package api

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"ivpn.net/email/api/internal/model"
)

type AdminService interface {
	GetAllUsers(context.Context) ([]model.User, error)
	GetSystemStats(context.Context) (any, error)
	GetAllLogs(context.Context) ([]model.Log, error)
	AdminUpdateUser(context.Context, model.User) error
	AdminDeleteUser(context.Context, string) error
	AdminAssignPlan(context.Context, string, string) error
	GetAllAliasesAdmin(context.Context, int, int, string) ([]model.Alias, int64, error)
	AdminDeleteAlias(context.Context, string) error
	AdminToggleAlias(context.Context, string, bool) error
	GetAllDomainsAdmin(context.Context) ([]model.Domain, error)
	AdminDeleteDomain(context.Context, string) error
	AdminToggleDomain(context.Context, string, bool) error
	GetAllRecipientsAdmin(context.Context, int, int, string) ([]model.Recipient, int64, error)
	AdminDeleteRecipient(context.Context, string) error
	GetLogsFiltered(context.Context, string, int, int) ([]model.Log, int64, error)
	AdminSearchUsers(context.Context, string, int, int) ([]model.User, int64, error)
	AdminGetUserDetail(context.Context, string) (model.User, model.Subscription, []model.Alias, []model.Recipient, []model.Domain, error)
	GetAllAccessKeysAdmin(context.Context, int, int) ([]model.AccessKey, int64, error)
	AdminDeleteAccessKey(context.Context, string) error
	GetAllSessionsAdmin(context.Context, int, int) ([]model.Session, int64, error)
	AdminDeleteSession(context.Context, string) error
	AdminForceLogout(context.Context, string) error
	GetAllCredentialsAdmin(context.Context, int, int) ([]model.Credential, int64, error)
	AdminDeleteCredential(context.Context, string) error
	AdminUpdateSubscription(context.Context, string, string, bool, string) error
	AdminBulkUpdateUsers(context.Context, []string, bool) error
	GetAllInboxMessagesAdmin(context.Context, int, int) ([]model.InboxMessage, int64, error)
	AdminDeleteInboxMessage(context.Context, uint) error
	AdminPurgeInboxByUser(context.Context, string) error
	AdminDisableTotp(context.Context, string) error
	AdminResetPassword(context.Context, string, string) error
	AdminGetSettings(context.Context, string) (model.Settings, error)
	AdminUpdateSettings(context.Context, string, map[string]interface{}) error
	AdminExportUsers(context.Context) ([]model.User, error)
	AdminExportAliases(context.Context) ([]model.Alias, error)
	GetAllSubscriptionsAdmin(context.Context, int, int, string) ([]model.Subscription, int64, error)
	AdminDeleteSubscription(context.Context, string) error
	AdminImpersonate(context.Context, string) (model.User, error)
	AdminBulkDeleteAliases(context.Context, []string) error
	AdminBulkDeleteDomains(context.Context, []string) error
	AdminBulkDeleteRecipients(context.Context, []string) error
	GetTableSizes(context.Context) (map[string]int64, error)
	GetRecentSignups(context.Context, int) ([]model.User, error)
	AdminVerifyDomain(context.Context, string, bool) error
	AdminImpersonateUser(context.Context, string) (string, error)
	SearchAccessKeys(context.Context, string, int, int) ([]model.AccessKey, int64, error)
	SearchSessions(context.Context, string, int, int) ([]model.Session, int64, error)
	SearchInboxMessages(context.Context, string, int, int) ([]model.InboxMessage, int64, error)
	GetAllMessagesAdmin(context.Context, int, int, string) ([]model.Message, int64, error)
	AdminGetUserStats(context.Context, string) (model.UserStats, error)
	SearchLogs(context.Context, string, string, int, int) ([]model.Log, int64, error)
	AdminToggleRecipient(context.Context, string, bool) error
	SearchDomainsAdmin(context.Context, string) ([]model.Domain, error)
	AdminExportRecipients(context.Context) ([]model.Recipient, error)
	AdminExportSubscriptions(context.Context) ([]model.Subscription, error)
	AdminChangeEmail(context.Context, string, string) error
	AdminExportDomains(context.Context) ([]model.Domain, error)
	AdminExportLogs(context.Context) ([]model.Log, error)
	AdminBulkDeleteUsers(context.Context, []string) error
	SearchMessages(context.Context, string, string, int, int) ([]model.Message, int64, error)
	AdminToggleRecipientPGP(context.Context, string, bool) error
	AdminRemoveRecipientPGPKey(context.Context, string) error
	AdminUpdateAlias(context.Context, string, map[string]interface{}) error
	AdminUpdateDomain(context.Context, string, map[string]interface{}) error
	AdminMarkInboxRead(context.Context, uint, bool) error
	AdminGetAllUsersPaginated(context.Context, int, int, string) ([]model.User, int64, error)
	AdminCreateRecipient(context.Context, model.Recipient) error
	AdminCreateDomain(context.Context, model.Domain) error
	AdminExportInbox(context.Context) ([]model.InboxMessage, error)
	AdminExportMessages(context.Context) ([]model.Message, error)
	AdminCreateAlias(context.Context, model.Alias) error
	AdminUpdateRecipient(context.Context, string, map[string]interface{}) error
	AdminDeleteLog(context.Context, string) error
	AdminBulkDeleteInbox(context.Context, []uint) error
	AdminExtendSubscription(context.Context, string, int) error
	AdminCreateAccessKey(context.Context, string, string) (string, error)
	AdminTransferAlias(context.Context, string, string) error
	AdminTransferDomain(context.Context, string, string) error
	AdminPurgeLogs(context.Context, int, string) (int64, error)
	AdminPurgeAllInbox(context.Context) (int64, error)
	AdminCreateUser(context.Context, string, string) (model.User, error)
	AdminGetInboxRaw(context.Context, uint) ([]byte, error)
	AdminSetAliasExpiry(context.Context, string, *time.Time) error
	AdminSetAccessKeyExpiry(context.Context, string, *time.Time) error
	LogAdminAction(context.Context, string, string, string, string)
	AdminGetAuditLog(context.Context, int, int) ([]model.AdminAudit, int64, error)
	AdminGetSessionData(context.Context, string) ([]byte, error)
	AdminGetLogsDateRange(context.Context, string, string, string, int, int) ([]model.Log, int64, error)
	AdminBulkDeleteAccessKeys(context.Context, []string) error
	AdminBulkDeleteCredentials(context.Context, []string) error
	AdminBulkExtendSubscriptions(context.Context, []string, int) (int64, error)
	AdminExportUsersEnriched(context.Context) ([]model.UserWithSub, error)
	AdminBulkDeleteMessages(context.Context, []uint) error
	AdminSetRecipientPGP(context.Context, string, string, bool) error
	AdminGetDomainDNS(context.Context, string) (model.DNSConfig, error)
	AdminUpdateUserNotes(context.Context, string, string) error
	AdminGetSubscriptionStats(context.Context) (int64, int64, int64, error)
}

