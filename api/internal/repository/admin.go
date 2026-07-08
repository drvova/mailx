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

func (d *Database) AdminChangeEmail(ctx context.Context, userID string, newEmail string) error {
	return d.Client.Model(&model.User{}).Where("id = ?", userID).Update("email", newEmail).Error
}

func (d *Database) AdminExportDomains(ctx context.Context) ([]model.Domain, error) {
	var domains []model.Domain
	err := d.Client.Order("created_at desc").Find(&domains).Error
	return domains, err
}

func (d *Database) AdminExportLogs(ctx context.Context) ([]model.Log, error) {
	var logs []model.Log
	err := d.Client.Order("created_at desc").Limit(10000).Find(&logs).Error
	return logs, err
}

func (d *Database) AdminBulkDeleteUsers(ctx context.Context, userIDs []string) error {
	return d.Client.Transaction(func(tx *gorm.DB) error {
		for _, model := range []interface{}{&model.Recipient{}, &model.Alias{}, &model.Settings{}, &model.Session{}, &model.Credential{}, &model.AccessKey{}, &model.Log{}, &model.Subscription{}, &model.Domain{}, &model.InboxMessage{}} {
			if err := tx.Where("user_id = ?", userIDs).Delete(model).Error; err != nil {
				return err
			}
		}
		return tx.Where("id IN ?", userIDs).Delete(&model.User{}).Error
	})
}

