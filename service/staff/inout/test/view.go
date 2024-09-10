package test

import (
	"go-template/service/staff/inout"
)

func (suite *PackageTestSuite) TestStaffToView() {
	given := suite.makeTestStaff()
	actual := inout.StaffToView(given, suite.datetime)

	suite.Equal(given.ID, actual.ID)
	suite.Equal(given.Name, actual.Name)
	suite.Equal(given.Tel, actual.Tel)
	suite.Equal(actual.CreatedAt, suite.datetime.ConvertUnixToDateTime(suite.datetime.GetUnixNow()))
	suite.Equal(actual.UpdatedAt, suite.datetime.ConvertUnixToDateTime(suite.datetime.GetUnixNow()))
}