func (h *Handler) AdminGetUsers(c *fiber.Ctx) error {
	users, err := h.Service.GetAllUsers(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch users"})
	}
	return c.JSON(users)
}

func (h *Handler) AdminGetStats(c *fiber.Ctx) error {
	stats, err := h.Service.GetSystemStats(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch stats"})
	}
	return c.JSON(stats)
}

func (h *Handler) AdminGetLogs(c *fiber.Ctx) error {
	logs, err := h.Service.GetAllLogs(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch logs"})
	}
	return c.JSON(logs)
}

type AdminUpdateUserReq struct {
	ID       string `json:"id" validate:"required,uuid"`
	IsActive *bool  `json:"is_active"`
	IsAdmin  *bool  `json:"is_admin"`
}

func (h *Handler) AdminUpdateUser(c *fiber.Ctx) error {
	var req AdminUpdateUserReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	user, err := h.Service.GetUser(c.Context(), req.ID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}
	if req.IsAdmin != nil {
		user.IsAdmin = *req.IsAdmin
	}

	if err := h.Service.AdminUpdateUser(c.Context(), user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to update user"})
	}

	return c.JSON(fiber.Map{"message": "User updated"})
}

func (h *Handler) AdminDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "User ID required"})
	}

	if err := h.Service.AdminDeleteUser(c.Context(), id); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to delete user"})
	}

	h.audit(c, "delete_user", id, "")
	return c.JSON(fiber.Map{"message": "User deleted"})
}

type AdminAssignPlanReq struct {
	UserID string `json:"user_id" validate:"required,uuid"`
	PlanID string `json:"plan_id" validate:"required,uuid"`
}

func (h *Handler) AdminAssignPlan(c *fiber.Ctx) error {
	var req AdminAssignPlanReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.Service.AdminAssignPlan(c.Context(), req.UserID, req.PlanID); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to assign plan"})
	}

	return c.JSON(fiber.Map{"message": "Plan assigned"})
}

func (h *Handler) AdminGetAliases(c *fiber.Ctx) error {
	search := c.Query("search", "")
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)
	aliases, total, err := h.Service.GetAllAliasesAdmin(c.Context(), limit, offset, search)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch aliases"})
	}
	return c.JSON(fiber.Map{"aliases": aliases, "total": total})
}

func (h *Handler) AdminDeleteAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Alias ID required"})
	}
	if err := h.Service.AdminDeleteAlias(c.Context(), id); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to delete alias"})
	}
	return c.JSON(fiber.Map{"message": "Alias deleted"})
}

func (h *Handler) AdminToggleAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Alias ID required"})
	}
	var req struct {
		Enabled bool `json:"enabled"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.Service.AdminToggleAlias(c.Context(), id, req.Enabled); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to toggle alias"})
	}
	return c.JSON(fiber.Map{"message": "Alias updated"})
}

func (h *Handler) AdminGetDomains(c *fiber.Ctx) error {
	domains, err := h.Service.GetAllDomainsAdmin(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch domains"})
	}
	return c.JSON(domains)
}

func (h *Handler) AdminDeleteDomain(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Domain ID required"})
	}
	if err := h.Service.AdminDeleteDomain(c.Context(), id); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to delete domain"})
	}
	return c.JSON(fiber.Map{"message": "Domain deleted"})
}

func (h *Handler) AdminToggleDomain(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Domain ID required"})
	}
	var req struct {
		Enabled bool `json:"enabled"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.Service.AdminToggleDomain(c.Context(), id, req.Enabled); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to toggle domain"})
	}
	return c.JSON(fiber.Map{"message": "Domain updated"})
}

func (h *Handler) AdminGetRecipients(c *fiber.Ctx) error {
	search := c.Query("search", "")
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)
	recipients, total, err := h.Service.GetAllRecipientsAdmin(c.Context(), limit, offset, search)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch recipients"})
	}
	return c.JSON(fiber.Map{"recipients": recipients, "total": total})
}

func (h *Handler) AdminDeleteRecipient(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Recipient ID required"})
	}
	if err := h.Service.AdminDeleteRecipient(c.Context(), id); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to delete recipient"})
	}
	return c.JSON(fiber.Map{"message": "Recipient deleted"})
}

func (h *Handler) AdminGetLogsFiltered(c *fiber.Ctx) error {
	logType := c.Query("type", "")
	limit := c.QueryInt("limit", 100)
	offset := c.QueryInt("offset", 0)
	logs, total, err := h.Service.GetLogsFiltered(c.Context(), logType, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch logs"})
	}
	return c.JSON(fiber.Map{"logs": logs, "total": total})
}

func (h *Handler) AdminSearchUsers(c *fiber.Ctx) error {
	search := c.Query("search", "")
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)
	users, total, err := h.Service.AdminSearchUsers(c.Context(), search, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to search users"})
	}
	return c.JSON(fiber.Map{"users": users, "total": total})
}

func (h *Handler) AdminGetUserDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "User ID required"})
	}
	user, sub, aliases, recipients, domains, err := h.Service.AdminGetUserDetail(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(fiber.Map{
		"user":       user,
		"subscription": sub,
		"aliases":    aliases,
		"recipients": recipients,
		"domains":    domains,
	})
}

func (h *Handler) AdminGetAccessKeys(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)
	keys, total, err := h.Service.GetAllAccessKeysAdmin(c.Context(), limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch access keys"})
	}
	return c.JSON(fiber.Map{"keys": keys, "total": total})
}

