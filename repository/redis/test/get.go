package test

func (suite *PackageTestSuite) TestGet() {
	var value string
	err := suite.redis.Get(suite.ctx, "test-get", &value)
	suite.NoError(err)
	suite.NotEmpty(value)
}
