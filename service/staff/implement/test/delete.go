package test

import (
	"go-template/service/util"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestDeleteSuccess() {
	givenInput := suite.makeTestDeleteInput()

	suite.filterString.On("MakeID", mock.Anything).Once().Return(givenIDFilter)
	suite.filterString.On("MakeDeletedAtIsNull").Once().Return(givenDeletedAtIsNullFilter)
	suite.repoStaff.On("Read", mock.Anything, []string{givenIDFilter, givenDeletedAtIsNullFilter}, givenStaff).Once().Return(nil)
	suite.repoStaff.On("SoftDelete", mock.Anything, []string{givenIDFilter, givenDeletedAtIsNullFilter}).Once().Return(nil)

	err := suite.service.Delete(suite.ctx, givenInput)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestDeleteReadStaffError() {
	givenInput := suite.makeTestDeleteInput()

	suite.filterString.On("MakeID", mock.Anything).Once().Return(givenIDFilter)
	suite.filterString.On("MakeDeletedAtIsNull").Once().Return(givenDeletedAtIsNullFilter)
	suite.repoStaff.On("Read", mock.Anything, []string{givenIDFilter, givenDeletedAtIsNullFilter}, givenStaff).Once().Return(givenError)

	err := suite.service.Delete(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
}

func (suite *PackageTestSuite) TestDeleteError() {
	givenInput := suite.makeTestDeleteInput()

	suite.filterString.On("MakeID", mock.Anything).Once().Return(givenIDFilter)
	suite.filterString.On("MakeDeletedAtIsNull").Once().Return(givenDeletedAtIsNullFilter)
	suite.repoStaff.On("Read", mock.Anything, []string{givenIDFilter, givenDeletedAtIsNullFilter}, givenStaff).Once().Return(nil)
	suite.repoStaff.On("SoftDelete", mock.Anything, []string{givenIDFilter, givenDeletedAtIsNullFilter}).Once().Return(givenError)

	err := suite.service.Delete(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoDeleteErr))
}
