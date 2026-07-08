package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"ivpn.net/email/api/internal/model"
	"ivpn.net/email/api/internal/utils"
)

type SystemStats struct {
	TotalUsers          int64 `json:"total_users"`
	ActiveUsers         int64 `json:"active_users"`
	SuspendedUsers      int64 `json:"suspended_users"`
	AdminUsers          int64 `json:"admin_users"`
	TotalAliases        int64 `json:"total_aliases"`
	TotalDomains        int64 `json:"total_domains"`
	TotalRecipients     int64 `json:"total_recipients"`
	TotalLogs           int64 `json:"total_logs"`
	TotalInboxMessages  int64 `json:"total_inbox_messages"`
	TotalSubscriptions  int64 `json:"total_subscriptions"`
	ActiveSubscriptions int64 `json:"active_subscriptions"`
	ActivePlans         int   `json:"active_plans"`
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
	GetAllAccessKeysAdmin(context.Context, int, int) ([]model.AccessKey, int64, error)
	AdminDeleteAccessKey(context.Context, string) error
	GetAllSessionsAdmin(context.Context, int, int) ([]model.Session, int64, error)
	AdminDeleteSession(context.Context, string) error
	AdminDeleteSessionsByUserID(context.Context, string) error
	GetAllCredentialsAdmin(context.Context, int, int) ([]model.Credential, int64, error)
	AdminDeleteCredential(context.Context, string) error
	AdminUpdateSubscription(context.Context, string, string, bool, string) error
	GetSuspendedUserCount(context.Context) (int64, error)
	GetAdminCount(context.Context) (int64, error)
	AdminBulkUpdateUsers(context.Context, []string, bool) error
	GetAllInboxMessagesAdmin(context.Context, int, int) ([]model.InboxMessage, int64, error)
	AdminDeleteInboxMessage(context.Context, uint) error
	AdminPurgeInboxByUser(context.Context, string) error
	AdminDisableTotp(context.Context, string) error
	AdminResetPassword(context.Context, string, string) error
	AdminGetSettings(context.Context, string) (model.Settings, error)
	AdminUpdateSettings(context.Context, string, map[string]interface{}) error
	GetInboxMessageCount(context.Context) (int64, error)
	GetRecipientCountAll(context.Context) (int64, error)
	GetSubscriptionCount(context.Context) (int64, error)
	GetActiveSubscriptionCount(context.Context) (int64, error)
	AdminExportUsers(context.Context) ([]model.User, error)
	AdminExportAliases(context.Context) ([]model.Alias, error)
	GetAllSubscriptionsAdmin(context.Context, int, int, string) ([]model.Subscription, int64, error)
	AdminDeleteSubscription(context.Context, string) error
	AdminImpersonate(context.Context, string) (model.User, error)
	AdminBulkDeleteAliases(context.Context, []string) error
	AdminBulkDeleteDomains(context.Context, []string) error
	AdminBulkDeleteRecipients(context.Context, []string) error
	GetTableSizes(context.Context) (map[string]int64, error)
	GetRecentSignups(context.Context, int) ([]model.User, error)
	SearchAccessKeys(context.Context, string, int, int) ([]model.AccessKey, int64, error)
	SearchSessions(context.Context, string, int, int) ([]model.Session, int64, error)
	SearchInboxMessages(context.Context, string, int, int) ([]model.InboxMessage, int64, error)
	AdminVerifyDomain(context.Context, string, bool) error
	AdminCreateSession(context.Context, string, string, time.Time) error
	GetAllMessagesAdmin(context.Context, int, int, string) ([]model.Message, int64, error)
	AdminGetUserStats(context.Context, string) (model.UserStats, error)
	SearchLogs(context.Context, string, string, int, int) ([]model.Log, int64, error)
	AdminToggleRecipient(context.Context, string, bool) error
	SearchDomainsAdmin(context.Context, string) ([]model.Domain, error)
	GetMessageCount(context.Context) (int64, error)
	AdminExportRecipients(context.Context) ([]model.Recipient, error)
	AdminExportSubscriptions(context.Context) ([]model.Subscription, error)
	AdminChangeEmail(context.Context, string, string) error
	AdminExportDomains(context.Context) ([]model.Domain, error)
	AdminExportLogs(context.Context) ([]model.Log, error)
	AdminBulkDeleteUsers(context.Context, []string) error
	SearchMessages(context.Context, string, string, int, int) ([]model.Message, int64, error)
	AdminToggleRecipientPGP(context.Context, string, bool) error
	AdminRemoveRecipientPGPKey(context.Context, string) error
	AdminUpdateAlias(context.Context, string, map[string]interface{}) error
	AdminUpdateDomain(context.Context, string, map[string]interface{}) error
	AdminMarkInboxRead(context.Context, uint, bool) error
	AdminGetAllUsersPaginated(context.Context, int, int, string) ([]model.User, int64, error)
	AdminCreateRecipient(context.Context, model.Recipient) error
	AdminCreateDomain(context.Context, model.Domain) error
	AdminExportInbox(context.Context) ([]model.InboxMessage, error)
	AdminExportMessages(context.Context) ([]model.Message, error)
}