func (h *Handler) AdminDeleteAccessKey(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Key ID required"})
	}
	if err := h.Service.AdminDeleteAccessKey(c.Context(), id); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to delete key"})
	}
	return c.JSON(fiber.Map{"message": "Key revoked"})
}

func (h *Handler) AdminGetSessions(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)
	sessions, total, err := h.Service.GetAllSessionsAdmin(c.Context(), limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch sessions"})
	}
	return c.JSON(fiber.Map{"sessions": sessions, "total": total})
}

func (h *Handler) AdminDeleteSession(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Session ID required"})
	}
	if err := h.Service.AdminDeleteSession(c.Context(), id); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to delete session"})
	}
	return c.JSON(fiber.Map{"message": "Session terminated"})
}

func (h *Handler) AdminForceLogout(c *fiber.Ctx) error {
	userID := c.Params("id")
	if userID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "User ID required"})
	}
	if err := h.Service.AdminForceLogout(c.Context(), userID); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to force logout"})
	}
	return c.JSON(fiber.Map{"message": "All sessions terminated"})
}

func (h *Handler) AdminGetCredentials(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)
	creds, total, err := h.Service.GetAllCredentialsAdmin(c.Context(), limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch credentials"})
	}
	return c.JSON(fiber.Map{"credentials": creds, "total": total})
}

func (h *Handler) AdminDeleteCredential(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Credential ID required"})
	}
	if err := h.Service.AdminDeleteCredential(c.Context(), id); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to delete credential"})
	}
	return c.JSON(fiber.Map{"message": "Passkey removed"})
}

type AdminUpdateSubReq struct {
	UserID      string `json:"user_id" validate:"required,uuid"`
	Tier        string `json:"tier"`
	IsActive    bool   `json:"is_active"`
	ActiveUntil string `json:"active_until"`
}

func (h *Handler) AdminUpdateSubscription(c *fiber.Ctx) error {
	var req AdminUpdateSubReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.Service.AdminUpdateSubscription(c.Context(), req.UserID, req.Tier, req.IsActive, req.ActiveUntil); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to update subscription"})
	}
	return c.JSON(fiber.Map{"message": "Subscription updated"})
}

type AdminBulkReq struct {
	UserIDs  []string `json:"user_ids" validate:"required"`
	IsActive bool     `json:"is_active"`
}

func (h *Handler) AdminBulkUpdateUsers(c *fiber.Ctx) error {
	var req AdminBulkReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.Service.AdminBulkUpdateUsers(c.Context(), req.UserIDs, req.IsActive); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to bulk update"})
	}
	h.audit(c, "bulk_update_users", fmt.Sprintf("%d users", len(req.UserIDs)), fmt.Sprintf("is_active=%v", req.IsActive))
	return c.JSON(fiber.Map{"message": "Users updated"})
}

func (h *Handler) AdminGetInboxMessages(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)
	msgs, total, err := h.Service.GetAllInboxMessagesAdmin(c.Context(), limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch inbox messages"})
	}
	return c.JSON(fiber.Map{"messages": msgs, "total": total})
}

func (h *Handler) AdminDeleteInboxMessage(c *fiber.Ctx) error {
	id := c.Params("id")
	msgID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid message ID"})
	}
	if err := h.Service.AdminDeleteInboxMessage(c.Context(), uint(msgID)); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to delete message"})
	}
	return c.JSON(fiber.Map{"message": "Inbox message deleted"})
}

func (h *Handler) AdminPurgeInbox(c *fiber.Ctx) error {
	userID := c.Params("id")
	if userID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "User ID required"})
	}
	if err := h.Service.AdminPurgeInboxByUser(c.Context(), userID); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to purge inbox"})
	}
	return c.JSON(fiber.Map{"message": "Inbox purged"})
}

func (h *Handler) AdminDisableTotp(c *fiber.Ctx) error {
	userID := c.Params("id")
	if userID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "User ID required"})
	}
	if err := h.Service.AdminDisableTotp(c.Context(), userID); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to disable TOTP"})
	}
	return c.JSON(fiber.Map{"message": "TOTP disabled"})
}

type AdminResetPasswordReq struct {
	UserID   string `json:"user_id" validate:"required,uuid"`
	Password string `json:"password" validate:"required,min=12"`
}

func (h *Handler) AdminResetPassword(c *fiber.Ctx) error {
	var req AdminResetPasswordReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.Service.AdminResetPassword(c.Context(), req.UserID, req.Password); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to reset password"})
	}
	return c.JSON(fiber.Map{"message": "Password reset"})
}

func (h *Handler) AdminGetSettings(c *fiber.Ctx) error {
	userID := c.Params("id")
	if userID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "User ID required"})
	}
	settings, err := h.Service.AdminGetSettings(c.Context(), userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Settings not found"})
	}
	return c.JSON(settings)
}

type AdminUpdateSettingsReq struct {
	UserID      string `json:"user_id" validate:"required,uuid"`
	Domain      string `json:"domain"`
	Recipient   string `json:"recipient"`
	FromName    string `json:"from_name"`
	AliasFormat string `json:"alias_format"`
	LogIssues   *bool  `json:"log_issues"`
	RemoveHeader *bool `json:"remove_header"`
}

func (h *Handler) AdminUpdateSettings(c *fiber.Ctx) error {
	var req AdminUpdateSettingsReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	updates := map[string]interface{}{}
	if req.Domain != "" { updates["domain"] = req.Domain }
	if req.Recipient != "" { updates["recipient"] = req.Recipient }
	if req.FromName != "" { updates["from_name"] = req.FromName }
	if req.AliasFormat != "" { updates["alias_format"] = req.AliasFormat }
	if req.LogIssues != nil { updates["log_issues"] = *req.LogIssues }
	if req.RemoveHeader != nil { updates["remove_header"] = *req.RemoveHeader }
	if err := h.Service.AdminUpdateSettings(c.Context(), req.UserID, updates); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to update settings"})
	}
	return c.JSON(fiber.Map{"message": "Settings updated"})
}

