package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"ivpn.net/email/api/internal/middleware/auth"
	"ivpn.net/email/api/internal/model"
)

var (
	UpdateSubscriptionSuccess = "Subscription updated successfully."
	AddSubscriptionSuccess    = "Subscription added successfully."
)

type SubscriptionService interface {
	GetSubscription(context.Context, string) (model.Subscription, error)
}

// @Summary Get subscription
// @Description Get subscription
// @Tags subscription
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.Subscription
// @Failure 400 {object} ErrorRes
// @Router /sub [get]
func (h *Handler) GetSubscription(c *fiber.Ctx) error {
	userID := auth.GetUserID(c)

	sub, err := h.Service.GetSubscription(c.Context(), userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(sub)
}
