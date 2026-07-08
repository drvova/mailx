package repository

import (
	"context"

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