func (h *Handler) AdminExportUsers(c *fiber.Ctx) error {
	users, err := h.Service.AdminExportUsers(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to export users"})
	}
	var buf bytes.Buffer
	buf.WriteString("id,email,is_active,is_admin,created_at\n")
	for _, u := range users {
		buf.WriteString(fmt.Sprintf("%s,%s,%t,%t,%s\n", u.ID, u.Email, u.IsActive, u.IsAdmin, u.CreatedAt.Format("2006-01-02")))
	}
	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", "attachment; filename=users.csv")
	return c.Send(buf.Bytes())
}

func (h *Handler) AdminExportAliases(c *fiber.Ctx) error {
	aliases, err := h.Service.AdminExportAliases(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to export aliases"})
	}
	var buf bytes.Buffer
	buf.WriteString("id,user_id,name,enabled,created_at\n")
	for _, a := range aliases {
		buf.WriteString(fmt.Sprintf("%s,%s,%s,%t,%s\n", a.ID, a.UserID, a.Name, a.Enabled, a.CreatedAt.Format("2006-01-02")))
	}
	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", "attachment; filename=aliases.csv")
	return c.Send(buf.Bytes())
}

func (h *Handler) AdminGetSubscriptions(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)
	tier := c.Query("tier", "")
	subs, total, err := h.Service.GetAllSubscriptionsAdmin(c.Context(), limit, offset, tier)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch subscriptions"})
	}
	return c.JSON(fiber.Map{"subscriptions": subs, "total": total})
}

func (h *Handler) AdminDeleteSubscription(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Subscription ID required"})
	}
	if err := h.Service.AdminDeleteSubscription(c.Context(), id); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to delete subscription"})
	}
	return c.JSON(fiber.Map{"message": "Subscription deleted"})
}

type AdminBulkDeleteReq struct {
	IDs []string `json:"ids"`
}

func (h *Handler) AdminBulkDeleteAliases(c *fiber.Ctx) error {
	var req AdminBulkDeleteReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.Service.AdminBulkDeleteAliases(c.Context(), req.IDs); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to bulk delete"})
	}
	return c.JSON(fiber.Map{"message": "Aliases deleted"})
}

func (h *Handler) AdminBulkDeleteDomains(c *fiber.Ctx) error {
	var req AdminBulkDeleteReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.Service.AdminBulkDeleteDomains(c.Context(), req.IDs); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to bulk delete"})
	}
	return c.JSON(fiber.Map{"message": "Domains deleted"})
}

func (h *Handler) AdminBulkDeleteRecipients(c *fiber.Ctx) error {
	var req AdminBulkDeleteReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.Service.AdminBulkDeleteRecipients(c.Context(), req.IDs); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to bulk delete"})
	}
	return c.JSON(fiber.Map{"message": "Recipients deleted"})
}

func (h *Handler) AdminGetTableSizes(c *fiber.Ctx) error {
	sizes, err := h.Service.GetTableSizes(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch table sizes"})
	}
	return c.JSON(sizes)
}

func (h *Handler) AdminGetRecentSignups(c *fiber.Ctx) error {
	days := c.QueryInt("days", 7)
	users, err := h.Service.GetRecentSignups(c.Context(), days)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch recent signups"})
	}
	return c.JSON(fiber.Map{"users": users, "count": len(users)})
}

type AdminVerifyDomainReq struct {
	Verified bool `json:"verified"`
}

func (h *Handler) AdminVerifyDomain(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Domain ID required"})
	}
	var req AdminVerifyDomainReq
	c.BodyParser(&req)
	if err := h.Service.AdminVerifyDomain(c.Context(), id, req.Verified); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to update domain verification"})
	}
	return c.JSON(fiber.Map{"message": "Domain verification updated"})
}

func (h *Handler) AdminImpersonate(c *fiber.Ctx) error {
	userID := c.Params("id")
	if userID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "User ID required"})
	}
	token, err := h.Service.AdminImpersonateUser(c.Context(), userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to impersonate user"})
	}
	c.Cookie(&fiber.Cookie{Name: "authn", Value: token, HTTPOnly: true, Secure: true, MaxAge: 86400})
	return c.JSON(fiber.Map{"token": token, "message": "Impersonation session created"})
}

func (h *Handler) AdminSearchAccessKeys(c *fiber.Ctx) error {
	userID := c.Query("user_id", "")
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)
	keys, total, err := h.Service.SearchAccessKeys(c.Context(), userID, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to search access keys"})
	}
	return c.JSON(fiber.Map{"keys": keys, "total": total})
}

func (h *Handler) AdminSearchSessions(c *fiber.Ctx) error {
	userID := c.Query("user_id", "")
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)
	sessions, total, err := h.Service.SearchSessions(c.Context(), userID, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to search sessions"})
	}
	return c.JSON(fiber.Map{"sessions": sessions, "total": total})
}

func (h *Handler) AdminSearchInbox(c *fiber.Ctx) error {
	search := c.Query("search", "")
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)
	msgs, total, err := h.Service.SearchInboxMessages(c.Context(), search, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to search inbox"})
	}
	return c.JSON(fiber.Map{"messages": msgs, "total": total})
}

func (h *Handler) AdminGetMessages(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)
	msgType := c.Query("type", "")
	msgs, total, err := h.Service.GetAllMessagesAdmin(c.Context(), limit, offset, msgType)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch messages"})
	}
	return c.JSON(fiber.Map{"messages": msgs, "total": total})
}

func (h *Handler) AdminGetUserStats(c *fiber.Ctx) error {
	userID := c.Params("id")
	if userID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "User ID required"})
	}
	stats, err := h.Service.AdminGetUserStats(c.Context(), userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch user stats"})
	}
	return c.JSON(stats)
}

func (h *Handler) AdminSearchLogs(c *fiber.Ctx) error {
	search := c.Query("search", "")
	logType := c.Query("type", "")
	limit := c.QueryInt("limit", 100)
	offset := c.QueryInt("offset", 0)
	logs, total, err := h.Service.SearchLogs(c.Context(), search, logType, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to search logs"})
	}
	return c.JSON(fiber.Map{"logs": logs, "total": total})
}

type AdminToggleRecipientReq struct {
	IsActive bool `json:"is_active"`
}

