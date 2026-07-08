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

func (s *Service) GetAllAliasesAdmin(ctx context.Context, limit, offset int, search string) ([]model.Alias, int64, error) {
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	return s.Store.GetAllAliasesAdmin(ctx, limit, offset, search)
}

func (s *Service) AdminDeleteAlias(ctx context.Context, aliasID string) error {
	return s.Store.AdminDeleteAlias(ctx, aliasID)
}

func (s *Service) AdminToggleAlias(ctx context.Context, aliasID string, enabled bool) error {
	return s.Store.AdminToggleAlias(ctx, aliasID, enabled)
}

func (s *Service) GetAllDomainsAdmin(ctx context.Context) ([]model.Domain, error) {
	return s.Store.GetAllDomainsAdmin(ctx)
}

func (s *Service) AdminDeleteDomain(ctx context.Context, domainID string) error {
	return s.Store.AdminDeleteDomain(ctx, domainID)
}

func (s *Service) AdminToggleDomain(ctx context.Context, domainID string, enabled bool) error {
	return s.Store.AdminToggleDomain(ctx, domainID, enabled)
}

func (s *Service) GetAllRecipientsAdmin(ctx context.Context, limit, offset int, search string) ([]model.Recipient, int64, error) {
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	return s.Store.GetAllRecipientsAdmin(ctx, limit, offset, search)
}

func (s *Service) AdminDeleteRecipient(ctx context.Context, recipientID string) error {
	return s.Store.AdminDeleteRecipient(ctx, recipientID)
}

func (s *Service) GetLogsFiltered(ctx context.Context, logType string, limit, offset int) ([]model.Log, int64, error) {
	if limit <= 0 || limit > 200 {
		limit = 100
	}
	return s.Store.GetLogsFiltered(ctx, logType, limit, offset)
}

func (s *Service) AdminSearchUsers(ctx context.Context, search string, limit, offset int) ([]model.User, int64, error) {
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	return s.Store.AdminSearchUsers(ctx, search, limit, offset)
}

func (s *Service) AdminGetUserDetail(ctx context.Context, userID string) (model.User, model.Subscription, []model.Alias, []model.Recipient, []model.Domain, error) {
	return s.Store.AdminGetUserDetail(ctx, userID)
}
