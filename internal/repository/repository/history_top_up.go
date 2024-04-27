package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"room-meet/domain/entity"
)

func (r *CacheRepository) Set(ctx context.Context, key string, value *entity.Room) error {
	data, err := json.Marshal(value)
	if err != nil {
		//r.logger.Error("redis set key error marshal", zap.Error(err))
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	if err := r.super.Redis.Set(ctx, key, data, 0).Err(); err != nil {
		r.logger.Error("redis set key error", zap.Error(err))
		return fmt.Errorf("failed to set key-value in redis: %w", err)
	}

	return nil
}
