package test

import "time"

func (suite *PackageTestSuite) TestSet() {
	err := suite.redis.Set(suite.ctx, "test-set", "true")
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestSetWithTTL() {
	timeToExpired := time.Duration(1) * time.Second
	err := suite.redis.SetWithTTL(suite.ctx, "test-set-ttl", "true", timeToExpired)
	suite.NoError(err)
}
