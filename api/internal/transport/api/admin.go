package api

import (
	"context"

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
