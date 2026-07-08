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

func (d *Database) AdminGetTopForwarders(ctx context.Context, days int) ([]model.UserForwardStats, error) {
	cutoff := time.Now().AddDate(0, 0, -days)
	type row struct { UserID string; Email string; Count int64 }
	var rows []row
	d.Client.Model(&model.Message{}).
		Select("messages.user_id, users.email, count(*) as count").
		Joins("join users on users.id = messages.user_id").
		Where("messages.created_at >= ? AND messages.type = 0", cutoff).
		Group("messages.user_id, users.email").
		Order("count desc").
		Limit(20).
		Scan(&rows)
	var results []model.UserForwardStats
	for _, r := range rows {
		// Count blocks, replies, sends for same user in same period
		var blocks, replies, sends int64
		d.Client.Model(&model.Message{}).Where("user_id = ? AND type = 1 AND created_at >= ?", r.UserID, cutoff).Count(&blocks)
		d.Client.Model(&model.Message{}).Where("user_id = ? AND type = 2 AND created_at >= ?", r.UserID, cutoff).Count(&replies)
		d.Client.Model(&model.Message{}).Where("user_id = ? AND type = 3 AND created_at >= ?", r.UserID, cutoff).Count(&sends)
		results = append(results, model.UserForwardStats{
			UserID: r.UserID, Email: r.Email,
			Forwards: r.Count, Blocks: blocks, Replies: replies, Sends: sends,
		})
	}
	return results, nil
}

func (d *Database) AdminGetMessageTypeStats(ctx context.Context, days int) (map[string]int64, error) {
	cutoff := time.Now().AddDate(0, 0, -days)
	type row struct { Type int; Count int64 }
	var rows []row
	d.Client.Model(&model.Message{}).Select("type, count(*) as count").Where("created_at >= ?", cutoff).Group("type").Scan(&rows)
	result := map[string]int64{"forward": 0, "block": 0, "reply": 0, "send": 0, "bounce": 0, "inbox": 0}
	names := map[int]string{0: "forward", 1: "block", 2: "reply", 3: "send", 4: "bounce", 5: "inbox"}
	for _, r := range rows { result[names[r.Type]] = r.Count }
	return result, nil
}

func (d *Database) AdminGetRecentAliases(ctx context.Context, limit int) ([]model.Alias, error) {
	var aliases []model.Alias
	d.Client.Order("created_at desc").Limit(limit).Find(&aliases)
	return aliases, nil
}