func (h *Handler) AdminToggleRecipient(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Recipient ID required"})
	}
	var req AdminToggleRecipientReq
	c.BodyParser(&req)
	if err := h.Service.AdminToggleRecipient(c.Context(), id, req.IsActive); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to toggle recipient"})
	}
	return c.JSON(fiber.Map{"message": "Recipient updated"})
}

func (h *Handler) AdminSearchDomains(c *fiber.Ctx) error {
	search := c.Query("search", "")
	domains, err := h.Service.SearchDomainsAdmin(c.Context(), search)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to search domains"})
	}
	return c.JSON(fiber.Map{"domains": domains, "total": len(domains)})
}

func (h *Handler) AdminExportRecipients(c *fiber.Ctx) error {
	recipients, err := h.Service.AdminExportRecipients(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to export recipients"})
	}
	var buf bytes.Buffer
	buf.WriteString("id,user_id,email,is_active,created_at\n")
	for _, r := range recipients {
		buf.WriteString(fmt.Sprintf("%s,%s,%s,%t,%s\n", r.ID, r.UserID, r.Email, r.IsActive, r.CreatedAt.Format("2006-01-02")))
	}
	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", "attachment; filename=recipients.csv")
	return c.Send(buf.Bytes())
}

func (h *Handler) AdminExportSubscriptions(c *fiber.Ctx) error {
	subs, err := h.Service.AdminExportSubscriptions(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to export subscriptions"})
	}
	var buf bytes.Buffer
	buf.WriteString("id,user_id,type,tier,is_active,active_until,created_at\n")
	for _, s := range subs {
		buf.WriteString(fmt.Sprintf("%s,%s,%s,%s,%t,%s,%s\n", s.ID, s.UserID, s.Type, s.Tier, s.IsActive, s.ActiveUntil.Format("2006-01-02"), s.CreatedAt.Format("2006-01-02")))
	}
	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", "attachment; filename=subscriptions.csv")
	return c.Send(buf.Bytes())
}

type AdminChangeEmailReq struct {
	UserID   string `json:"user_id" validate:"required,uuid"`
	NewEmail string `json:"new_email" validate:"required,email"`
}

func (h *Handler) AdminChangeEmail(c *fiber.Ctx) error {
	var req AdminChangeEmailReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.Service.AdminChangeEmail(c.Context(), req.UserID, req.NewEmail); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to change email"})
	}
	return c.JSON(fiber.Map{"message": "Email changed"})
}

func (h *Handler) AdminExportDomains(c *fiber.Ctx) error {
	domains, err := h.Service.AdminExportDomains(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to export domains"})
	}
	var buf bytes.Buffer
	buf.WriteString("id,user_id,name,enabled,mx_verified,created_at\n")
	for _, d := range domains {
		verified := d.MXVerifiedAt != nil
		buf.WriteString(fmt.Sprintf("%s,%s,%s,%t,%t,%s\n", d.ID, d.UserID, d.Name, d.Enabled, verified, d.CreatedAt.Format("2006-01-02")))
	}
	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", "attachment; filename=domains.csv")
	return c.Send(buf.Bytes())
}

func (h *Handler) AdminExportLogs(c *fiber.Ctx) error {
	logs, err := h.Service.AdminExportLogs(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to export logs"})
	}
	var buf bytes.Buffer
	buf.WriteString("id,created_at,log_type,from,destination,status,message\n")
	for _, l := range logs {
		buf.WriteString(fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s\n", l.ID, l.CreatedAt.Format("2006-01-02 15:04"), l.Type, l.From, l.Destination, l.Status, l.Message))
	}
	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", "attachment; filename=logs.csv")
	return c.Send(buf.Bytes())
}

func (h *Handler) AdminBulkDeleteUsers(c *fiber.Ctx) error {
	var req AdminBulkReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.Service.AdminBulkDeleteUsers(c.Context(), req.UserIDs); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to bulk delete users"})
	}
	h.audit(c, "bulk_delete_users", fmt.Sprintf("%d users", len(req.UserIDs)), "")
	return c.JSON(fiber.Map{"message": "Users deleted"})
}

func (h *Handler) AdminGetConfig(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"fqdn":                h.Cfg.FQDN,
		"name":                h.Cfg.Name,
		"port":                h.Cfg.Port,
		"api_allow_origin":    h.Cfg.ApiAllowOrigin,
		"domains":             h.Cfg.Domains,
		"token_expiration":    h.Cfg.TokenExpiration.String(),
		"admin_emails":        h.Cfg.AdminEmails,
		"preauth_url_set":     h.Cfg.PreauthURL != "",
		"preauth_psk_set":     h.Cfg.PreauthPSK != "",
		"signup_webhook_set":  h.Cfg.SignupWebhookURL != "",
		"smtp_configured":     false,
		"oxapay_configured":   false,
	})
}

func (h *Handler) AdminSearchMessages(c *fiber.Ctx) error {
	search := c.Query("search", "")
	msgType := c.Query("type", "")
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)
	msgs, total, err := h.Service.SearchMessages(c.Context(), search, msgType, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to search messages"})
	}
	return c.JSON(fiber.Map{"messages": msgs, "total": total})
}

type AdminTogglePGPReq struct {
	PGPEnabled bool `json:"pgp_enabled"`
}

func (h *Handler) AdminToggleRecipientPGP(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Recipient ID required"})
	}
	var req AdminTogglePGPReq
	c.BodyParser(&req)
	if err := h.Service.AdminToggleRecipientPGP(c.Context(), id, req.PGPEnabled); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to toggle PGP"})
	}
	return c.JSON(fiber.Map{"message": "PGP updated"})
}

func (h *Handler) AdminRemoveRecipientPGPKey(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Recipient ID required"})
	}
	if err := h.Service.AdminRemoveRecipientPGPKey(c.Context(), id); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to remove PGP key"})
	}
	return c.JSON(fiber.Map{"message": "PGP key removed"})
}

type AdminUpdateAliasReq struct {
	Description string `json:"description"`
	Recipients  string `json:"recipients"`
	FromName    string `json:"from_name"`
}

