package repository

import (
	"context"
	"time"

	"ivpn.net/email/api/internal/model"
	"gorm.io/gorm"
)

func (d *Database) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := d.Client.Order("created_at desc").Find(&users).Error
	return users, err
}

func (d *Database) GetUserCount(ctx context.Context) (int64, error) {
	var count int64
	err := d.Client.Model(&model.User{}).Count(&count).Error
	return count, err
}

func (d *Database) GetActiveUserCount(ctx context.Context) (int64, error) {
	var count int64
	err := d.Client.Model(&model.User{}).Where("is_active = ?", true).Count(&count).Error
	return count, err
}

func (d *Database) GetAllSubscriptions(ctx context.Context) ([]model.Subscription, error) {
	var subs []model.Subscription
	err := d.Client.Order("created_at desc").Find(&subs).Error
	return subs, err
}

func (d *Database) GetAllLogs(ctx context.Context) ([]model.Log, error) {
	var logs []model.Log
	err := d.Client.Order("created_at desc").Limit(500).Find(&logs).Error
	return logs, err
}

func (d *Database) GetLogCount(ctx context.Context) (int64, error) {
	var count int64
	err := d.Client.Model(&model.Log{}).Count(&count).Error
	return count, err
}

func (d *Database) GetAliasCountAll(ctx context.Context) (int64, error) {
	var count int64
	err := d.Client.Model(&model.Alias{}).Count(&count).Error
	return count, err
}

func (d *Database) GetDomainCountAll(ctx context.Context) (int64, error) {
	var count int64
	err := d.Client.Model(&model.Domain{}).Count(&count).Error
	return count, err
}

func (d *Database) AdminUpdateUser(ctx context.Context, user model.User) error {
	return d.Client.Model(&model.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"is_active": user.IsActive,
		"is_admin":  user.IsAdmin,
	}).Error
}

func (d *Database) AdminDeleteUser(ctx context.Context, userID string) error {
	return d.Client.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", userID).Delete(&model.Recipient{}).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", userID).Delete(&model.Alias{}).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", userID).Delete(&model.Settings{}).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", userID).Delete(&model.Session{}).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", userID).Delete(&model.Credential{}).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", userID).Delete(&model.AccessKey{}).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", userID).Delete(&model.Log{}).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", userID).Delete(&model.Subscription{}).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", userID).Delete(&model.Domain{}).Error; err != nil {
			return err
		}
		return tx.Where("id = ?", userID).Delete(&model.User{}).Error
	})
}

func (d *Database) AdminAssignPlan(ctx context.Context, userID string, planID string, tier string) error {
	sub := model.Subscription{}
	err := d.Client.Where("user_id = ?", userID).First(&sub).Error
	if err != nil {
		return err
	}
	sub.PlanID = &planID
	sub.Tier = tier
	sub.IsActive = true
	return d.Client.Save(&sub).Error
}