func (d *Database) SearchMessages(ctx context.Context, search string, msgType string, limit, offset int) ([]model.Message, int64, error) {
	var msgs []model.Message
	q := d.Client.Model(&model.Message{})
	if msgType != "" {
		q = q.Where("type = ?", msgType)
	}
	if search != "" {
		q = q.Where("user_id LIKE ? OR alias_id LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	var total int64
	q.Count(&total)
	err := q.Order("created_at desc").Limit(limit).Offset(offset).Find(&msgs).Error
	return msgs, total, err
}

func (d *Database) AdminToggleRecipientPGP(ctx context.Context, recipientID string, pgpEnabled bool) error {
	return d.Client.Model(&model.Recipient{}).Where("id = ?", recipientID).Update("pgp_enabled", pgpEnabled).Error
}

func (d *Database) AdminRemoveRecipientPGPKey(ctx context.Context, recipientID string) error {
	return d.Client.Model(&model.Recipient{}).Where("id = ?", recipientID).Updates(map[string]interface{}{"pgp_key": "", "pgp_enabled": false, "pgp_inline": false}).Error
}

func (d *Database) AdminUpdateAlias(ctx context.Context, aliasID string, updates map[string]interface{}) error {
	return d.Client.Model(&model.Alias{}).Where("id = ?", aliasID).Updates(updates).Error
}

func (d *Database) AdminUpdateDomain(ctx context.Context, domainID string, updates map[string]interface{}) error {
	return d.Client.Model(&model.Domain{}).Where("id = ?", domainID).Updates(updates).Error
}

func (d *Database) AdminMarkInboxRead(ctx context.Context, msgID uint, isRead bool) error {
	return d.Client.Model(&model.InboxMessage{}).Where("id = ?", msgID).Update("read", isRead).Error
}

func (d *Database) AdminGetAllUsersPaginated(ctx context.Context, limit, offset int, search string) ([]model.User, int64, error) {
	q := d.Client.Model(&model.User{})
	if search != "" {
		q = q.Where("email LIKE ?", "%"+search+"%")
	}
	var total int64
	q.Count(&total)
	var users []model.User
	err := q.Order("created_at desc").Limit(limit).Offset(offset).Find(&users).Error
	return users, total, err
}

func (d *Database) AdminCreateRecipient(ctx context.Context, r model.Recipient) error {
	return d.Client.Create(&r).Error
}

func (d *Database) AdminCreateDomain(ctx context.Context, dm model.Domain) error {
	return d.Client.Create(&dm).Error
}

func (d *Database) AdminExportInbox(ctx context.Context) ([]model.InboxMessage, error) {
	var msgs []model.InboxMessage
	err := d.Client.Select("id, user_id, alias_id, \"from\", from_name, subject, read, size, created_at").Order("created_at desc").Limit(10000).Find(&msgs).Error
	return msgs, err
}

func (d *Database) AdminExportMessages(ctx context.Context) ([]model.Message, error) {
	var msgs []model.Message
	err := d.Client.Order("created_at desc").Limit(10000).Find(&msgs).Error
	return msgs, err
}

func (d *Database) AdminCreateAlias(ctx context.Context, a model.Alias) error {
	return d.Client.Create(&a).Error
}

func (d *Database) AdminUpdateRecipient(ctx context.Context, recipientID string, updates map[string]interface{}) error {
	return d.Client.Model(&model.Recipient{}).Where("id = ?", recipientID).Updates(updates).Error
}

func (d *Database) AdminDeleteLog(ctx context.Context, logID string) error {
	return d.Client.Where("id = ?", logID).Delete(&model.Log{}).Error
}

func (d *Database) AdminBulkDeleteInbox(ctx context.Context, msgIDs []uint) error {
	return d.Client.Where("id IN ?", msgIDs).Delete(&model.InboxMessage{}).Error
}

func (d *Database) AdminExtendSubscription(ctx context.Context, subID string, days int) error {
	var sub model.Subscription
	if err := d.Client.First(&sub, "id = ?", subID).Error; err != nil {
		return err
	}
	base := sub.ActiveUntil
	if !base.After(time.Now()) {
		base = time.Now()
	}
	sub.ActiveUntil = base.AddDate(0, 0, days)
	return d.Client.Save(&sub).Error
}

func (d *Database) AdminCreateAccessKey(ctx context.Context, k model.AccessKey) (model.AccessKey, error) {
	err := d.Client.Create(&k).Error
	return k, err
}

func (d *Database) AdminTransferAlias(ctx context.Context, aliasID string, newUserID string) error {
	return d.Client.Model(&model.Alias{}).Where("id = ?", aliasID).Update("user_id", newUserID).Error
}

func (d *Database) AdminTransferDomain(ctx context.Context, domainID string, newUserID string) error {
	return d.Client.Model(&model.Domain{}).Where("id = ?", domainID).Update("user_id", newUserID).Error
}

func (d *Database) AdminPurgeLogs(ctx context.Context, days int, logType string) (int64, error) {
	q := d.Client.Where("created_at < ?", time.Now().AddDate(0, 0, -days))
	if logType != "" {
		q = q.Where("type = ?", logType)
	}
	res := q.Delete(&model.Log{})
	return res.RowsAffected, res.Error
}

func (d *Database) AdminPurgeAllInbox(ctx context.Context) (int64, error) {
	res := d.Client.Where("1 = 1").Delete(&model.InboxMessage{})
	return res.RowsAffected, res.Error
}

func (d *Database) AdminGetInboxRaw(ctx context.Context, msgID uint) ([]byte, error) {
	var msg model.InboxMessage
	err := d.Client.Select("raw").First(&msg, "id = ?", msgID).Error
	return msg.Raw, err
}

func (d *Database) AdminSetAliasExpiry(ctx context.Context, aliasID string, expiresAt *time.Time) error {
	return d.Client.Model(&model.Alias{}).Where("id = ?", aliasID).Update("expires_at", expiresAt).Error
}

func (d *Database) AdminSetAccessKeyExpiry(ctx context.Context, keyID string, expiresAt *time.Time) error {
	return d.Client.Model(&model.AccessKey{}).Where("id = ?", keyID).Update("expires_at", expiresAt).Error
}

func (d *Database) AdminLogAudit(ctx context.Context, entry model.AdminAudit) error {
	return d.Client.Create(&entry).Error
}

func (d *Database) AdminGetAuditLog(ctx context.Context, limit, offset int) ([]model.AdminAudit, int64, error) {
	var entries []model.AdminAudit
	var total int64
	d.Client.Model(&model.AdminAudit{}).Count(&total)
	err := d.Client.Order("created_at desc").Limit(limit).Offset(offset).Find(&entries).Error
	return entries, total, err
}

func (d *Database) AdminGetSessionData(ctx context.Context, sessionID string) ([]byte, error) {
	var s model.Session
	err := d.Client.Select("data").First(&s, "id = ?", sessionID).Error
	return s.Data, err
}

func (d *Database) AdminBulkDeleteAccessKeys(ctx context.Context, keyIDs []string) error {
	return d.Client.Where("id IN ?", keyIDs).Delete(&model.AccessKey{}).Error
}

func (d *Database) AdminBulkDeleteCredentials(ctx context.Context, credIDs []string) error {
	return d.Client.Where("id IN ?", credIDs).Delete(&model.Credential{}).Error
}

func (d *Database) AdminBulkExtendSubscriptions(ctx context.Context, subIDs []string, days int) (int64, error) {
	res := d.Client.Model(&model.Subscription{}).Where("id IN ?", subIDs).
		Update("active_until", gorm.Expr("CASE WHEN active_until > NOW() THEN DATE_ADD(active_until, INTERVAL ? DAY) ELSE DATE_ADD(NOW(), INTERVAL ? DAY) END", days, days))
	return res.RowsAffected, res.Error
}

func (d *Database) AdminExportUsersEnriched(ctx context.Context) ([]model.UserWithSub, error) {
	var results []model.UserWithSub
	err := d.Client.Model(&model.User{}).
		Select("users.id, users.email, users.is_active, users.is_admin, users.created_at, subscriptions.tier, subscriptions.type as sub_type, subscriptions.is_active as sub_active, subscriptions.active_until").
		Joins("left join subscriptions on subscriptions.user_id = users.id").
		Order("users.created_at desc").
		Scan(&results).Error
	return results, err
}

func (d *Database) AdminBulkDeleteMessages(ctx context.Context, msgIDs []uint) error {
	return d.Client.Where("id IN ?", msgIDs).Delete(&model.Message{}).Error
}

func (d *Database) AdminSetRecipientPGP(ctx context.Context, recipientID string, pgpKey string, pgpInline bool) error {
	return d.Client.Model(&model.Recipient{}).Where("id = ?", recipientID).Updates(map[string]interface{}{"pgp_key": pgpKey, "pgp_enabled": pgpKey != "", "pgp_inline": pgpInline}).Error
}

func (d *Database) AdminGetDomainDNS(ctx context.Context, domainID string) (model.DNSConfig, error) {
	var dm model.Domain
	err := d.Client.First(&dm, "id = ?", domainID).Error
	if err != nil {
		return model.DNSConfig{}, err
	}
	return model.DNSConfig{Domain: dm.Name}, nil
}

func (d *Database) AdminUpdateUserNotes(ctx context.Context, userID string, notes string) error {
	return d.Client.Model(&model.User{}).Where("id = ?", userID).Update("notes", notes).Error
}

func (d *Database) AdminGetSubscriptionStats(ctx context.Context) (active, expired, grace int64, err error) {
	now := time.Now()
	graceCut := now.AddDate(0, 0, -30)
	d.Client.Model(&model.Subscription{}).Where("is_active = true AND active_until > ?", now).Count(&active)
	d.Client.Model(&model.Subscription{}).Where("active_until <= ? AND active_until > '0001-01-01'", now).Count(&expired)
	d.Client.Model(&model.Subscription{}).Where("is_active = true AND active_until > ? AND active_until <= ?", graceCut, now).Count(&grace)
	return
}

func (d *Database) AdminGetDailyActivity(ctx context.Context, days int) ([]model.DailyStats, error) {
	var results []model.DailyStats
	start := time.Now().AddDate(0, 0, -days)
	// Get daily forwards, blocks, replies, sends from messages
	rows, err := d.Client.Model(&model.Message{}).
		Select("date(created_at) as date, count(*) as total, sum(case when type = 0 then 1 else 0 end) as forwards, sum(case when type = 1 then 1 else 0 end) as blocks, sum(case when type = 2 then 1 else 0 end) as replies, sum(case when type = 3 then 1 else 0 end) as sends").
		Where("created_at >= ?", start).
		Group("date(created_at)").
		Order("date desc").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ds model.DailyStats
		rows.Scan(&ds.Date, &ds.Total, &ds.Forwards, &ds.Blocks, &ds.Replies, &ds.Sends)
		results = append(results, ds)
	}
	// Also get daily signups
	signupRows, err2 := d.Client.Model(&model.User{}).
		Select("date(created_at) as date, count(*) as signups").
		Where("created_at >= ?", start).
		Group("date(created_at)").
		Order("date desc").
		Rows()
	if err2 == nil {
		defer signupRows.Close()
		signupMap := map[string]int64{}
		for signupRows.Next() {
			var date string; var count int64
			signupRows.Scan(&date, &count)
			signupMap[date] = count
		}
		for i := range results { results[i].Signups = signupMap[results[i].Date] }
	}
	return results, nil
}

func (d *Database) AdminGetPlanDistribution(ctx context.Context) (map[string]int64, error) {
	var rows []struct { Tier string; Count int64 }
	err := d.Client.Model(&model.Subscription{}).
		Select("tier, count(*) as count").
		Where("tier != '' AND is_active = true").
		Group("tier").
		Scan(&rows).Error
	if err != nil { return nil, err }
	result := map[string]int64{}
	for _, r := range rows { result[r.Tier] = r.Count }
	return result, nil
}

func (d *Database) AdminGetDomainHealth(ctx context.Context) (verified, unverified int64, err error) {
	d.Client.Model(&model.Domain{}).Where("mx_verified_at is not null AND mx_verified_at > '0001-01-01'").Count(&verified)
	d.Client.Model(&model.Domain{}).Where("mx_verified_at is null OR mx_verified_at <= '0001-01-01'").Count(&unverified)
	return
}

func (d *Database) AdminGlobalUserSearch(ctx context.Context, query string) (*model.User, *model.Subscription, []model.Alias, []model.Domain, []model.Recipient, error) {
	var user model.User
	q := d.Client.Where("email LIKE ?", "%"+query+"%").First(&user)
	if q.Error != nil {
		return nil, nil, nil, nil, nil, q.Error
	}
	var sub model.Subscription
	d.Client.Where("user_id = ?", user.ID).First(&sub)
	var aliases []model.Alias
	d.Client.Where("user_id = ?", user.ID).Limit(50).Find(&aliases)
	var domains []model.Domain
	d.Client.Where("user_id = ?", user.ID).Find(&domains)
	var recipients []model.Recipient
	d.Client.Where("user_id = ?", user.ID).Limit(50).Find(&recipients)
	return &user, &sub, aliases, domains, recipients, nil
}

func (d *Database) AdminGetUserLastActive(ctx context.Context, userID string) (*time.Time, error) {
	var t time.Time
	// Check most recent session
	err := d.Client.Model(&model.Session{}).Select("created_at").Where("user_id = ?", userID).Order("created_at desc").Limit(1).Pluck("created_at", &t).Error
	if err == nil && !t.IsZero() {
		return &t, nil
	}
	// Fallback: most recent message
	err = d.Client.Model(&model.Message{}).Select("created_at").Where("user_id = ?", userID).Order("created_at desc").Limit(1).Pluck("created_at", &t).Error
	if err == nil && !t.IsZero() {
		return &t, nil
	}
	return nil, nil
}

func (d *Database) AdminGetInactiveUsers(ctx context.Context, days int) ([]model.User, int64, error) {
	cutoff := time.Now().AddDate(0, 0, -days)
	// Users whose most recent message or session is older than cutoff
	var ids []string
	d.Client.Model(&model.User{}).Select("users.id").
		Joins("left join (select user_id, max(created_at) as last_act from (select user_id, created_at from messages union all select user_id, created_at from sessions) t group by user_id) last_activity on last_activity.user_id = users.id").
		Where("last_activity.last_act < ? OR last_activity.last_act IS NULL", cutoff).
		Limit(200).
		Pluck("users.id", &ids)
	var users []model.User
	var total int64
	if len(ids) > 0 {
		d.Client.Where("id IN ?", ids).Order("created_at desc").Find(&users)
		total = int64(len(ids))
	}
	return users, total, nil
}

func (d *Database) AdminToggleAliasCatchAll(ctx context.Context, aliasID string, catchAll bool) error {
	return d.Client.Model(&model.Alias{}).Where("id = ?", aliasID).Update("catch_all", catchAll).Error
}

func (d *Database) AdminExportUserData(ctx context.Context, userID string) (*model.User, *model.Subscription, []model.Alias, []model.Domain, []model.Recipient, []model.AccessKey, *model.Settings, error) {
	var user model.User
	if err := d.Client.First(&user, "id = ?", userID).Error; err != nil { return nil, nil, nil, nil, nil, nil, nil, err }
	var sub model.Subscription
	d.Client.Where("user_id = ?", userID).First(&sub)
	var aliases []model.Alias
	d.Client.Where("user_id = ?", userID).Find(&aliases)
	var domains []model.Domain
	d.Client.Where("user_id = ?", userID).Find(&domains)
	var recipients []model.Recipient
	d.Client.Where("user_id = ?", userID).Find(&recipients)
	var keys []model.AccessKey
	d.Client.Where("user_id = ?", userID).Find(&keys)
	var settings model.Settings
	d.Client.Where("user_id = ?", userID).First(&settings)
	return &user, &sub, aliases, domains, recipients, keys, &settings, nil
}

func (d *Database) AdminPurgeExpiredSessions(ctx context.Context) (int64, error) {
	res := d.Client.Where("expires_at < ?", time.Now()).Delete(&model.Session{})
	return res.RowsAffected, res.Error
}

func (d *Database) AdminGetDomainWithAliasCounts(ctx context.Context) ([]model.DomainStats, error) {
	var domains []model.Domain
	d.Client.Order("created_at desc").Find(&domains)
	var results []model.DomainStats
	for _, dm := range domains {
		var count int64
		d.Client.Model(&model.Alias{}).Where("name LIKE ?", "%@"+dm.Name).Count(&count)
		results = append(results, model.DomainStats{
			Domain:    dm.Name,
			Enabled:   dm.Enabled,
			Verified:  dm.MXVerifiedAt != nil && dm.MXVerifiedAt.After(time.Time{}),
			AliasCount: count,
		})
	}
	return results, nil
}

func (d *Database) AdminGetLogsDateRange(ctx context.Context, logType, from, to string, limit, offset int) ([]model.Log, int64, error) {
	q := d.Client.Model(&model.Log{})
	if logType != "" {
		q = q.Where("type = ?", logType)
	}
	if from != "" {
		ft, err := time.Parse("2006-01-02", from)
		if err == nil {
			q = q.Where("created_at >= ?", ft)
		}
	}
	if to != "" {
		tt, err := time.Parse("2006-01-02", to)
		if err == nil {
			q = q.Where("created_at < ?", tt.AddDate(0, 0, 1))
		}
	}
	var total int64
	q.Count(&total)
	var logs []model.Log
	if limit <= 0 || limit > 200 { limit = 100 }
	err := q.Order("created_at desc").Limit(limit).Offset(offset).Find(&logs).Error
	return logs, total, err
}
