package redis

import (
	"context"
	"encoding/json"
	"time"
)

func (r *Redis) Get(ctx context.Context, key string, dest interface{}) (err error) {
	cmd := r.Client.Get(ctx, key)
	if err = cmd.Err(); err != nil {
		return err
	}
	if dest != nil {
		return json.Unmarshal([]byte(cmd.Val()), dest)
	}
	return nil
}

func (r *Redis) GetTTL(ctx context.Context, key string) (result time.Duration, err error) {
	return r.Client.TTL(ctx, key).Result()
}
