package redis

import (
	"context"

	"go-template/config"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Client *redis.Client
	Config *config.Config
}

func New(ctx context.Context, appConfig *config.Config) (r *Redis, err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     appConfig.RedisEndpoint,
		Password: appConfig.RedisPassword,
		DB:       0,
	})

	retry := 3
	for true {
		_, err = client.Ping(ctx).Result()
		if err != nil {
			retry--
			if retry < 0 {
				return nil, err
			}
			continue
		}
		break
	}

	r = &Redis{
		Client: client,
		Config: appConfig,
	}

	return r, nil
}
