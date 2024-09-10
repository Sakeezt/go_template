package test

import (
	"context"
	"time"

	"go-template/config"
	"go-template/repository/redis"

	"github.com/stretchr/testify/suite"
)

var err error

type PackageTestSuite struct {
	suite.Suite
	ctx   context.Context
	redis *redis.Redis
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	conf := config.New()
	suite.redis, err = redis.New(suite.ctx, conf)
	suite.NoError(err)

	_ = suite.redis.SetWithTTL(suite.ctx, "test-get", "true", time.Duration(10)*time.Second)
	_ = suite.redis.SetWithTTL(suite.ctx, "test-clear1", "true", time.Duration(10)*time.Second)
	_ = suite.redis.SetWithTTL(suite.ctx, "test-clear2", "true", time.Duration(10)*time.Second)
}

func (suite *PackageTestSuite) TearDownSuite() {
	keys := []string{
		"test-get",
		"test-clear1",
		"test-clear2",
		"test-set",
		"test-set-ttl",
	}
	_ = suite.redis.Clear(suite.ctx, keys...)
}
