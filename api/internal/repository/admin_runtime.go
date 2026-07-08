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
