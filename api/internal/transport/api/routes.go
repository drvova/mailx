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
