package repository

import (
	"context"
	"runtime"
	"time"

	"ivpn.net/email/api/internal/model"
)

var startTime = time.Now()

func (d *Database) AdminGetPlanUsage(ctx context.Context) ([]model.PlanUsage, error) {
	var results []model.PlanUsage
	var subs []model.Subscription
	d.Client.Where("is_active = true AND plan_id IS NOT NULL").Find(&subs)
	for _, sub := range subs {
		var user model.User
		if err := d.Client.First(&user, "id = ?", sub.UserID).Error; err != nil {
			continue
		}
		var aliasCount, recipientCount, credCount, sessCount int64
		d.Client.Model(&model.Alias{}).Where("user_id = ?", sub.UserID).Count(&aliasCount)
		d.Client.Model(&model.Recipient{}).Where("user_id = ?", sub.UserID).Count(&recipientCount)
		d.Client.Model(&model.Credential{}).Where("user_id = ?", sub.UserID).Count(&credCount)
		d.Client.Model(&model.Session{}).Where("user_id = ? AND expires_at > ?", sub.UserID, time.Now()).Count(&sessCount)
		var plan model.Plan
		if sub.PlanID != nil {
			if err := d.Client.First(&plan, "id = ?", *sub.PlanID).Error; err != nil {
				plan = model.Plan{}
			}
		}
		results = append(results, model.PlanUsage{
			UserID:          sub.UserID,
			Email:           user.Email,
			Tier:            sub.Tier,
			AliasCount:      aliasCount,
			MaxAliases:      int64(plan.MaxDailyAliases),
			RecipientCount:  recipientCount,
			MaxRecipients:   int64(plan.MaxRecipients),
			CredentialCount: credCount,
			MaxCredentials:  int64(plan.MaxCredentials),
			SessionCount:    sessCount,
			MaxSessions:     int64(plan.MaxSessions),
		})
	}
	return results, nil
}

func (d *Database) AdminGetInactiveAliases(ctx context.Context, days int) ([]model.InactiveAlias, error) {
	cutoff := time.Now().AddDate(0, 0, -days)
	var activeIDs []string
	d.Client.Model(&model.Message{}).Distinct("alias_id").Where("created_at >= ?", cutoff).Pluck("alias_id", &activeIDs)
	var aliases []model.Alias
	if len(activeIDs) > 0 {
		d.Client.Where("enabled = true AND id NOT IN ?", activeIDs).Order("created_at desc").Limit(100).Find(&aliases)
	} else {
		d.Client.Where("enabled = true").Order("created_at desc").Limit(100).Find(&aliases)
	}
	var results []model.InactiveAlias
	for _, a := range aliases {
		results = append(results, model.InactiveAlias{
			AliasID: a.ID, AliasName: a.Name, UserID: a.UserID,
			CreatedAt: a.CreatedAt, DaysInactive: days,
		})
	}
	return results, nil
}

func (d *Database) AdminGetCatchAllStats(ctx context.Context) (map[string]interface{}, error) {
	var total, catchAll int64
	d.Client.Model(&model.Alias{}).Count(&total)
	d.Client.Model(&model.Alias{}).Where("catch_all = true").Count(&catchAll)
	var byDomain []struct {
		Domain string
		Count  int64
	}
	d.Client.Model(&model.Alias{}).Select("substring_index(name, '@', -1) as domain, count(*) as count").Where("catch_all = true").Group("substring_index(name, '@', -1)").Order("count desc").Limit(10).Scan(&byDomain)
	var pct float64
	if total > 0 {
		pct = float64(catchAll) / float64(total) * 100
	}
	return map[string]interface{}{
		"total_aliases":    total,
		"catchall_aliases": catchAll,
		"percentage":       pct,
		"by_domain":        byDomain,
	}, nil
}

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

func (d *Database) AdminGetHourlyVolume(ctx context.Context, days int) ([]model.HourlyVolume, error) {
	cutoff := time.Now().AddDate(0, 0, -days)
	type row struct {
		Hour  int
		Count int64
	}
	var rows []row
	d.Client.Model(&model.Message{}).
		Select("hour(created_at) as hour, count(*) as count").
		Where("created_at >= ?", cutoff).
		Group("hour(created_at)").
		Order("hour asc").
		Scan(&rows)
	var results []model.HourlyVolume
	for _, r := range rows {
		results = append(results, model.HourlyVolume{Hour: r.Hour, Count: r.Count})
	}
	return results, nil
}

func (d *Database) AdminGetTopSenders(ctx context.Context, days int) ([]model.UserForwardStats, error) {
	cutoff := time.Now().AddDate(0, 0, -days)
	type row struct {
		UserID string
		Email  string
		Count  int64
	}
	var rows []row
	d.Client.Model(&model.Message{}).
		Select("messages.user_id, users.email, count(*) as count").
		Joins("join users on users.id = messages.user_id").
		Where("messages.created_at >= ? AND messages.type = 3", cutoff).
		Group("messages.user_id, users.email").
		Order("count desc").
		Limit(20).
		Scan(&rows)
	var results []model.UserForwardStats
	for _, r := range rows {
		results = append(results, model.UserForwardStats{UserID: r.UserID, Email: r.Email, Sends: r.Count})
	}
	return results, nil
}

func (d *Database) AdminGetAliasForwardStats(ctx context.Context, days int) ([]model.AliasForwardStats, error) {
	cutoff := time.Now().AddDate(0, 0, -days)
	type row struct {
		AliasID   string
		AliasName string
		UserEmail string
		Count     int64
		Type      int
	}
	var rows []row
	d.Client.Model(&model.Message{}).
		Select("messages.alias_id, aliases.name as alias_name, users.email as user_email, messages.type, count(*) as count").
		Joins("join aliases on aliases.id = messages.alias_id").
		Joins("join users on users.id = messages.user_id").
		Where("messages.created_at >= ?", cutoff).
		Group("messages.alias_id, aliases.name, users.email, messages.type").
		Order("count desc").Limit(200).Scan(&rows)
	result := map[string]*model.AliasForwardStats{}
	for _, r := range rows {
		s, ok := result[r.AliasID]
		if !ok {
			s = &model.AliasForwardStats{AliasID: r.AliasID, AliasName: r.AliasName, UserEmail: r.UserEmail}
			result[r.AliasID] = s
		}
		switch r.Type {
		case 0:
			s.Forwards = r.Count
		case 1:
			s.Blocks = r.Count
		case 2:
			s.Replies = r.Count
		case 3:
			s.Sends = r.Count
		}
	}
	var stats []model.AliasForwardStats
	for _, s := range result {
		stats = append(stats, *s)
	}
	return stats, nil
}
