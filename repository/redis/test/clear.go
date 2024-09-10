package test

func (suite *PackageTestSuite) TestClear() {
	err := suite.redis.Clear(suite.ctx, "test-clear1", "test-clear2")
	suite.NoError(err)
}
