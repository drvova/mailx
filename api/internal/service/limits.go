package service

import (
	"context"
	"log"
)

type Limits struct {
	MaxRecipients     int
	MaxCredentials    int
	MaxDailyAliases   int
	MaxDailySendReply int
	MaxSessions       int
}

func (s *Service) GetLimits(ctx context.Context, userID string) Limits {
	sub, err := s.GetSubscription(ctx, userID)
	if err != nil {
		return s.envLimits()
	}

	if sub.Tier == "self-hosted" || sub.PlanID == nil || *sub.PlanID == "" {
		return s.envLimits()
	}

	plan, err := s.Store.GetPlan(ctx, *sub.PlanID)
	if err != nil {
		log.Printf("error getting plan for limits: %s", err.Error())
		return s.envLimits()
	}

	return Limits{
		MaxRecipients:     plan.MaxRecipients,
		MaxCredentials:    plan.MaxCredentials,
		MaxDailyAliases:   plan.MaxDailyAliases,
		MaxDailySendReply: plan.MaxDailySendReply,
		MaxSessions:       plan.MaxSessions,
	}
}

func (s *Service) envLimits() Limits {
	return Limits{
		MaxRecipients:     s.Cfg.Service.MaxRecipients,
		MaxCredentials:    s.Cfg.Service.MaxCredentials,
		MaxDailyAliases:   s.Cfg.Service.MaxDailyAliases,
		MaxDailySendReply: s.Cfg.Service.MaxDailySendReply,
		MaxSessions:       s.Cfg.Service.MaxSessions,
	}
}
