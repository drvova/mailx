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
