package api

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"ivpn.net/email/api/internal/middleware/auth"
	"ivpn.net/email/api/internal/model"
)

var DeleteInboxMessageSuccess = "Message deleted successfully."

type InboxService interface {
	GetInboxMessages(context.Context, string, string) ([]model.InboxMessage, error)
	GetInboxUnreadCount(context.Context, string) (int64, error)
	GetRenderedInboxMessage(context.Context, uint, string) (model.RenderedMessage, error)
	DeleteInboxMessage(context.Context, uint, string) error
}

// @Summary Get inbox messages
// @Description Get stored messages for an inbox alias
// @Tags inbox
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Alias ID"
// @Success 200 {array} model.InboxMessage
// @Failure 400 {object} ErrorRes
// @Router /alias/{id}/inbox [get]
func (h *Handler) GetInboxMessages(c *fiber.Ctx) error {
	userID := auth.GetUserID(c)
	messages, err := h.Service.GetInboxMessages(c.Context(), c.Params("id"), userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(messages)
}

// @Summary Get inbox unread count
// @Description Get the count of unread inbox messages for the authenticated user
// @Tags inbox
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]int64
// @Failure 400 {object} ErrorRes
// @Router /inbox/unread [get]
func (h *Handler) GetInboxUnread(c *fiber.Ctx) error {
	userID := auth.GetUserID(c)
	count, err := h.Service.GetInboxUnreadCount(c.Context(), userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"unread": count,
	})
}

// @Summary Get inbox message
// @Description Get a parsed, sanitized inbox message by ID
// @Tags inbox
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Message ID"
// @Success 200 {object} model.RenderedMessage
// @Failure 400 {object} ErrorRes
// @Router /inbox/message/{id} [get]
func (h *Handler) GetInboxMessage(c *fiber.Ctx) error {
	userID := auth.GetUserID(c)
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": ErrInvalidRequest,
		})
	}

	message, err := h.Service.GetRenderedInboxMessage(c.Context(), uint(id), userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(message)
}

// @Summary Delete inbox message
// @Description Delete an inbox message by ID
// @Tags inbox
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Message ID"
// @Success 200 {object} SuccessRes
// @Failure 400 {object} ErrorRes
// @Router /inbox/message/{id} [delete]
func (h *Handler) DeleteInboxMessage(c *fiber.Ctx) error {
	userID := auth.GetUserID(c)
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": ErrInvalidRequest,
		})
	}

	err = h.Service.DeleteInboxMessage(c.Context(), uint(id), userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": DeleteInboxMessageSuccess,
	})
}