func (s *Service) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.Store.GetAllUsers(ctx)
}

func (s *Service) GetSystemStats(ctx context.Context) (any, error) {
	totalUsers, _ := s.Store.GetUserCount(ctx)
	activeUsers, _ := s.Store.GetActiveUserCount(ctx)
	suspendedUsers, _ := s.Store.GetSuspendedUserCount(ctx)
	adminUsers, _ := s.Store.GetAdminCount(ctx)
	totalAliases, _ := s.Store.GetAliasCountAll(ctx)
	totalDomains, _ := s.Store.GetDomainCountAll(ctx)
	totalRecipients, _ := s.Store.GetRecipientCountAll(ctx)
	totalLogs, _ := s.Store.GetLogCount(ctx)
	totalInbox, _ := s.Store.GetInboxMessageCount(ctx)
	totalSubs, _ := s.Store.GetSubscriptionCount(ctx)
	activeSubs, _ := s.Store.GetActiveSubscriptionCount(ctx)
	plans, _ := s.Store.GetActivePlans(ctx)

	return SystemStats{
		TotalUsers:          totalUsers,
		ActiveUsers:         activeUsers,
		SuspendedUsers:      suspendedUsers,
		AdminUsers:          adminUsers,
		TotalAliases:        totalAliases,
		TotalDomains:        totalDomains,
		TotalRecipients:     totalRecipients,
		TotalLogs:           totalLogs,
		TotalInboxMessages:  totalInbox,
		TotalSubscriptions:  totalSubs,
		ActiveSubscriptions: activeSubs,
		ActivePlans:         len(plans),
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

func (s *Service) GetAllAccessKeysAdmin(ctx context.Context, limit, offset int) ([]model.AccessKey, int64, error) {
	if limit <= 0 || limit > 100 { limit = 50 }
	return s.Store.GetAllAccessKeysAdmin(ctx, limit, offset)
}

func (s *Service) AdminDeleteAccessKey(ctx context.Context, keyID string) error {
	return s.Store.AdminDeleteAccessKey(ctx, keyID)
}

func (s *Service) GetAllSessionsAdmin(ctx context.Context, limit, offset int) ([]model.Session, int64, error) {
	if limit <= 0 || limit > 100 { limit = 50 }
	return s.Store.GetAllSessionsAdmin(ctx, limit, offset)
}

func (s *Service) AdminDeleteSession(ctx context.Context, sessionID string) error {
	return s.Store.AdminDeleteSession(ctx, sessionID)
}

func (s *Service) AdminForceLogout(ctx context.Context, userID string) error {
	return s.Store.AdminDeleteSessionsByUserID(ctx, userID)
}

func (s *Service) GetAllCredentialsAdmin(ctx context.Context, limit, offset int) ([]model.Credential, int64, error) {
	if limit <= 0 || limit > 100 { limit = 50 }
	return s.Store.GetAllCredentialsAdmin(ctx, limit, offset)
}

func (s *Service) AdminDeleteCredential(ctx context.Context, credID string) error {
	return s.Store.AdminDeleteCredential(ctx, credID)
}

func (s *Service) AdminUpdateSubscription(ctx context.Context, userID string, tier string, isActive bool, activeUntil string) error {
	return s.Store.AdminUpdateSubscription(ctx, userID, tier, isActive, activeUntil)
}

func (s *Service) AdminBulkUpdateUsers(ctx context.Context, userIDs []string, isActive bool) error {
	return s.Store.AdminBulkUpdateUsers(ctx, userIDs, isActive)
}

func (s *Service) GetAllInboxMessagesAdmin(ctx context.Context, limit, offset int) ([]model.InboxMessage, int64, error) {
	if limit <= 0 || limit > 100 { limit = 50 }
	return s.Store.GetAllInboxMessagesAdmin(ctx, limit, offset)
}

func (s *Service) AdminDeleteInboxMessage(ctx context.Context, msgID uint) error {
	return s.Store.AdminDeleteInboxMessage(ctx, msgID)
}

func (s *Service) AdminPurgeInboxByUser(ctx context.Context, userID string) error {
	return s.Store.AdminPurgeInboxByUser(ctx, userID)
}

func (s *Service) AdminDisableTotp(ctx context.Context, userID string) error {
	return s.Store.AdminDisableTotp(ctx, userID)
}

func (s *Service) AdminResetPassword(ctx context.Context, userID string, passwordPlain string) error {
	hash, err := utils.HashPassword(passwordPlain)
	if err != nil {
		return fmt.Errorf("password hash failed")
	}
	return s.Store.AdminResetPassword(ctx, userID, hash)
}

func (s *Service) AdminGetSettings(ctx context.Context, userID string) (model.Settings, error) {
	return s.Store.AdminGetSettings(ctx, userID)
}

func (s *Service) AdminUpdateSettings(ctx context.Context, userID string, updates map[string]interface{}) error {
	return s.Store.AdminUpdateSettings(ctx, userID, updates)
}

func (s *Service) AdminExportUsers(ctx context.Context) ([]model.User, error) {
	return s.Store.AdminExportUsers(ctx)
}

func (s *Service) AdminExportAliases(ctx context.Context) ([]model.Alias, error) {
	return s.Store.AdminExportAliases(ctx)
}

func (s *Service) GetAllSubscriptionsAdmin(ctx context.Context, limit, offset int, tier string) ([]model.Subscription, int64, error) {
	if limit <= 0 || limit > 100 { limit = 50 }
	return s.Store.GetAllSubscriptionsAdmin(ctx, limit, offset, tier)
}

func (s *Service) AdminDeleteSubscription(ctx context.Context, subID string) error {
	return s.Store.AdminDeleteSubscription(ctx, subID)
}

func (s *Service) AdminImpersonate(ctx context.Context, userID string) (model.User, error) {
	return s.Store.AdminImpersonate(ctx, userID)
}

func (s *Service) AdminBulkDeleteAliases(ctx context.Context, ids []string) error {
	return s.Store.AdminBulkDeleteAliases(ctx, ids)
}

func (s *Service) AdminBulkDeleteDomains(ctx context.Context, ids []string) error {
	return s.Store.AdminBulkDeleteDomains(ctx, ids)
}

func (s *Service) AdminBulkDeleteRecipients(ctx context.Context, ids []string) error {
	return s.Store.AdminBulkDeleteRecipients(ctx, ids)
}

func (s *Service) GetTableSizes(ctx context.Context) (map[string]int64, error) {
	return s.Store.GetTableSizes(ctx)
}

func (s *Service) GetRecentSignups(ctx context.Context, days int) ([]model.User, error) {
	if days <= 0 { days = 7 }
	return s.Store.GetRecentSignups(ctx, days)
}

func (s *Service) AdminVerifyDomain(ctx context.Context, domainID string, verified bool) error {
	return s.Store.AdminVerifyDomain(ctx, domainID, verified)
}

func (s *Service) AdminImpersonateUser(ctx context.Context, userID string) (string, error) {
	user, err := s.Store.AdminImpersonate(ctx, userID)
	if err != nil {
		return "", err
	}
	token, err := model.GenSessionToken()
	if err != nil {
		return "", err
	}
	exp := time.Now().Add(24 * time.Hour)
	if err := s.Store.AdminCreateSession(ctx, token, user.ID, exp); err != nil {
		return "", err
	}
	return token, nil
}

func (s *Service) SearchAccessKeys(ctx context.Context, userID string, limit, offset int) ([]model.AccessKey, int64, error) {
	if limit <= 0 || limit > 100 { limit = 50 }
	return s.Store.SearchAccessKeys(ctx, userID, limit, offset)
}

func (s *Service) SearchSessions(ctx context.Context, userID string, limit, offset int) ([]model.Session, int64, error) {
	if limit <= 0 || limit > 100 { limit = 50 }
	return s.Store.SearchSessions(ctx, userID, limit, offset)
}

func (s *Service) SearchInboxMessages(ctx context.Context, search string, limit, offset int) ([]model.InboxMessage, int64, error) {
	if limit <= 0 || limit > 100 { limit = 50 }
	return s.Store.SearchInboxMessages(ctx, search, limit, offset)
}

func (s *Service) GetAllMessagesAdmin(ctx context.Context, limit, offset int, msgType string) ([]model.Message, int64, error) {
	if limit <= 0 || limit > 100 { limit = 50 }
	return s.Store.GetAllMessagesAdmin(ctx, limit, offset, msgType)
}

func (s *Service) AdminGetUserStats(ctx context.Context, userID string) (model.UserStats, error) {
	return s.Store.AdminGetUserStats(ctx, userID)
}

func (s *Service) SearchLogs(ctx context.Context, search string, logType string, limit, offset int) ([]model.Log, int64, error) {
	if limit <= 0 || limit > 200 { limit = 100 }
	return s.Store.SearchLogs(ctx, search, logType, limit, offset)
}

func (s *Service) AdminToggleRecipient(ctx context.Context, recipientID string, isActive bool) error {
	return s.Store.AdminToggleRecipient(ctx, recipientID, isActive)
}

func (s *Service) SearchDomainsAdmin(ctx context.Context, search string) ([]model.Domain, error) {
	return s.Store.SearchDomainsAdmin(ctx, search)
}

func (s *Service) AdminExportRecipients(ctx context.Context) ([]model.Recipient, error) {
	return s.Store.AdminExportRecipients(ctx)
}

func (s *Service) AdminExportSubscriptions(ctx context.Context) ([]model.Subscription, error) {
	return s.Store.AdminExportSubscriptions(ctx)
}

func (s *Service) AdminChangeEmail(ctx context.Context, userID string, newEmail string) error {
	return s.Store.AdminChangeEmail(ctx, userID, newEmail)
}

func (s *Service) AdminExportDomains(ctx context.Context) ([]model.Domain, error) {
	return s.Store.AdminExportDomains(ctx)
}

func (s *Service) AdminExportLogs(ctx context.Context) ([]model.Log, error) {
	return s.Store.AdminExportLogs(ctx)
}

func (s *Service) AdminBulkDeleteUsers(ctx context.Context, userIDs []string) error {
	return s.Store.AdminBulkDeleteUsers(ctx, userIDs)
}

func (s *Service) SearchMessages(ctx context.Context, search string, msgType string, limit, offset int) ([]model.Message, int64, error) {
	if limit <= 0 || limit > 100 { limit = 50 }
	return s.Store.SearchMessages(ctx, search, msgType, limit, offset)
}

func (s *Service) AdminToggleRecipientPGP(ctx context.Context, recipientID string, pgpEnabled bool) error {
	return s.Store.AdminToggleRecipientPGP(ctx, recipientID, pgpEnabled)
}

func (s *Service) AdminRemoveRecipientPGPKey(ctx context.Context, recipientID string) error {
	return s.Store.AdminRemoveRecipientPGPKey(ctx, recipientID)
}

func (s *Service) AdminUpdateAlias(ctx context.Context, aliasID string, updates map[string]interface{}) error {
	return s.Store.AdminUpdateAlias(ctx, aliasID, updates)
}

func (s *Service) AdminUpdateDomain(ctx context.Context, domainID string, updates map[string]interface{}) error {
	return s.Store.AdminUpdateDomain(ctx, domainID, updates)
}

func (s *Service) AdminMarkInboxRead(ctx context.Context, msgID uint, isRead bool) error {
	return s.Store.AdminMarkInboxRead(ctx, msgID, isRead)
}

func (s *Service) AdminGetAllUsersPaginated(ctx context.Context, limit, offset int, search string) ([]model.User, int64, error) {
	if limit <= 0 || limit > 100 { limit = 50 }
	return s.Store.AdminGetAllUsersPaginated(ctx, limit, offset, search)
}

func (s *Service) AdminCreateRecipient(ctx context.Context, r model.Recipient) error {
	return s.Store.AdminCreateRecipient(ctx, r)
}

func (s *Service) AdminCreateDomain(ctx context.Context, dm model.Domain) error {
	return s.Store.AdminCreateDomain(ctx, dm)
}

func (s *Service) AdminExportInbox(ctx context.Context) ([]model.InboxMessage, error) {
	return s.Store.AdminExportInbox(ctx)
}

func (s *Service) AdminExportMessages(ctx context.Context) ([]model.Message, error) {
	return s.Store.AdminExportMessages(ctx)
}