func (h *Handler) AdminUpdateAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Alias ID required"})
	}
	var req AdminUpdateAliasReq
	c.BodyParser(&req)
	updates := map[string]interface{}{}
	if req.Description != "" { updates["description"] = req.Description }
	if req.Recipients != "" { updates["recipients"] = req.Recipients }
	if req.FromName != "" { updates["from_name"] = req.FromName }
	if len(updates) == 0 { return c.Status(400).JSON(fiber.Map{"error": "No updates provided"}) }
	if err := h.Service.AdminUpdateAlias(c.Context(), id, updates); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to update alias"})
	}
	return c.JSON(fiber.Map{"message": "Alias updated"})
}

type AdminUpdateDomainReq struct {
	Description string `json:"description"`
	Recipient   string `json:"recipient"`
	FromName    string `json:"from_name"`
}

func (h *Handler) AdminUpdateDomain(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Domain ID required"})
	}
	var req AdminUpdateDomainReq
	c.BodyParser(&req)
	updates := map[string]interface{}{}
	if req.Description != "" { updates["description"] = req.Description }
	if req.Recipient != "" { updates["recipient"] = req.Recipient }
	if req.FromName != "" { updates["from_name"] = req.FromName }
	if len(updates) == 0 { return c.Status(400).JSON(fiber.Map{"error": "No updates provided"}) }
	if err := h.Service.AdminUpdateDomain(c.Context(), id, updates); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to update domain"})
	}
	return c.JSON(fiber.Map{"message": "Domain updated"})
}

type AdminMarkReadReq struct {
	IsRead bool `json:"is_read"`
}

func (h *Handler) AdminMarkInboxRead(c *fiber.Ctx) error {
	id := c.Params("id")
	msgID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid message ID"})
	}
	var req AdminMarkReadReq
	c.BodyParser(&req)
	if err := h.Service.AdminMarkInboxRead(c.Context(), uint(msgID), req.IsRead); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to update message"})
	}
	return c.JSON(fiber.Map{"message": "Message updated"})
}

func (h *Handler) AdminGetUsersPaginated(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)
	search := c.Query("search", "")
	users, total, err := h.Service.AdminGetAllUsersPaginated(c.Context(), limit, offset, search)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch users"})
	}
	return c.JSON(fiber.Map{"users": users, "total": total, "limit": limit, "offset": offset})
}

type AdminCreateRecipientReq struct {
	UserID string `json:"user_id" validate:"required,uuid"`
	Email  string `json:"email" validate:"required,email"`
}

func (h *Handler) AdminCreateRecipient(c *fiber.Ctx) error {
	var req AdminCreateRecipientReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	r := model.Recipient{UserID: req.UserID, Email: req.Email}
	if err := h.Service.AdminCreateRecipient(c.Context(), r); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to create recipient"})
	}
	return c.JSON(fiber.Map{"message": "Recipient created"})
}

type AdminCreateDomainReq struct {
	UserID string `json:"user_id" validate:"required,uuid"`
	Name   string `json:"name" validate:"required"`
}

func (h *Handler) AdminCreateDomain(c *fiber.Ctx) error {
	var req AdminCreateDomainReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	dm := model.Domain{UserID: req.UserID, Name: req.Name}
	if err := h.Service.AdminCreateDomain(c.Context(), dm); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to create domain"})
	}
	return c.JSON(fiber.Map{"message": "Domain created"})
}

func (h *Handler) AdminExportInbox(c *fiber.Ctx) error {
	msgs, err := h.Service.AdminExportInbox(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to export inbox messages"})
	}
	var buf bytes.Buffer
	buf.WriteString("id,user_id,alias_id,from,from_name,subject,read,size,created_at\n")
	for _, m := range msgs {
		buf.WriteString(fmt.Sprintf("%d,%s,%s,%s,%s,%s,%t,%d,%s\n", m.ID, m.UserID, m.AliasID, m.From, m.FromName, m.Subject, m.Read, m.Size, m.CreatedAt.Format(time.RFC3339)))
	}
	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", "attachment; filename=inbox_messages.csv")
	return c.Send(buf.Bytes())
}

func (h *Handler) AdminExportMessages(c *fiber.Ctx) error {
	msgs, err := h.Service.AdminExportMessages(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to export messages"})
	}
	var buf bytes.Buffer
	buf.WriteString("id,user_id,alias_id,type,created_at\n")
	for _, m := range msgs {
		buf.WriteString(fmt.Sprintf("%d,%s,%s,%s,%s\n", m.ID, m.UserID, m.AliasID, m.Type, m.CreatedAt.Format(time.RFC3339)))
	}
	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", "attachment; filename=messages.csv")
	return c.Send(buf.Bytes())
}

type AdminCreateAliasReq struct {
	UserID  string `json:"user_id" validate:"required,uuid"`
	Name    string `json:"name" validate:"required"`
	Enabled bool   `json:"enabled"`
}

func (h *Handler) AdminCreateAlias(c *fiber.Ctx) error {
	var req AdminCreateAliasReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	a := model.Alias{UserID: req.UserID, Name: req.Name, Enabled: req.Enabled}
	if err := h.Service.AdminCreateAlias(c.Context(), a); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to create alias"})
	}
	return c.JSON(fiber.Map{"message": "Alias created"})
}

type AdminUpdateRecipientReq struct {
	Email string `json:"email"`
}

func (h *Handler) AdminUpdateRecipient(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Recipient ID required"})
	}
	var req AdminUpdateRecipientReq
	c.BodyParser(&req)
	updates := map[string]interface{}{}
	if req.Email != "" { updates["email"] = req.Email }
	if len(updates) == 0 { return c.Status(400).JSON(fiber.Map{"error": "No updates provided"}) }
	if err := h.Service.AdminUpdateRecipient(c.Context(), id, updates); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to update recipient"})
	}
	return c.JSON(fiber.Map{"message": "Recipient updated"})
}

func (h *Handler) AdminDeleteLog(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Log ID required"})
	}
	if err := h.Service.AdminDeleteLog(c.Context(), id); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to delete log"})
	}
	return c.JSON(fiber.Map{"message": "Log deleted"})
}

type AdminBulkDeleteInboxReq struct {
	IDs []uint `json:"ids"`
}

