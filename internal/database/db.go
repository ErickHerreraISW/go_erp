package database

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Option func(*gorm.Config)

func New(dbURL string) (*gorm.DB, error) {
	cfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	dial := postgres.Open(dbURL)
	db, err := gorm.Open(dial, cfg)
	if err != nil {
		return nil, err
	}

	// Ping básico
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, err
	}
	log.Info().Msg("Connected to PostgreSQL")
	return db, nil
}
