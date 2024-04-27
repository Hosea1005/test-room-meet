package repository

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
)

func (r *CacheRepository) DeleteToken(ctx context.Context, key string) error {
	log.Println("initiate delete key: ", key)
	if err := r.super.Redis.Del(ctx, key).Err(); err != nil {
		r.logger.Error("redis error delete key", zap.Error(err))
		return fmt.Errorf("failed to set key-value in redis: %w", err)
	}

	return nil
}