func (d *Database) GetAllAliasesAdmin(ctx context.Context, limit, offset int, search string) ([]model.Alias, int64, error) {
	var aliases []model.Alias
	q := d.Client.Model(&model.Alias{})
	if search != "" {
		q = q.Where("name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	var total int64
	q.Count(&total)
	err := q.Order("created_at desc").Limit(limit).Offset(offset).Find(&aliases).Error
	return aliases, total, err
}

func (d *Database) AdminDeleteAlias(ctx context.Context, aliasID string) error {
	return d.Client.Where("id = ?", aliasID).Delete(&model.Alias{}).Error
}

func (d *Database) AdminToggleAlias(ctx context.Context, aliasID string, enabled bool) error {
	return d.Client.Model(&model.Alias{}).Where("id = ?", aliasID).Update("enabled", enabled).Error
}

func (d *Database) GetAllDomainsAdmin(ctx context.Context) ([]model.Domain, error) {
	var domains []model.Domain
	err := d.Client.Order("created_at desc").Find(&domains).Error
	return domains, err
}

func (d *Database) AdminDeleteDomain(ctx context.Context, domainID string) error {
	return d.Client.Where("id = ?", domainID).Delete(&model.Domain{}).Error
}

func (d *Database) AdminToggleDomain(ctx context.Context, domainID string, enabled bool) error {
	return d.Client.Model(&model.Domain{}).Where("id = ?", domainID).Update("enabled", enabled).Error
}

func (d *Database) GetAllRecipientsAdmin(ctx context.Context, limit, offset int, search string) ([]model.Recipient, int64, error) {
	var recipients []model.Recipient
	q := d.Client.Model(&model.Recipient{})
	if search != "" {
		q = q.Where("email LIKE ?", "%"+search+"%")
	}
	var total int64
	q.Count(&total)
	err := q.Order("created_at desc").Limit(limit).Offset(offset).Find(&recipients).Error
	return recipients, total, err
}

func (d *Database) AdminDeleteRecipient(ctx context.Context, recipientID string) error {
	return d.Client.Where("id = ?", recipientID).Delete(&model.Recipient{}).Error
}

func (d *Database) GetLogsFiltered(ctx context.Context, logType string, limit, offset int) ([]model.Log, int64, error) {
	var logs []model.Log
	q := d.Client.Model(&model.Log{})
	if logType != "" {
		q = q.Where("log_type = ?", logType)
	}
	var total int64
	q.Count(&total)
	err := q.Order("created_at desc").Limit(limit).Offset(offset).Find(&logs).Error
	return logs, total, err
}

func (d *Database) AdminSearchUsers(ctx context.Context, search string, limit, offset int) ([]model.User, int64, error) {
	var users []model.User
	q := d.Client.Model(&model.User{})
	if search != "" {
		q = q.Where("email LIKE ?", "%"+search+"%")
	}
	var total int64
	q.Count(&total)
	err := q.Order("created_at desc").Limit(limit).Offset(offset).Find(&users).Error
	return users, total, err
}

func (d *Database) AdminGetUserDetail(ctx context.Context, userID string) (model.User, model.Subscription, []model.Alias, []model.Recipient, []model.Domain, error) {
	var user model.User
	if err := d.Client.First(&user, "id = ?", userID).Error; err != nil {
		return model.User{}, model.Subscription{}, nil, nil, nil, err
	}
	var sub model.Subscription
	d.Client.Where("user_id = ?", userID).First(&sub)
	var aliases []model.Alias
	d.Client.Where("user_id = ?", userID).Order("created_at desc").Limit(50).Find(&aliases)
	var recipients []model.Recipient
	d.Client.Where("user_id = ?", userID).Find(&recipients)
	var domains []model.Domain
	d.Client.Where("user_id = ?", userID).Find(&domains)
	return user, sub, aliases, recipients, domains, nil
}

func (d *Database) GetAllAccessKeysAdmin(ctx context.Context, limit, offset int) ([]model.AccessKey, int64, error) {
	var keys []model.AccessKey
	var total int64
	d.Client.Model(&model.AccessKey{}).Count(&total)
	err := d.Client.Order("created_at desc").Limit(limit).Offset(offset).Find(&keys).Error
	return keys, total, err
}

func (d *Database) AdminDeleteAccessKey(ctx context.Context, keyID string) error {
	return d.Client.Where("id = ?", keyID).Delete(&model.AccessKey{}).Error
}

func (d *Database) GetAllSessionsAdmin(ctx context.Context, limit, offset int) ([]model.Session, int64, error) {
	var sessions []model.Session
	var total int64
	d.Client.Model(&model.Session{}).Count(&total)
	err := d.Client.Order("created_at desc").Limit(limit).Offset(offset).Find(&sessions).Error
	return sessions, total, err
}

func (d *Database) AdminDeleteSession(ctx context.Context, sessionID string) error {
	return d.Client.Where("id = ?", sessionID).Delete(&model.Session{}).Error
}

func (d *Database) AdminDeleteSessionsByUserID(ctx context.Context, userID string) error {
	return d.Client.Where("user_id = ?", userID).Delete(&model.Session{}).Error
}

func (d *Database) GetAllCredentialsAdmin(ctx context.Context, limit, offset int) ([]model.Credential, int64, error) {
	var creds []model.Credential
	var total int64
	d.Client.Model(&model.Credential{}).Count(&total)
	err := d.Client.Order("created_at desc").Limit(limit).Offset(offset).Find(&creds).Error
	return creds, total, err
}

func (d *Database) AdminDeleteCredential(ctx context.Context, credID string) error {
	return d.Client.Where("id = ?", credID).Delete(&model.Credential{}).Error
}

func (d *Database) AdminUpdateSubscription(ctx context.Context, userID string, tier string, isActive bool, activeUntil string) error {
	updates := map[string]interface{}{}
	if tier != "" {
		updates["tier"] = tier
	}
	updates["is_active"] = isActive
	if activeUntil != "" {
		updates["active_until"] = activeUntil
	}
	return d.Client.Model(&model.Subscription{}).Where("user_id = ?", userID).Updates(updates).Error
}

func (d *Database) GetSuspendedUserCount(ctx context.Context) (int64, error) {
	var count int64
	err := d.Client.Model(&model.User{}).Where("is_active = ?", false).Count(&count).Error
	return count, err
}

func (d *Database) GetAdminCount(ctx context.Context) (int64, error) {
	var count int64
	err := d.Client.Model(&model.User{}).Where("is_admin = ?", true).Count(&count).Error
	return count, err
}

func (d *Database) AdminBulkUpdateUsers(ctx context.Context, userIDs []string, isActive bool) error {
	return d.Client.Model(&model.User{}).Where("id IN ?", userIDs).Update("is_active", isActive).Error
}

func (d *Database) GetAllInboxMessagesAdmin(ctx context.Context, limit, offset int) ([]model.InboxMessage, int64, error) {
	var msgs []model.InboxMessage
	var total int64
	d.Client.Model(&model.InboxMessage{}).Count(&total)
	err := d.Client.Select("id", "created_at", "user_id", "alias_id", "from", "from_name", "subject", "read", "size").
		Order("created_at desc").Limit(limit).Offset(offset).Find(&msgs).Error
	return msgs, total, err
}

func (d *Database) AdminDeleteInboxMessage(ctx context.Context, msgID uint) error {
	return d.Client.Where("id = ?", msgID).Delete(&model.InboxMessage{}).Error
}

func (d *Database) AdminPurgeInboxByUser(ctx context.Context, userID string) error {
	return d.Client.Where("user_id = ?", userID).Delete(&model.InboxMessage{}).Error
}

func (d *Database) AdminDisableTotp(ctx context.Context, userID string) error {
	return d.Client.Model(&model.User{}).Where("id = ?", userID).Updates(map[string]interface{}{"totp_secret": "", "totp_backup": "", "totp_backup_used": ""}).Error
}

func (d *Database) AdminResetPassword(ctx context.Context, userID string, passwordHash string) error {
	return d.Client.Model(&model.User{}).Where("id = ?", userID).Update("password_hash", passwordHash).Error
}

func (d *Database) AdminGetSettings(ctx context.Context, userID string) (model.Settings, error) {
	var s model.Settings
	err := d.Client.Where("user_id = ?", userID).First(&s).Error
	return s, err
}

func (d *Database) AdminUpdateSettings(ctx context.Context, userID string, updates map[string]interface{}) error {
	return d.Client.Model(&model.Settings{}).Where("user_id = ?", userID).Updates(updates).Error
}

func (d *Database) GetInboxMessageCount(ctx context.Context) (int64, error) {
	var count int64
	err := d.Client.Model(&model.InboxMessage{}).Count(&count).Error
	return count, err
}

func (d *Database) GetRecipientCountAll(ctx context.Context) (int64, error) {
	var count int64
	err := d.Client.Model(&model.Recipient{}).Count(&count).Error
	return count, err
}

func (d *Database) GetSubscriptionCount(ctx context.Context) (int64, error) {
	var count int64
	err := d.Client.Model(&model.Subscription{}).Count(&count).Error
	return count, err
}

func (d *Database) GetActiveSubscriptionCount(ctx context.Context) (int64, error) {
	var count int64
	err := d.Client.Model(&model.Subscription{}).Where("is_active = ?", true).Count(&count).Error
	return count, err
}

func (d *Database) AdminExportUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := d.Client.Order("created_at desc").Find(&users).Error
	return users, err
}

func (d *Database) AdminExportAliases(ctx context.Context) ([]model.Alias, error) {
	var aliases []model.Alias
	err := d.Client.Order("created_at desc").Find(&aliases).Error
	return aliases, err
}

func (d *Database) GetAllSubscriptionsAdmin(ctx context.Context, limit, offset int, tier string) ([]model.Subscription, int64, error) {
	var subs []model.Subscription
	q := d.Client.Model(&model.Subscription{})
	if tier != "" {
		q = q.Where("tier = ?", tier)
	}
	var total int64
	q.Count(&total)
	err := q.Order("created_at desc").Limit(limit).Offset(offset).Find(&subs).Error
	return subs, total, err
}

func (d *Database) AdminDeleteSubscription(ctx context.Context, subID string) error {
	return d.Client.Where("id = ?", subID).Delete(&model.Subscription{}).Error
}

func (d *Database) AdminImpersonate(ctx context.Context, userID string) (model.User, error) {
	var user model.User
	err := d.Client.First(&user, "id = ?", userID).Error
	return user, err
}

func (d *Database) AdminBulkDeleteAliases(ctx context.Context, aliasIDs []string) error {
	return d.Client.Where("id IN ?", aliasIDs).Delete(&model.Alias{}).Error
}

func (d *Database) AdminBulkDeleteDomains(ctx context.Context, domainIDs []string) error {
	return d.Client.Where("id IN ?", domainIDs).Delete(&model.Domain{}).Error
}

func (d *Database) AdminBulkDeleteRecipients(ctx context.Context, recipientIDs []string) error {
	return d.Client.Where("id IN ?", recipientIDs).Delete(&model.Recipient{}).Error
}

func (d *Database) GetTableSizes(ctx context.Context) (map[string]int64, error) {
	sizes := map[string]int64{}
	tables := []string{"users", "subscriptions", "aliases", "domains", "recipients", "messages", "inbox_messages", "settings", "sessions", "credentials", "access_keys", "logs", "plans"}
	for _, t := range tables {
		var count int64
		d.Client.Table(t).Count(&count)
		sizes[t] = count
	}
	return sizes, nil
}

func (d *Database) GetRecentSignups(ctx context.Context, days int) ([]model.User, error) {
	var users []model.User
	cutoff := time.Now().AddDate(0, 0, -days)
	err := d.Client.Where("created_at >= ?", cutoff).Order("created_at desc").Find(&users).Error
	return users, err
}

func (d *Database) SearchAccessKeys(ctx context.Context, userID string, limit, offset int) ([]model.AccessKey, int64, error) {
	var keys []model.AccessKey
	q := d.Client.Model(&model.AccessKey{})
	if userID != "" {
		q = q.Where("user_id = ?", userID)
	}
	var total int64
	q.Count(&total)
	err := q.Order("created_at desc").Limit(limit).Offset(offset).Find(&keys).Error
	return keys, total, err
}

func (d *Database) SearchSessions(ctx context.Context, userID string, limit, offset int) ([]model.Session, int64, error) {
	var sessions []model.Session
	q := d.Client.Model(&model.Session{})
	if userID != "" {
		q = q.Where("user_id = ?", userID)
	}
	var total int64
	q.Count(&total)
	err := q.Order("created_at desc").Limit(limit).Offset(offset).Find(&sessions).Error
	return sessions, total, err
}

func (d *Database) SearchInboxMessages(ctx context.Context, search string, limit, offset int) ([]model.InboxMessage, int64, error) {
	var msgs []model.InboxMessage
	q := d.Client.Model(&model.InboxMessage{})
	if search != "" {
		q = q.Where("from LIKE ? OR subject LIKE ? OR from_name LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	var total int64
	q.Count(&total)
	err := q.Select("id", "created_at", "user_id", "alias_id", "from", "from_name", "subject", "read", "size").
		Order("created_at desc").Limit(limit).Offset(offset).Find(&msgs).Error
	return msgs, total, err
}

func (d *Database) AdminVerifyDomain(ctx context.Context, domainID string, verified bool) error {
	now := time.Now()
	updates := map[string]interface{}{}
	if verified {
		updates["mx_verified_at"] = now
		updates["owner_verified_at"] = now
		updates["send_verified_at"] = now
	} else {
		updates["mx_verified_at"] = nil
		updates["owner_verified_at"] = nil
		updates["send_verified_at"] = nil
	}
	return d.Client.Model(&model.Domain{}).Where("id = ?", domainID).Updates(updates).Error
}

func (d *Database) AdminBulkActivateUsers(ctx context.Context, userIDs []string, isActive bool) error {
	return d.Client.Model(&model.User{}).Where("id IN ?", userIDs).Update("is_active", isActive).Error
}

func (d *Database) AdminGetConfig(ctx context.Context) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

func (d *Database) AdminCreateSession(ctx context.Context, token string, userID string, exp time.Time) error {
	sessionData := model.Session{}
	sessionData.UserID = userID
	sessionData.Token = token
	sessionData.ExpiresAt = exp
	return d.Client.Create(&sessionData).Error
}

func (d *Database) GetAllMessagesAdmin(ctx context.Context, limit, offset int, msgType string) ([]model.Message, int64, error) {
	var msgs []model.Message
	q := d.Client.Model(&model.Message{})
	if msgType != "" {
		q = q.Where("type = ?", msgType)
	}
	var total int64
	q.Count(&total)
	err := q.Order("created_at desc").Limit(limit).Offset(offset).Find(&msgs).Error
	return msgs, total, err
}

func (d *Database) AdminGetUserStats(ctx context.Context, userID string) (model.UserStats, error) {
	var stats model.UserStats
	var fw, bl, rp, sd int64
	d.Client.Model(&model.Message{}).Where("user_id = ? AND type = ?", userID, model.Forward).Count(&fw)
	d.Client.Model(&model.Message{}).Where("user_id = ? AND type = ?", userID, model.Block).Count(&bl)
	d.Client.Model(&model.Message{}).Where("user_id = ? AND type = ?", userID, model.Reply).Count(&rp)
	d.Client.Model(&model.Message{}).Where("user_id = ? AND type = ?", userID, model.Send).Count(&sd)
	stats.Forwards = int(fw)
	stats.Blocks = int(bl)
	stats.Replies = int(rp)
	stats.Sends = int(sd)
	d.Client.Model(&model.Alias{}).Where("user_id = ?", userID).Count(&stats.Aliases)
	return stats, nil
}

func (d *Database) SearchLogs(ctx context.Context, search string, logType string, limit, offset int) ([]model.Log, int64, error) {
	q := d.Client.Model(&model.Log{})
	if logType != "" {
		q = q.Where("log_type = ?", logType)
	}
	if search != "" {
		q = q.Where("from LIKE ? OR destination LIKE ? OR message LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	var total int64
	q.Count(&total)
	var logs []model.Log
	err := q.Order("created_at desc").Limit(limit).Offset(offset).Find(&logs).Error
	return logs, total, err
}

func (d *Database) AdminToggleRecipient(ctx context.Context, recipientID string, isActive bool) error {
	return d.Client.Model(&model.Recipient{}).Where("id = ?", recipientID).Update("is_active", isActive).Error
}

func (d *Database) SearchDomainsAdmin(ctx context.Context, search string) ([]model.Domain, error) {
	var domains []model.Domain
	q := d.Client.Model(&model.Domain{})
	if search != "" {
		q = q.Where("name LIKE ?", "%"+search+"%")
	}
	err := q.Order("created_at desc").Find(&domains).Error
	return domains, err
}

func (d *Database) GetMessageCount(ctx context.Context) (int64, error) {
	var count int64
	err := d.Client.Model(&model.Message{}).Count(&count).Error
	return count, err
}

func (d *Database) AdminExportRecipients(ctx context.Context) ([]model.Recipient, error) {
	var recipients []model.Recipient
	err := d.Client.Order("created_at desc").Find(&recipients).Error
	return recipients, err
}

func (d *Database) AdminExportSubscriptions(ctx context.Context) ([]model.Subscription, error) {
	var subs []model.Subscription
	err := d.Client.Order("created_at desc").Find(&subs).Error
	return subs, err
}
