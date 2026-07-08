package service

import (
	"context"
	"log"

	"ivpn.net/email/api/internal/model"
)

type SystemStats struct {
	TotalUsers   int64 `json:"total_users"`
	ActiveUsers  int64 `json:"active_users"`
	TotalAliases int64 `json:"total_aliases"`
	TotalDomains int64 `json:"total_domains"`
	TotalLogs    int64 `json:"total_logs"`
	ActivePlans  int   `json:"active_plans"`
}

type AdminStore interface {
	GetAllUsers(context.Context) ([]model.User, error)
	GetUserCount(context.Context) (int64, error)
	GetActiveUserCount(context.Context) (int64, error)
	GetAllSubscriptions(context.Context) ([]model.Subscription, error)
	GetAllLogs(context.Context) ([]model.Log, error)
	GetLogCount(context.Context) (int64, error)
	GetAliasCountAll(context.Context) (int64, error)
	GetDomainCountAll(context.Context) (int64, error)
	AdminUpdateUser(context.Context, model.User) error
	AdminDeleteUser(context.Context, string) error
	AdminAssignPlan(context.Context, string, string, string) error
}

func (s *Service) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.Store.GetAllUsers(ctx)
}

func (s *Service) GetSystemStats(ctx context.Context) (any, error) {
	totalUsers, _ := s.Store.GetUserCount(ctx)
	activeUsers, _ := s.Store.GetActiveUserCount(ctx)
	totalAliases, _ := s.Store.GetAliasCountAll(ctx)
	totalDomains, _ := s.Store.GetDomainCountAll(ctx)
	totalLogs, _ := s.Store.GetLogCount(ctx)
	plans, _ := s.Store.GetActivePlans(ctx)

	return SystemStats{
		TotalUsers:   totalUsers,
		ActiveUsers:  activeUsers,
		TotalAliases: totalAliases,
		TotalDomains: totalDomains,
		TotalLogs:    totalLogs,
		ActivePlans:  len(plans),
	}, nil
}

func (s *Service) GetAllLogs(ctx context.Context) ([]model.Log, error) {
	return s.Store.GetAllLogs(ctx)
}

func (s *Service) AdminUpdateUser(ctx context.Context, user model.User) error {
	err := s.Store.AdminUpdateUser(ctx, user)
	if err != nil {
		log.Printf("error admin updating user: %s", err.Error())
	}
	return err
}

func (s *Service) AdminDeleteUser(ctx context.Context, userID string) error {
	err := s.Store.AdminDeleteUser(ctx, userID)
	if err != nil {
		log.Printf("error admin deleting user: %s", err.Error())
	}
	return err
}

func (s *Service) AdminAssignPlan(ctx context.Context, userID string, planID string) error {
	plan, err := s.Store.GetPlan(ctx, planID)
	if err != nil {
		return err
	}
	return s.Store.AdminAssignPlan(ctx, userID, planID, plan.Name)
}
