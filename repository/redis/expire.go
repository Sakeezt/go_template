package redis

import (
	"context"
	"time"
)

func (r *Redis) ExpireAt(ctx context.Context, key string, tm time.Time) (err error) {
	return r.Client.ExpireAt(ctx, key, tm).Err()
}
