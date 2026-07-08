package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"ivpn.net/email/api/internal/client/oxapay"
)

type OxapayClient interface {
	CreateInvoice(req oxapay.InvoiceReq) (*oxapay.InvoiceResp, error)
	GetPaymentInfo(trackID string) (*oxapay.PaymentInfo, error)
}

func (s *Service) CreateCheckout(ctx context.Context, userID string, planID string) (string, error) {
	plan, err := s.Store.GetPlan(ctx, planID)
	if err != nil {
		return "", fmt.Errorf("plan not found")
	}

	if plan.PriceCents == 0 {
		return "", fmt.Errorf("free plan does not require checkout")
	}

	user, err := s.Store.GetUser(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	amount := float64(plan.PriceCents) / 100.0
	orderID := fmt.Sprintf("plan-%s-user-%s", plan.ID, userID)

	resp, err := s.Oxapay.CreateInvoice(oxapay.InvoiceReq{
		Amount:      amount,
		Currency:    plan.Currency,
		Lifetime:    60,
		CallbackURL: s.Cfg.Oxapay.WebhookURL,
		ReturnURL:   s.Cfg.Oxapay.ReturnURL,
		Email:       user.Email,
		OrderID:     orderID,
		Description: fmt.Sprintf("FreeTheMail %s plan - %s", plan.DisplayName, plan.Interval),
	})
	if err != nil {
		log.Printf("error creating oxapay invoice: %s", err.Error())
		return "", err
	}

	return resp.Data.PaymentURL, nil
}

type WebhookData struct {
	TrackID string  `json:"track_id"`
	Status  string  `json:"status"`
	Type    string  `json:"type"`
	Amount  float64 `json:"amount"`
	OrderID string  `json:"order_id"`
	Email   string  `json:"email"`
}

func (s *Service) VerifyWebhook(rawBody []byte, hmacHeader string) bool {
	mac := hmac.New(sha512.New, []byte(s.Cfg.Oxapay.MerchantKey))
	mac.Write(rawBody)
	expected := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expected), []byte(hmacHeader))
}

func (s *Service) HandleWebhook(ctx context.Context, rawBody []byte) error {
	var data WebhookData
	if err := json.Unmarshal(rawBody, &data); err != nil {
		return err
	}

	if data.Status != "Paid" {
		return nil
	}

	userID, planID, err := parseOrderID(data.OrderID)
	if err != nil {
		log.Printf("error parsing order_id from webhook: %s", err.Error())
		return nil
	}

	plan, err := s.Store.GetPlan(ctx, planID)
	if err != nil {
		log.Printf("error getting plan from webhook: %s", err.Error())
		return nil
	}

	sub, err := s.Store.GetSubscription(ctx, userID)
	if err != nil {
		log.Printf("error getting subscription from webhook: %s", err.Error())
		return nil
	}

	var activeUntil time.Time
	switch plan.Interval {
	case "monthly":
		activeUntil = time.Now().AddDate(0, 1, 0)
	case "yearly":
		activeUntil = time.Now().AddDate(1, 0, 0)
	default:
		activeUntil = time.Now().AddDate(100, 0, 0)
	}

	sub.PlanID = &plan.ID
	sub.Tier = plan.Name
	sub.Type = "paid"
	sub.ActiveUntil = activeUntil
	sub.IsActive = true

	return s.Store.UpdateSubscription(ctx, sub)
}

func parseOrderID(orderID string) (userID string, planID string, err error) {
	var p, u string
	_, err = fmt.Sscanf(orderID, "plan-%s-user-%s", &p, &u)
	if err != nil {
		return "", "", err
	}
	return u, p, nil
}
