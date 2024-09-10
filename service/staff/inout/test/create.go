package test

func (suite *PackageTestSuite) TestCreateInputToStaffDomain() {
	given := suite.makeTestCreateInput()
	actual := given.ToDomain(suite.datetime)

	suite.Equal(given.ID, actual.ID)
	suite.Equal(given.Name, actual.Name)
	suite.Equal(given.Tel, actual.Tel)
}
