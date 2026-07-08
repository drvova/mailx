package service

import (
	"context"
	"log"

	"ivpn.net/email/api/internal/model"
)

type PlanStore interface {
	GetPlan(context.Context, string) (model.Plan, error)
	GetPlanByName(context.Context, string) (model.Plan, error)
	GetActivePlans(context.Context) ([]model.Plan, error)
	GetAllPlans(context.Context) ([]model.Plan, error)
	PostPlan(context.Context, model.Plan) error
	UpdatePlan(context.Context, model.Plan) error
	DeletePlan(context.Context, string) error
}

func (s *Service) GetActivePlans(ctx context.Context) ([]model.Plan, error) {
	return s.Store.GetActivePlans(ctx)
}

func (s *Service) GetAllPlans(ctx context.Context) ([]model.Plan, error) {
	return s.Store.GetAllPlans(ctx)
}

func (s *Service) GetPlan(ctx context.Context, id string) (model.Plan, error) {
	return s.Store.GetPlan(ctx, id)
}

func (s *Service) CreatePlan(ctx context.Context, plan model.Plan) error {
	err := s.Store.PostPlan(ctx, plan)
	if err != nil {
		log.Printf("error creating plan: %s", err.Error())
	}
	return err
}

func (s *Service) UpdatePlan(ctx context.Context, plan model.Plan) error {
	err := s.Store.UpdatePlan(ctx, plan)
	if err != nil {
		log.Printf("error updating plan: %s", err.Error())
	}
	return err
}

func (s *Service) DeletePlan(ctx context.Context, id string) error {
	return s.Store.DeletePlan(ctx, id)
}
