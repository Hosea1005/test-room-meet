package config

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"room-meet/internal/model"
	"strconv"
)

type BaseRepository struct {
	PostgresSql *gorm.DB
	MySql       *sql.DB
	Redis       *redis.Client

	ctx context.Context
}

func Redis(ctx context.Context, logger *zap.Logger) *redis.Client {
	logger.Info("initializing redis connection")

	dbRedis, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		logger.Error("Failed get redis db from env", zap.Error(err))
		dbRedis = 0
	}

	redis := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbRedis,
	})

	logger.Info("Redis Connection", zap.Any("PING", redis.Ping(ctx)))

	return redis
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	SSLMode  string
}

// NewDB creates a new database connection, migrates the schema, and returns a pointer to the GORM DB instance.
func PostgresDB(logger *zap.Logger) *gorm.DB {
	config := DatabaseConfig{
		Host:     os.Getenv("DB_POSTGRES_HOST"),
		User:     os.Getenv("DB_POSTGRES_USERNAME"),
		Password: os.Getenv("DB_POSTGRES_PASSWORD"),
		DBName:   os.Getenv("DB_POSTGRES_NAME"),
		Port:     os.Getenv("DB_POSTGRES_PORT"),
		SSLMode:  os.Getenv("DB_POSTGRES_SSL"),
	}
	// Build the connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.Host, config.Port, config.User, config.DBName, config.Password, config.SSLMode)

	// Open a new database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("error connect to postgres : ", zap.String("error", err.Error()))
		return nil
	}

	// Auto-migrate the schema
	if err := db.Migrator().AutoMigrate(&model.Room{}); err != nil {
		logger.Error("AutoMigrate error", zap.Error(err))
		return nil
	}

	return db
}
