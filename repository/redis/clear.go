package redis

import (
	"context"
)

func (r *Redis) Clear(ctx context.Context, keys ...string) (err error) {
	return r.Client.Del(ctx, keys...).Err()
}
