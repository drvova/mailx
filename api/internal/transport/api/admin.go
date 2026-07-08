package api

import (
	"bytes"
	"context"
	"fmt"
	"strconv"

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
