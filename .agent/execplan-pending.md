# ExecPlan: Temp Inbox Aliases (AdGuard-Mail-style stored email)

> STATUS: Phases 1-5 implemented in commit 99b359b. Phase 6 (frontend app/wxt) pending.

## Goal

Add a second alias kind: **inbox**. A relay alias forwards mail to a verified
recipient; an inbox alias has no recipient — incoming mail is stored in the DB
and read in the web app. Inbox aliases carry a TTL and self-destruct.
Receive-only in v1 (no reply from temp inboxes — same as AdGuard temp mail).

## Facts about the current system (verified)

- Mail path: Postfix → `mailserver/daemon` (policy, domain-level only — **no change needed**) → `POST /v1/email` (`HandleEmail`) → `service.ProcessMessage` (`api/internal/service/processor.go`).
- `ProcessMessage` → `FindRecipients` (`api/internal/service/recipient.go:291`) which errors with `ErrNoRecipients` when an alias has no recipients. That error is exactly where inbox behavior forks.
- `model.Message` (`api/internal/model/message.go`) is a **metadata counter only** (ID, CreatedAt, UserID, AliasID, Type). Do not touch it; stats keep working.
- `model.Alias` (`api/internal/model/alias.go`): BaseModel + Name/UserID/Enabled/Recipients/CatchAll/...
- DB is **MySQL** (`api/internal/repository/db.go` uses `gorm.io/driver/mysql`; README's "SQLite" is stale). Cron jobs already use MySQL `NOW() - INTERVAL ? DAY` syntax.
- AutoMigrate list lives in `repository/db.go: migrate()`.
- Cron pattern: `api/internal/cron/cron.go` + one func per file in `cron/jobs/`.
- Routes: `api/internal/transport/api/routes.go`, v1 group auth'd via `auth.New`.
- `PostAlias` service (`service/alias.go:159`) already checks subscription + limits.

## Phase 1 — Model + migration

**`api/internal/model/alias.go`**
```go
type AliasType int

const (
    AliasRelay AliasType = 0
    AliasInbox AliasType = 1
)

// add to Alias struct:
Type      AliasType  `gorm:"default:0" json:"type"`
ExpiresAt *time.Time `json:"expires_at,omitempty"` // nil = never (relay aliases)
```
Existing rows default to `0` (relay) — zero-downtime migration via AutoMigrate.

**New `api/internal/model/inbox.go`**
```go
type InboxMessage struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UserID    string    `gorm:"index" json:"-"`
    AliasID   string    `gorm:"index" json:"alias_id"`
    From      string    `json:"from"`
    FromName  string    `json:"from_name"`
    Subject   string    `json:"subject"`
    Read      bool      `gorm:"default:false" json:"read"`
    Size      int       `json:"size"`
    Raw       []byte    `gorm:"type:mediumblob" json:"-"` // raw MIME, parsed on read
}
```
Cap raw size at **5 MB** (`MaxInboxMessageSize` const in model); oversize mail is
truncated-rejected in the processor (logged, dropped — same "fail silently"
posture the processor already uses for unauthenticated mail).

**`api/internal/repository/db.go`**: append `&model.InboxMessage{}` to AutoMigrate.

## Phase 2 — Processor fork (the core change)

**`api/internal/service/recipient.go` — `FindRecipients` (line ~291)**
After the enabled-alias and enabled-domain checks (lines 300–335), before the
Reply|Send block (line 337):
```go
// Inbox alias: mail is stored, not relayed. No recipients required.
if alias.Type == model.AliasInbox {
    return nil, alias, model.Inbox, nil
}
```

**`api/internal/model/message.go`**: add `Inbox MessageType = 5` (stats row type).

**`api/internal/service/processor.go` — `ProcessMessage`**
In the per-`to` loop, after `FindRecipients` succeeds and subscription is
loaded, before the recipients fan-out:
```go
if relayType == model.Inbox {
    if !sub.ActiveStatus() { continue } // same gate as Reply|Send
    if err := s.StoreInboxMessage(context.Background(), alias, msg, data); err != nil {
        log.Println("error storing inbox message", err)
    }
    continue
}
```

**New `api/internal/service/inbox.go`** — `InboxStore` interface (mirrors
`AliasStore` pattern) + service methods:
- `StoreInboxMessage(ctx, alias, msg, raw)` — enforce size cap, persist
  `InboxMessage{From: msg.From, FromName: msg.FromName, Subject: msg.Subject, Raw: raw, ...}`,
  then `SaveMessage(ctx, alias, model.Inbox)` for stats.
- `GetInboxMessages(ctx, aliasID, userID)` — metadata only (omit `Raw` via `Select`).
- `GetInboxMessage(ctx, id, userID)` — full row, marks `Read = true`.
- `DeleteInboxMessage(ctx, id, userID)`.
- Per-alias cap: keep newest 50 messages (delete oldest on insert — one query).

**New `api/internal/repository/inbox.go`** — GORM impls, all queries scoped by
`user_id` (same discipline as `repository/message.go`).

## Phase 3 — Reading mail safely (the only security-sensitive part)

**New deps (justified — trust boundary, not convenience):**
- `github.com/jhillyerd/enmime` — MIME parsing (nested multipart, charsets, QP/base64). Stdlib `mime/multipart` does not handle real-world mail.
- `github.com/microcosm-cc/bluemonday` — HTML sanitization.

**In `service/inbox.go`**: `RenderInboxMessage` returns
```go
type RenderedMessage struct {
    From, FromName, Subject string
    Date                    time.Time
    HTML                    string   // bluemonday.UGCPolicy() output, remote images stripped
    Text                    string
    Attachments             []string // names+sizes only; bodies not served in v1
}
```
Sanitize policy: `UGCPolicy()` minus `img` with remote `src` (rewrite to blocked
placeholder), no `style` attributes, `target="_blank" rel="noopener"` forced on links.

## Phase 4 — API surface

**`api/internal/transport/api/inbox.go`** (new) + wire in `routes.go` under the
existing `v1` auth group:
```
GET    /v1/alias/:id/inbox        → list (metadata)
GET    /v1/inbox/message/:id      → RenderedMessage
DELETE /v1/inbox/message/:id
```

**`PostAlias`** (`transport/api/alias.go` + `service/alias.go:159`): accept
`type` ("relay"|"inbox") and `ttl_hours` (bounded: 1h–720h, default 24h for
inbox). For inbox type: skip the recipients requirement, set
`ExpiresAt = now + ttl`. Relay path unchanged. Existing subscription/limit
checks apply to both.

**`FindRecipients` guard**: expired inbox alias (`ExpiresAt < now`) behaves as
disabled → reuse `ErrDisabledAlias` path (blocks + optional discard log for free).

## Phase 5 — Expiry cron

**New `api/internal/cron/jobs/inbox.go`**, both scheduled hourly in `cron.go`:
```go
// Soft-delete expired inbox aliases (CleanupDeletedAliases purges later)
db.Where("type = 1 AND expires_at IS NOT NULL AND expires_at < NOW()").Delete(&model.Alias{})
// Hard-delete stored mail older than 7 days regardless of alias state
db.Where("created_at < NOW() - INTERVAL ? DAY", 7).Delete(&model.InboxMessage{})
// Orphan sweep: inbox messages whose alias is soft-deleted
db.Where("alias_id IN (SELECT id FROM aliases WHERE deleted_at IS NOT NULL)").Delete(&model.InboxMessage{})
```

## Phase 6 — Frontend (separate PR)

- `app/src/api/inbox.ts` (fetch wrappers, mirrors `alias.ts`).
- Inbox view: message list + reader pane rendering sanitized HTML in a
  sandboxed iframe (`sandbox=""` attr — belt and suspenders on top of bluemonday).
- Alias create dialog: type toggle + TTL select.
- `wxt/` extension: "Generate temp email" button → `POST /alias {type: inbox, ttl_hours: 24}` → copy to clipboard.

## Validation

1. `cd api && go build ./... && go test ./...` — zero errors.
2. Integration (MailHog per README): create inbox alias via API, pipe a raw
   MIME file to `POST /v1/email` with PSK, assert message listed and rendered;
   send >5 MB mail, assert dropped; disable alias, assert Block stat row.
3. Sanitizer test (`service/inbox_test.go`): `<script>`, `onerror=`, remote
   `<img>`, `javascript:` links — all stripped. This is the one test that must exist.
4. Relay regression: existing forward path untouched (`relayType != model.Inbox`
   falls through to the exact current code).

## Explicitly out of scope (v1)

- Anonymous/no-account temp inboxes (whole auth model is user-scoped; abuse surface).
- Replying/sending from inbox aliases.
- Attachment download (list names only).
- Converting inbox → relay alias or vice versa.

## Order of work

Phases 1–2 (one commit, API functional via curl) → 3–4 (one commit) → 5 (one
commit) → 6 (separate PR). Each phase leaves the relay path fully working.
