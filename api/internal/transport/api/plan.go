package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"ivpn.net/email/api/internal/middleware/auth"
	"ivpn.net/email/api/internal/model"
)

type PlanService interface {
	GetActivePlans(context.Context) ([]model.Plan, error)
	GetAllPlans(context.Context) ([]model.Plan, error)
	GetPlan(context.Context, string) (model.Plan, error)
	CreatePlan(context.Context, model.Plan) error
	UpdatePlan(context.Context, model.Plan) error
	DeletePlan(context.Context, string) error
}

type BillingService interface {
	CreateCheckout(context.Context, string, string) (string, error)
	VerifyWebhook([]byte, string) bool
	HandleWebhook(context.Context, []byte) error
}

func (h *Handler) GetPlans(c *fiber.Ctx) error {
	plans, err := h.Service.GetActivePlans(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch plans"})
	}
	return c.JSON(plans)
}

func (h *Handler) GetAllPlans(c *fiber.Ctx) error {
	plans, err := h.Service.GetAllPlans(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch plans"})
	}
	return c.JSON(plans)
}

func (h *Handler) CreatePlan(c *fiber.Ctx) error {
	var req PlanReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	plan := model.Plan{
		Name:              req.Name,
		DisplayName:       req.DisplayName,
		PriceCents:        req.PriceCents,
		Currency:          req.Currency,
		Interval:          req.Interval,
		MaxRecipients:     req.MaxRecipients,
		MaxCredentials:    req.MaxCredentials,
		MaxDailyAliases:   req.MaxDailyAliases,
		MaxDailySendReply: req.MaxDailySendReply,
		MaxSessions:       req.MaxSessions,
		IsActive:          true,
		SortOrder:         req.SortOrder,
	}

	if err := h.Service.CreatePlan(c.Context(), plan); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to create plan"})
	}

	return c.Status(201).JSON(plan)
}

func (h *Handler) UpdatePlan(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Plan ID required"})
	}

	plan, err := h.Service.GetPlan(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Plan not found"})
	}

	var req PlanReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	plan.Name = req.Name
	plan.DisplayName = req.DisplayName
	plan.PriceCents = req.PriceCents
	plan.Currency = req.Currency
	plan.Interval = req.Interval
	plan.MaxRecipients = req.MaxRecipients
	plan.MaxCredentials = req.MaxCredentials
	plan.MaxDailyAliases = req.MaxDailyAliases
	plan.MaxDailySendReply = req.MaxDailySendReply
	plan.MaxSessions = req.MaxSessions
	plan.SortOrder = req.SortOrder

	if err := h.Service.UpdatePlan(c.Context(), plan); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to update plan"})
	}

	return c.JSON(plan)
}

func (h *Handler) DeletePlan(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Plan ID required"})
	}

	if err := h.Service.DeletePlan(c.Context(), id); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to delete plan"})
	}

	return c.JSON(fiber.Map{"message": "Plan deactivated"})
}

func (h *Handler) CreateCheckoutSession(c *fiber.Ctx) error {
	userID := auth.GetUserID(c)

	var req CheckoutReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	url, err := h.Service.CreateCheckout(c.Context(), userID, req.PlanID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"url": url})
}

func (h *Handler) StripeWebhook(c *fiber.Ctx) error {
	hmacHeader := c.Get("HMAC")
	rawBody := c.Body()

	if !h.Service.VerifyWebhook(rawBody, hmacHeader) {
		return c.Status(401).SendString("Invalid signature")
	}

	if err := h.Service.HandleWebhook(c.Context(), rawBody); err != nil {
		return c.Status(500).SendString("Error")
	}

	return c.SendString("OK")
}
