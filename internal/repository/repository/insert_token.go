package repository

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"
)

func (r *CacheRepository) SetToken(ctx context.Context, key string, token string) error {
	expiration := 10 * time.Minute
	if err := r.super.Redis.Set(ctx, key, token, expiration).Err(); err != nil {
		r.logger.Error("redis set key error", zap.Error(err))
		return fmt.Errorf("failed to set key-value in redis: %w", err)
	}

	return nil
}
