package test

func (suite *PackageTestSuite) TestUpdateInputToStaffDomain() {
	given := suite.makeTestUpdateInput()
	actual := given.ToDomain(suite.datetime)

	suite.Equal(given.Name, actual.Name)
	suite.Equal(given.Tel, actual.Tel)
	suite.Equal(actual.UpdatedAt, suite.datetime.GetUnixNow())
}