func (h *Handler) AdminBulkDeleteInbox(c *fiber.Ctx) error {
	var req AdminBulkDeleteInboxReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if len(req.IDs) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "No IDs provided"})
	}
	if err := h.Service.AdminBulkDeleteInbox(c.Context(), req.IDs); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to delete messages"})
	}
	return c.JSON(fiber.Map{"message": fmt.Sprintf("%d messages deleted", len(req.IDs))})
}

type AdminExtendSubReq struct {
	Days int `json:"days"`
}

func (h *Handler) AdminExtendSubscription(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Subscription ID required"})
	}
	var req AdminExtendSubReq
	if err := c.BodyParser(&req); err != nil || req.Days == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Days required"})
	}
	if err := h.Service.AdminExtendSubscription(c.Context(), id, req.Days); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to extend subscription"})
	}
	return c.JSON(fiber.Map{"message": fmt.Sprintf("Subscription extended by %d days", req.Days)})
}

type AdminCreateAccessKeyReq struct {
	UserID string `json:"user_id" validate:"required,uuid"`
	Name   string `json:"name" validate:"required"`
}

func (h *Handler) AdminCreateAccessKey(c *fiber.Ctx) error {
	var req AdminCreateAccessKeyReq
	if err := c.BodyParser(&req); err != nil || req.UserID == "" || req.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "user_id and name required"})
	}
	key, err := h.Service.AdminCreateAccessKey(c.Context(), req.UserID, req.Name)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to create access key"})
	}
	return c.JSON(fiber.Map{"key": key, "message": "Access key created. Save it now - it will not be shown again."})
}

type AdminTransferReq struct {
	NewUserID string `json:"new_user_id" validate:"required,uuid"`
}

func (h *Handler) AdminTransferAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var req AdminTransferReq
	if err := c.BodyParser(&req); err != nil || req.NewUserID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "new_user_id required"})
	}
	if err := h.Service.AdminTransferAlias(c.Context(), id, req.NewUserID); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to transfer alias"})
	}
	h.audit(c, "transfer_alias", id, req.NewUserID)
	return c.JSON(fiber.Map{"message": "Alias transferred"})
}

func (h *Handler) AdminTransferDomain(c *fiber.Ctx) error {
	id := c.Params("id")
	var req AdminTransferReq
	if err := c.BodyParser(&req); err != nil || req.NewUserID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "new_user_id required"})
	}
	if err := h.Service.AdminTransferDomain(c.Context(), id, req.NewUserID); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to transfer domain"})
	}
	h.audit(c, "transfer_domain", id, req.NewUserID)
	return c.JSON(fiber.Map{"message": "Domain transferred"})
}

type AdminPurgeLogsReq struct {
	Days    int    `json:"days"`
	LogType string `json:"log_type"`
}

func (h *Handler) AdminPurgeLogs(c *fiber.Ctx) error {
	var req AdminPurgeLogsReq
	c.BodyParser(&req)
	count, err := h.Service.AdminPurgeLogs(c.Context(), req.Days, req.LogType)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to purge logs"})
	}
	h.audit(c, "purge_logs", fmt.Sprintf("%d days", req.Days), fmt.Sprintf("type=%s count=%d", req.LogType, count))
	return c.JSON(fiber.Map{"message": fmt.Sprintf("%d logs purged", count)})
}

func (h *Handler) AdminPurgeAllInbox(c *fiber.Ctx) error {
	count, err := h.Service.AdminPurgeAllInbox(c.Context())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to purge inbox"})
	}
	h.audit(c, "purge_all_inbox", "", fmt.Sprintf("count=%d", count))
	return c.JSON(fiber.Map{"message": fmt.Sprintf("%d inbox messages purged", count)})
}

type AdminCreateUserReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=12"`
}

func (h *Handler) AdminCreateUser(c *fiber.Ctx) error {
	var req AdminCreateUserReq
	if err := c.BodyParser(&req); err != nil || req.Email == "" {
		return c.Status(400).JSON(fiber.Map{"error": "email and password required"})
	}
	if len(req.Password) < 12 {
		return c.Status(400).JSON(fiber.Map{"error": "Password must be at least 12 characters"})
	}
	user, err := h.Service.AdminCreateUser(c.Context(), req.Email, req.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to create user"})
	}
	h.audit(c, "create_user", req.Email, "")
	return c.JSON(fiber.Map{"message": "User created", "user": user})
}

func (h *Handler) AdminGetInboxRaw(c *fiber.Ctx) error {
	id := c.Params("id")
	msgID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid message ID"})
	}
	raw, err := h.Service.AdminGetInboxRaw(c.Context(), uint(msgID))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Message not found"})
	}
	c.Set("Content-Type", "text/plain; charset=utf-8")
	return c.Send(raw)
}

type AdminAliasExpiryReq struct {
	ExpiresAt string `json:"expires_at"`
}

func (h *Handler) AdminSetAliasExpiry(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Alias ID required"})
	}
	var req AdminAliasExpiryReq
	c.BodyParser(&req)
	var expiresAt *time.Time
	if req.ExpiresAt != "" {
		t, err := time.Parse("2006-01-02", req.ExpiresAt)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid date, use YYYY-MM-DD"})
		}
		expiresAt = &t
	}
	if err := h.Service.AdminSetAliasExpiry(c.Context(), id, expiresAt); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to set alias expiry"})
	}
	return c.JSON(fiber.Map{"message": "Alias expiry updated"})
}

func (h *Handler) AdminSetAccessKeyExpiry(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Access key ID required"})
	}
	var req AdminAliasExpiryReq
	c.BodyParser(&req)
	var expiresAt *time.Time
	if req.ExpiresAt != "" {
		t, err := time.Parse("2006-01-02", req.ExpiresAt)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid date, use YYYY-MM-DD"})
		}
		expiresAt = &t
	}
	if err := h.Service.AdminSetAccessKeyExpiry(c.Context(), id, expiresAt); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to set access key expiry"})
	}
	return c.JSON(fiber.Map{"message": "Access key expiry updated"})
}

func (h *Handler) AdminGetAuditLog(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)
	entries, total, err := h.Service.AdminGetAuditLog(c.Context(), limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch audit log"})
	}
	return c.JSON(fiber.Map{"entries": entries, "total": total})
}

