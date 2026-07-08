package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"ivpn.net/email/api/internal/model"
)

var (
	ErrGetSubscription    = errors.New("Unable to retrieve subscription by user ID.")
	ErrAddSubscription    = errors.New("Unable to add subscription.")
	ErrPostSubscription   = errors.New("Unable to create subscription.")
	ErrUpdateSubscription = errors.New("Unable to update subscription.")
	ErrDeleteSubscription = errors.New("Unable to delete subscription.")
)

type SubscriptionStore interface {
	GetSubscription(context.Context, string) (model.Subscription, error)
	PostSubscription(context.Context, model.Subscription) error
	UpdateSubscription(context.Context, model.Subscription) error
	DeleteSubscription(context.Context, string) error
}

func (s *Service) GetSubscription(ctx context.Context, userID string) (model.Subscription, error) {
	sub, err := s.Store.GetSubscription(ctx, userID)
	if err != nil {
		return model.Subscription{}, ErrGetSubscription
	}

	sub.Status = sub.GetStatus()
	sub.Outage = sub.IsOutage()

	return sub, nil
}

func (s *Service) AddSubscription(ctx context.Context, subscription model.Subscription, activeUntil string) error {
	err := s.Cache.Set(ctx, "sub_"+subscription.ID, activeUntil, s.Cfg.Service.OTPExpiration)
	if err != nil {
		log.Printf("error adding subscription: %s", err.Error())
		return ErrAddSubscription
	}

	return nil
}

func (s *Service) DeleteSubscription(ctx context.Context, userID string) error {
	err := s.Store.DeleteSubscription(ctx, userID)
	if err != nil {
		log.Printf("error deleting subscription: %s", err.Error())
		return ErrDeleteSubscription
	}

	return nil
}

func (s *Service) CreateSelfHostedSubscription(ctx context.Context, userID string) error {
	sub := model.Subscription{
		UserID:      userID,
		Type:        "self",
		ActiveUntil: time.Now().AddDate(100, 0, 0),
		IsActive:    true,
		Tier:        "self-hosted",
	}
	sub.ID = uuid.New().String()

	// Assign free plan if one exists
	plans, perr := s.Store.GetActivePlans(ctx)
	if perr == nil {
		for _, p := range plans {
			if p.PriceCents == 0 {
				sub.PlanID = &p.ID
				sub.Tier = p.Name
				break
			}
		}
	}

	err := s.Store.PostSubscription(ctx, sub)
	if err != nil {
		log.Printf("error creating subscription: %s", err.Error())
		return ErrPostSubscription
	}

	return nil
}
