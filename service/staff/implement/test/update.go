package test

import (
	"go-template/service/util"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestUpdateSuccess() {
	givenInput := suite.makeTestUpdateInput()

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.filterString.On("MakeIDFilters", mock.Anything).Once().Return(givenIDFilters)
	suite.repoStaff.On("Read", mock.Anything, givenIDFilters, givenStaff).Once().Return(nil)
	suite.repoStaff.On("Update", mock.Anything, givenIDFilters, givenStaff).Once().Return(nil)

	err := suite.service.Update(suite.ctx, givenInput)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestUpdateValidateError() {
	givenInput := suite.makeTestUpdateInput()

	suite.validator.On("Validate", givenInput).Once().Return(givenError)

	err := suite.service.Update(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.ValidationUpdateErr))
}

func (suite *PackageTestSuite) TestUpdateReadStaffError() {
	givenInput := suite.makeTestUpdateInput()

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.filterString.On("MakeIDFilters", mock.Anything).Once().Return(givenIDFilters)
	suite.repoStaff.On("Read", mock.Anything, givenIDFilters, givenStaff).Once().Return(givenError)

	err := suite.service.Update(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
}

func (suite *PackageTestSuite) TestUpdateError() {
	givenInput := suite.makeTestUpdateInput()

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.filterString.On("MakeIDFilters", mock.Anything).Once().Return(givenIDFilters)
	suite.repoStaff.On("Read", mock.Anything, givenIDFilters, givenStaff).Once().Return(nil)
	suite.repoStaff.On("Update", mock.Anything, givenIDFilters, givenStaff).Once().Return(givenError)

	err := suite.service.Update(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoUpdateErr))
}