func (h *Handler) AdminGetSessionData(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Session ID required"})
	}
	data, err := h.Service.AdminGetSessionData(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Session not found"})
	}
	c.Set("Content-Type", "application/json; charset=utf-8")
	return c.Send(data)
}

type AdminBulkIDsReq struct {
	IDs []string `json:"ids"`
}

func (h *Handler) AdminBulkDeleteAccessKeys(c *fiber.Ctx) error {
	var req AdminBulkIDsReq
	if err := c.BodyParser(&req); err != nil || len(req.IDs) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "ids required"})
	}
	if err := h.Service.AdminBulkDeleteAccessKeys(c.Context(), req.IDs); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to bulk revoke keys"})
	}
	h.audit(c, "bulk_revoke_keys", fmt.Sprintf("%d keys", len(req.IDs)), "")
	return c.JSON(fiber.Map{"message": fmt.Sprintf("%d keys revoked", len(req.IDs))})
}

func (h *Handler) AdminBulkDeleteCredentials(c *fiber.Ctx) error {
	var req AdminBulkIDsReq
	if err := c.BodyParser(&req); err != nil || len(req.IDs) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "ids required"})
	}
	if err := h.Service.AdminBulkDeleteCredentials(c.Context(), req.IDs); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to bulk remove credentials"})
	}
	h.audit(c, "bulk_remove_passkeys", fmt.Sprintf("%d creds", len(req.IDs)), "")
	return c.JSON(fiber.Map{"message": fmt.Sprintf("%d passkeys removed", len(req.IDs))})
}

type AdminBulkExtendReq struct {
	IDs  []string `json:"ids"`
	Days int      `json:"days"`
}

func (h *Handler) AdminBulkExtendSubscriptions(c *fiber.Ctx) error {
	var req AdminBulkExtendReq
	if err := c.BodyParser(&req); err != nil || len(req.IDs) == 0 || req.Days == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "ids and days required"})
	}
	count, err := h.Service.AdminBulkExtendSubscriptions(c.Context(), req.IDs, req.Days)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to bulk extend"})
	}
	h.audit(c, "bulk_extend_subs", fmt.Sprintf("%d subs", len(req.IDs)), fmt.Sprintf("%d days", req.Days))
	return c.JSON(fiber.Map{"message": fmt.Sprintf("%d subscriptions extended by %d days", count, req.Days)})
}

func (h *Handler) AdminGetLogsDateRange(c *fiber.Ctx) error {
	logType := c.Query("type", "")
	from := c.Query("from", "")
	to := c.Query("to", "")
	limit := c.QueryInt("limit", 100)
	offset := c.QueryInt("offset", 0)
	logs, total, err := h.Service.AdminGetLogsDateRange(c.Context(), logType, from, to, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch logs"})
	}
	return c.JSON(fiber.Map{"logs": logs, "total": total})
}

func (h *Handler) AdminExportUsersEnriched(c *fiber.Ctx) error {
	users, err := h.Service.AdminExportUsersEnriched(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to export users"})
	}
	var buf bytes.Buffer
	buf.WriteString("id,email,is_active,is_admin,tier,sub_type,sub_active,active_until,created_at\n")
	for _, u := range users {
		au := ""; if u.ActiveUntil != nil { au = u.ActiveUntil.Format(time.RFC3339) };
		buf.WriteString(fmt.Sprintf("%s,%s,%t,%t,%s,%s,%t,%s,%s\n", u.ID, u.Email, u.IsActive, u.IsAdmin, u.Tier, u.SubType, u.SubActive, au, u.CreatedAt.Format(time.RFC3339)))
	}
	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", "attachment; filename=users_enriched.csv")
	return c.Send(buf.Bytes())
}

type AdminBulkDeleteMsgsReq struct {
	IDs []uint `json:"ids"`
}

func (h *Handler) AdminBulkDeleteMessages(c *fiber.Ctx) error {
	var req AdminBulkDeleteMsgsReq
	if err := c.BodyParser(&req); err != nil || len(req.IDs) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "ids required"})
	}
	if err := h.Service.AdminBulkDeleteMessages(c.Context(), req.IDs); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to bulk delete messages"})
	}
	h.audit(c, "bulk_delete_messages", fmt.Sprintf("%d msgs", len(req.IDs)), "")
	return c.JSON(fiber.Map{"message": fmt.Sprintf("%d messages deleted", len(req.IDs))})
}

type AdminPGPUploadReq struct {
	PGPKey    string `json:"pgp_key"`
	PGPInline bool   `json:"pgp_inline"`
}

func (h *Handler) AdminSetRecipientPGP(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Recipient ID required"})
	}
	var req AdminPGPUploadReq
	c.BodyParser(&req)
	if err := h.Service.AdminSetRecipientPGP(c.Context(), id, req.PGPKey, req.PGPInline); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to set PGP key"})
	}
	h.audit(c, "set_pgp_key", id, "")
	return c.JSON(fiber.Map{"message": "PGP key set"})
}

func (h *Handler) AdminGetDomainDNS(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Domain ID required"})
	}
	dns, err := h.Service.AdminGetDomainDNS(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Domain not found"})
	}
	return c.JSON(dns)
}

type AdminUserNotesReq struct {
	Notes string `json:"notes"`
}

func (h *Handler) AdminUpdateUserNotes(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "User ID required"})
	}
	var req AdminUserNotesReq
	c.BodyParser(&req)
	if err := h.Service.AdminUpdateUserNotes(c.Context(), id, req.Notes); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to update notes"})
	}
	h.audit(c, "update_user_notes", id, req.Notes)
	return c.JSON(fiber.Map{"message": "Notes updated"})
}

func (h *Handler) AdminGetSubscriptionStats(c *fiber.Ctx) error {
	active, expired, grace, err := h.Service.AdminGetSubscriptionStats(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch subscription stats"})
	}
	return c.JSON(fiber.Map{"active": active, "expired": expired, "grace_period": grace})
}

func (h *Handler) audit(c *fiber.Ctx, action, target, details string) {
	email, _ := c.Locals("admin_email").(string)
	if email == "" { email = "unknown" }
	h.Service.LogAdminAction(c.Context(), email, action, target, details)
}
