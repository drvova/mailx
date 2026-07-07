package repository

import (
	"context"

	"ivpn.net/email/api/internal/model"
)

func (d *Database) PostInboxMessage(ctx context.Context, msg model.InboxMessage) error {
	err := d.Client.WithContext(ctx).Create(&msg).Error
	if err != nil {
		return err
	}

	// Evict oldest messages beyond the per-alias cap.
	var staleIDs []uint
	err = d.Client.WithContext(ctx).Model(&model.InboxMessage{}).
		Where("alias_id = ?", msg.AliasID).
		Order("id DESC").
		Limit(1000).
		Offset(model.MaxInboxMessagesPerAlias).
		Pluck("id", &staleIDs).Error
	if err != nil || len(staleIDs) == 0 {
		return err
	}

	return d.Client.WithContext(ctx).Delete(&model.InboxMessage{}, staleIDs).Error
}

func (d *Database) GetInboxMessages(ctx context.Context, aliasID string, userID string) ([]model.InboxMessage, error) {
	var messages []model.InboxMessage
	err := d.Client.WithContext(ctx).
		Select("id", "created_at", "user_id", "alias_id", "`from`", "from_name", "subject", "`read`", "size").
		Where("alias_id = ? AND user_id = ?", aliasID, userID).
		Order("created_at DESC").
		Find(&messages).Error
	return messages, err
}

func (d *Database) GetInboxMessage(ctx context.Context, ID uint, userID string) (model.InboxMessage, error) {
	var message model.InboxMessage
	err := d.Client.WithContext(ctx).
		Where("id = ? AND user_id = ?", ID, userID).
		First(&message).Error
	if err != nil {
		return model.InboxMessage{}, err
	}

	if !message.Read {
		err = d.Client.WithContext(ctx).Model(&model.InboxMessage{}).
			Where("id = ?", ID).Update("read", true).Error
	}

	return message, err
}

func (d *Database) DeleteInboxMessage(ctx context.Context, ID uint, userID string) error {
	return d.Client.WithContext(ctx).
		Where("id = ? AND user_id = ?", ID, userID).
		Delete(&model.InboxMessage{}).Error
}
