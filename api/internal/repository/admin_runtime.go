package repository

import (
	"context"
	"runtime"
	"time"

	"ivpn.net/email/api/internal/model"
)

var startTime = time.Now()

func (d *Database) AdminGetRuntimeStats(ctx context.Context) (map[string]interface{}, error) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	stats := map[string]interface{}{
		"goroutines":  runtime.NumGoroutine(),
		"heap_alloc_mb": float64(m.HeapAlloc) / 1024 / 1024,
		"heap_sys_mb":   float64(m.HeapSys) / 1024 / 1024,
		"total_alloc_mb": float64(m.TotalAlloc) / 1024 / 1024,
		"num_gc":        m.NumGC,
		"uptime_seconds":  int64(time.Since(startTime).Seconds()),
	}
	return stats, nil
}

func (d *Database) AdminGetUserQuota(ctx context.Context, userID string) (*model.UserQuota, error) {
	var uq model.UserQuota
	uq.UserID = userID

	d.Client.Model(&model.Alias{}).Where("user_id = ?", userID).Count(&uq.AliasCount)
	d.Client.Model(&model.Recipient{}).Where("user_id = ?", userID).Count(&uq.RecipientCount)
	d.Client.Model(&model.Credential{}).Where("user_id = ?", userID).Count(&uq.CredentialCount)
	d.Client.Model(&model.Session{}).Where("user_id = ?", userID).Count(&uq.SessionCount)

	// Get subscription + plan limits
	var sub model.Subscription
	if err := d.Client.Where("user_id = ?", userID).First(&sub).Error; err == nil {
		uq.Tier = sub.Tier
		if sub.PlanID != nil {
			var plan model.Plan
			if err := d.Client.First(&plan, "id = ?", *sub.PlanID).Error; err == nil {
				uq.MaxAliases = int64(plan.MaxDailyAliases)
				uq.MaxRecipients = int64(plan.MaxRecipients)
				uq.MaxCredentials = int64(plan.MaxCredentials)
				uq.MaxSessions = int64(plan.MaxSessions)
			}
		}
	}
	return &uq, nil
}

func (d *Database) AdminCompareUsers(ctx context.Context, id1, id2 string) ([]model.User, []model.Subscription, error) {
	var users []model.User
	d.Client.Where("id IN ?", []string{id1, id2}).Find(&users)
	var subs []model.Subscription
	d.Client.Where("user_id IN ?", []string{id1, id2}).Find(&subs)
	return users, subs, nil
}

func (d *Database) AdminGetRecipientDomains(ctx context.Context) (map[string]int64, error) {
	type row struct { Domain string; Count int64 }
	var rows []row
	d.Client.Model(&model.Recipient{}).Select("substring_index(email, '@', -1) as domain, count(*) as count").Group("substring_index(email, '@', -1)").Order("count desc").Limit(30).Scan(&rows)
	result := map[string]int64{}
	for _, r := range rows { result[r.Domain] = r.Count }
	return result, nil
}
