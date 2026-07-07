package jobs

import (
	"log"

	"gorm.io/gorm"
	"ivpn.net/email/api/internal/model"
)

// DeleteExpiredInboxAliases soft-deletes inbox aliases past their TTL.
// CleanupDeletedAliases purges the rows later.
func DeleteExpiredInboxAliases(db *gorm.DB) {
	err := db.Where("type = ? AND expires_at IS NOT NULL AND expires_at < NOW()", model.AliasInbox).
		Delete(&model.Alias{}).Error
	if err != nil {
		log.Println("Error deleting expired inbox aliases:", err)
		return
	}
}

// DeleteOldInboxMessages hard-deletes stored mail older than 7 days and
// sweeps messages orphaned by soft-deleted aliases.
func DeleteOldInboxMessages(db *gorm.DB) {
	err := db.Where("created_at < NOW() - INTERVAL ? DAY", 7).
		Delete(&model.InboxMessage{}).Error
	if err != nil {
		log.Println("Error deleting old inbox messages:", err)
		return
	}

	err = db.Where("alias_id IN (SELECT id FROM aliases WHERE deleted_at IS NOT NULL)").
		Delete(&model.InboxMessage{}).Error
	if err != nil {
		log.Println("Error deleting orphaned inbox messages:", err)
		return
	}
}
