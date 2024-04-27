package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"strings"
)

func (r *CacheRepository) GetToken(ctx context.Context, key string) (string, error) {
	data, err := r.super.Redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			r.logger.Error("get redis trx_id error nil", zap.Error(err))
			return "", fmt.Errorf("key not found in redis: %w", err)
		}
		r.logger.Error("get redis trx_id error", zap.Error(err))
		return "", fmt.Errorf("failed to get key-value from redis: %w", err)
	}
	strSplit := strings.Split(data, `"`)
	if len(strSplit) > 1 {
		for _, v := range strSplit {
			if v != "" {
				data = v
			}
		}
	}

	return data, nil
}
