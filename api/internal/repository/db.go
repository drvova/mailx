package repository

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"ivpn.net/email/api/config"
	"ivpn.net/email/api/internal/model"
)

type Database struct {
	Client *gorm.DB
}

func NewDB(cfg config.DBConfig) (*Database, error) {
	db, err := connect(cfg)
	if err != nil {
		return nil, err
	}

	err = migrate(db)
	if err != nil {
		return nil, err
	}

	return &Database{
		Client: db,
	}, nil
}

func (d *Database) Close() error {
	db, err := d.Client.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func connect(cfg config.DBConfig) (*gorm.DB, error) {
	gormCfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	host := ""
	if len(cfg.Hosts) > 0 && cfg.Hosts[0] != "" {
		host = cfg.Hosts[0]
	}

	dsn := cfg.User + ":" + cfg.Password + "@tcp(" + host + ":" + cfg.Port + ")/" + cfg.Name + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), gormCfg)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxIdleTime(time.Hour)
	sqlDB.SetConnMaxLifetime(24 * time.Hour)

	log.Println("DB connection OK")
	return db, nil
}

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.User{},
		&model.Subscription{},
		&model.Recipient{},
		&model.Alias{},
		&model.Message{},
		&model.InboxMessage{},
		&model.Settings{},
		&model.Session{},
		&model.Credential{},
		&model.Log{},
		&model.AccessKey{},
		&model.Domain{},
		&model.Plan{},
		&model.AdminAudit{},
	)
	if err != nil {
		return err
	}

	log.Println("DB migration OK")
	return nil
}
