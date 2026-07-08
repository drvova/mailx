package api

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/swagger"
	"ivpn.net/email/api/config"
	_ "ivpn.net/email/api/docs"
	"ivpn.net/email/api/internal/middleware/auth"
	"ivpn.net/email/api/internal/middleware/limit"
)

func (h *Handler) SetupRoutes(cfg config.APIConfig) {
	email := h.Server.Group("/v1/email")
	email.Use(auth.NewIPFilter(cfg.ApiAllowIPs))
	email.Use(auth.NewPSK(cfg.PSK))
	email.Post("", h.HandleEmail)
	email.Post("/domain/check", h.CheckDomain)

	h.Server.Use(auth.NewAPICORS(cfg))
	h.Server.Use(helmet.New())
	h.Server.Use(healthcheck.New())

	h.Server.Post("/v1/register", limiter.New(), h.Register)
	h.Server.Post("/v1/login", limit.New(5, 10*time.Minute), h.Login)
	h.Server.Post("/v1/initiatepasswordreset", limiter.New(), h.InitiatePasswordReset)
	h.Server.Put("/v1/resetpassword", limiter.New(), h.ResetPassword)
	h.Server.Put("/v1/rotatepasession", limiter.New(), h.RotatePASession)
	h.Server.Post("/v1/api/authenticate", limit.New(5, 10*time.Minute), h.Authenticate)

	h.Server.Post("/v1/register/begin", limiter.New(), h.BeginRegistration)
	h.Server.Post("/v1/register/finish", limiter.New(), h.FinishRegistration)
	h.Server.Post("/v1/login/begin", limiter.New(), h.BeginLogin)
	h.Server.Post("/v1/login/finish", limiter.New(), h.FinishLogin)

	session := h.Server.Group("/v1/pasession")
	session.Use(auth.NewIPFilter(cfg.ApiAllowIPs))
	session.Use(auth.NewPSK(cfg.PSK))
	session.Post("/add", h.AddPASession)

	api := h.Server.Group("/v1/api")
	api.Use(auth.NewAPIAuth(cfg, h.Service))
	api.Get("/aliases", h.GetAliases)
	api.Post("/alias", limiter.New(), h.PostAlias)
	api.Put("/alias/:id", h.UpdateAlias)
	api.Delete("/alias/:id", h.DeleteAlias)
	api.Get("/defaults", h.GetDefaults)
	api.Post("/logout", h.ApiLogout)

	v1 := h.Server.Group("/v1")
	v1.Use(auth.New(cfg, h.Cache, h.Service))

	v1.Post("/register/add", limiter.New(), h.AddPasskey)
	v1.Post("/register/add/finish", limiter.New(), h.FinishAddPasskey)
	v1.Post("/user/sendotp", limit.New(5, 10*time.Minute), h.SendUserOTP)
	v1.Post("/user/activate", limiter.New(), h.Activate)
	v1.Post("/user/logout", h.Logout)
	v1.Post("/user/delete/request", limit.New(5, 10*time.Minute), h.DeleteUserRequest)
	v1.Post("/user/delete", limit.New(5, 10*time.Minute), h.DeleteUser)
	v1.Get("/user", h.GetUser)
	v1.Get("/user/stats", h.GetUserStats)
	v1.Get("/user/credentials", h.GetCredentials)
	v1.Delete("/user/credential/:id", h.DeleteCredential)
	v1.Put("/user/changepassword", limit.New(5, 10*time.Minute), h.ChangePassword)
	v1.Put("/user/changeemail", limit.New(5, 10*time.Minute), h.ChangeEmail)
	v1.Put("/user/totp/enable", limit.New(5, 10*time.Minute), h.TotpEnable)
	v1.Put("/user/totp/enable/confirm", limit.New(5, 10*time.Minute), h.TotpEnableConfirm)
	v1.Put("/user/totp/disable", limit.New(5, 10*time.Minute), h.TotpDisable)

	v1.Get("/sub", h.GetSubscription)
	v1.Put("/sub/update", limiter.New(), h.UpdateSubscription)

	v1.Get("/settings", h.GetSettings)
	v1.Get("/defaults", h.GetDefaults)
	v1.Put("/settings", h.UpdateSettings)

	v1.Get("/recipient/:id", h.GetRecipient)
	v1.Get("/recipients", h.GetRecipients)
	v1.Post("/recipient", limit.New(5, 10*time.Minute), h.PostRecipient)
	v1.Put("/recipient", h.UpdateRecipient)
	v1.Post("/recipient/sendotp/:id", limit.New(5, 10*time.Minute), h.SendRecipientOTP)
	v1.Post("/recipient/activate/:id", limit.New(5, 10*time.Minute), h.ActivateRecipient)
	v1.Put("/recipient/delete/:id", h.DeleteRecipient)

	v1.Get("/alias/:id", h.GetAlias)
	v1.Get("/alias/:id/inbox", h.GetInboxMessages)
	v1.Get("/inbox/unread", h.GetInboxUnread)
	v1.Get("/inbox/message/:id", h.GetInboxMessage)
	v1.Delete("/inbox/message/:id", h.DeleteInboxMessage)
	v1.Get("/aliases", h.GetAliases)
	v1.Get("/aliases/export", h.ExportAliases)
	v1.Post("/alias", limiter.New(), h.PostAlias)
	v1.Put("/alias/:id", h.UpdateAlias)
	v1.Delete("/alias/:id", h.DeleteAlias)

	v1.Get("/logs", h.GetLogs)
	v1.Delete("/logs", h.DeleteLogs)
	v1.Get("/log/file/:id", h.GetLogFile)

	v1.Get("/accesskeys", h.GetAccessKeys)
	v1.Post("/accesskeys", limiter.New(), h.PostAccessKey)
	v1.Delete("/accesskeys/:id", h.DeleteAccessKey)

	v1.Get("/domains", h.GetDomains)
	v1.Get("/domains/dns-config", h.GetDNSConfig)
	v1.Post("/domain", limiter.New(), h.PostDomain)
	v1.Put("/domain/:id", h.UpdateDomain)
	v1.Delete("/domain/:id", h.DeleteDomain)
	v1.Post("/domain/:id/verify-dns", h.VerifyDomainDNSRecords)

	// Plans - public list, admin CRUD
	v1.Get("/plans", h.GetPlans)
	admin := v1.Group("/admin", auth.NewAdminGuard(h.Service))
	admin.Get("/plans", h.GetAllPlans)
	admin.Post("/plan", h.CreatePlan)
	admin.Put("/plan/:id", h.UpdatePlan)
	admin.Delete("/plan/:id", h.DeletePlan)

	// Admin user management
	admin.Get("/users", h.AdminGetUsers)
	admin.Get("/stats", h.AdminGetStats)
	admin.Get("/logs", h.AdminGetLogs)
	admin.Put("/user", h.AdminUpdateUser)
	admin.Delete("/user/:id", h.AdminDeleteUser)
	admin.Post("/user/assign-plan", h.AdminAssignPlan)

	// Admin alias moderation
	admin.Get("/aliases", h.AdminGetAliases)
	admin.Delete("/alias/:id", h.AdminDeleteAlias)
	admin.Put("/alias/:id/toggle", h.AdminToggleAlias)

	// Admin domain moderation
	admin.Get("/domains", h.AdminGetDomains)
	admin.Delete("/domain/:id", h.AdminDeleteDomain)
	admin.Put("/domain/:id/toggle", h.AdminToggleDomain)

	// Admin recipient moderation
	admin.Get("/recipients", h.AdminGetRecipients)
	admin.Delete("/recipient/:id", h.AdminDeleteRecipient)

	// Admin log filtering + user search + user detail
	admin.Get("/logs/filter", h.AdminGetLogsFiltered)
	admin.Get("/users/search", h.AdminSearchUsers)
	admin.Get("/user/:id/detail", h.AdminGetUserDetail)

	// Admin access key moderation
	admin.Get("/accesskeys", h.AdminGetAccessKeys)
	admin.Delete("/accesskey/:id", h.AdminDeleteAccessKey)

	// Admin session moderation
	admin.Get("/sessions", h.AdminGetSessions)
	admin.Delete("/session/:id", h.AdminDeleteSession)
	admin.Delete("/user/:id/sessions", h.AdminForceLogout)

	// Admin credential (passkey) moderation
	admin.Get("/credentials", h.AdminGetCredentials)
	admin.Delete("/credential/:id", h.AdminDeleteCredential)

	// Admin subscription override
	admin.Put("/subscription", h.AdminUpdateSubscription)

	// Admin bulk operations
	admin.Post("/users/bulk", h.AdminBulkUpdateUsers)

	// Admin inbox moderation
	admin.Get("/inbox", h.AdminGetInboxMessages)
	admin.Delete("/inbox/message/:id", h.AdminDeleteInboxMessage)
	admin.Delete("/inbox/purge/:id", h.AdminPurgeInbox)

	// Admin TOTP and password management
	admin.Delete("/user/:id/totp", h.AdminDisableTotp)
	admin.Post("/user/reset-password", h.AdminResetPassword)

	// Admin settings override
	admin.Get("/user/:id/settings", h.AdminGetSettings)
	admin.Put("/user/settings", h.AdminUpdateSettings)

	// Admin CSV export
	admin.Get("/export/users", h.AdminExportUsers)
	admin.Get("/export/aliases", h.AdminExportAliases)

	// Admin subscription management
	admin.Get("/subscriptions", h.AdminGetSubscriptions)
	admin.Delete("/subscription/:id", h.AdminDeleteSubscription)

	// Admin bulk delete for aliases/domains/recipients
	admin.Post("/aliases/bulk-delete", h.AdminBulkDeleteAliases)
	admin.Post("/domains/bulk-delete", h.AdminBulkDeleteDomains)
	admin.Post("/recipients/bulk-delete", h.AdminBulkDeleteRecipients)

	// Admin system health
	admin.Get("/system/tables", h.AdminGetTableSizes)
	admin.Get("/system/recent-signups", h.AdminGetRecentSignups)

	// Admin domain verification override
	admin.Put("/domain/:id/verify", h.AdminVerifyDomain)

	// Admin impersonation
	admin.Post("/user/:id/impersonate", h.AdminImpersonate)

	// Admin search for keys/sessions/inbox
	admin.Get("/accesskeys/search", h.AdminSearchAccessKeys)
	admin.Get("/sessions/search", h.AdminSearchSessions)
	admin.Get("/inbox/search", h.AdminSearchInbox)

	// Admin message log + user stats
	admin.Get("/messages", h.AdminGetMessages)
	admin.Get("/user/:id/stats", h.AdminGetUserStats)

	// Admin log search (text + type)
	admin.Get("/logs/search", h.AdminSearchLogs)

	// Admin recipient toggle
	admin.Put("/recipient/:id/toggle", h.AdminToggleRecipient)

	// Admin domain search
	admin.Get("/domains/search", h.AdminSearchDomains)

	// Admin CSV export for recipients and subscriptions
	admin.Get("/export/recipients", h.AdminExportRecipients)
	admin.Get("/export/subscriptions", h.AdminExportSubscriptions)

	// Admin user email change
	admin.Put("/user/email", h.AdminChangeEmail)

	// Admin CSV export for domains and logs
	admin.Get("/export/domains", h.AdminExportDomains)
	admin.Get("/export/logs", h.AdminExportLogs)

	// Admin bulk delete users
	admin.Post("/users/bulk-delete", h.AdminBulkDeleteUsers)

	// Admin config viewer (redacted)
	admin.Get("/system/config", h.AdminGetConfig)

	// Admin message search
	admin.Get("/messages/search", h.AdminSearchMessages)

	// Admin recipient PGP management
	admin.Put("/recipient/:id/pgp", h.AdminToggleRecipientPGP)
	admin.Delete("/recipient/:id/pgp", h.AdminRemoveRecipientPGPKey)

	// Admin alias/domain edit
	admin.Put("/alias/:id", h.AdminUpdateAlias)
	admin.Put("/domain/:id", h.AdminUpdateDomain)

	// Admin inbox mark as read
	admin.Put("/inbox/message/:id/read", h.AdminMarkInboxRead)

	// Admin paginated users
	admin.Get("/users/paginated", h.AdminGetUsersPaginated)

	// Admin create resources for users
	admin.Post("/recipient", h.AdminCreateRecipient)
	admin.Post("/domain", h.AdminCreateDomain)

	// Admin CSV exports for inbox and messages
	admin.Get("/export/inbox", h.AdminExportInbox)
	admin.Get("/export/messages", h.AdminExportMessages)

	// Admin alias creation, recipient edit, log deletion, inbox bulk delete, subscription extend
	admin.Post("/alias", h.AdminCreateAlias)
	admin.Put("/recipient/:id", h.AdminUpdateRecipient)
	admin.Delete("/log/:id", h.AdminDeleteLog)
	admin.Post("/inbox/bulk-delete", h.AdminBulkDeleteInbox)
	admin.Post("/subscription/:id/extend", h.AdminExtendSubscription)

	// Admin access key creation, ownership transfer, purge operations
	admin.Post("/accesskey", h.AdminCreateAccessKey)
	admin.Post("/alias/:id/transfer", h.AdminTransferAlias)
	admin.Post("/domain/:id/transfer", h.AdminTransferDomain)
	admin.Post("/logs/purge", h.AdminPurgeLogs)
	admin.Delete("/inbox/purge-all", h.AdminPurgeAllInbox)

	// Admin user creation, inbox raw view, alias expiry
	admin.Post("/user/create", h.AdminCreateUser)
	admin.Get("/inbox/message/:id/raw", h.AdminGetInboxRaw)
	admin.Put("/alias/:id/expiry", h.AdminSetAliasExpiry)
	admin.Put("/accesskey/:id/expiry", h.AdminSetAccessKeyExpiry)

	// Admin audit log
	admin.Get("/audit", h.AdminGetAuditLog)

	// Session data viewer
	admin.Get("/session/:id/data", h.AdminGetSessionData)

	// Date range logs
	admin.Get("/logs/date-range", h.AdminGetLogsDateRange)

	// Bulk operations on keys, credentials, subscriptions
	admin.Post("/accesskeys/bulk-revoke", h.AdminBulkDeleteAccessKeys)
	admin.Post("/credentials/bulk-remove", h.AdminBulkDeleteCredentials)
	admin.Post("/subscriptions/bulk-extend", h.AdminBulkExtendSubscriptions)

	// Enriched user export with subscription data
	admin.Get("/export/users-enriched", h.AdminExportUsersEnriched)

	// Bulk message delete, PGP upload, domain DNS, user notes, subscription stats
	admin.Post("/messages/bulk-delete", h.AdminBulkDeleteMessages)
	admin.Put("/recipient/:id/pgp-key", h.AdminSetRecipientPGP)
	admin.Get("/domain/:id/dns", h.AdminGetDomainDNS)
	admin.Put("/user/:id/notes", h.AdminUpdateUserNotes)
	admin.Get("/subscriptions/stats", h.AdminGetSubscriptionStats)

	// Daily activity, plan distribution, domain health, global search
	admin.Get("/daily-activity", h.AdminGetDailyActivity)
	admin.Get("/plan-distribution", h.AdminGetPlanDistribution)
	admin.Get("/domain-health", h.AdminGetDomainHealth)
	admin.Get("/global-search", h.AdminGlobalUserSearch)

	// User last active, inactive users, catch-all toggle, full user data export
	admin.Get("/user/:id/last-active", h.AdminGetUserLastActive)
	admin.Get("/inactive-users", h.AdminGetInactiveUsers)
	admin.Put("/alias/:id/catch-all", h.AdminToggleAliasCatchAll)
	admin.Get("/user/:id/export-data", h.AdminExportUserData)

	// Session cleanup and domain stats
	admin.Delete("/sessions/expired", h.AdminPurgeExpiredSessions)
	admin.Get("/domain-stats", h.AdminGetDomainWithAliasCounts)

	// Billing - Oxapay checkout + webhook
	v1.Post("/billing/checkout", h.CreateCheckoutSession)
	h.Server.Post("/v1/billing/webhook", h.StripeWebhook)

	docs := h.Server.Group("/docs")
	docs.Use(auth.NewBasicAuth(cfg))
	docs.Get("/*", swagger.HandlerDefault)

	// Anonymous endpoint hit counters — no PII, no cookies, no client JS.
	// Protected by the same basic-auth as docs.
	metrics := h.Server.Group("/metrics")
	metrics.Use(auth.NewBasicAuth(cfg))
	metrics.Get("/endpoints", h.Metrics.HandleMetrics)
}
