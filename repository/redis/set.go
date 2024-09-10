package redis

import (
	"context"
	"encoding/json"
	"time"
)

func (r *Redis) Set(ctx context.Context, key string, value interface{}) (err error) {
	str, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Client.Set(ctx, key, str, time.Duration(r.Config.RedisTTL)*time.Minute).Err()
}

func (r *Redis) SetWithTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) (err error) {
	str, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Client.Set(ctx, key, str, ttl).Err()
}
