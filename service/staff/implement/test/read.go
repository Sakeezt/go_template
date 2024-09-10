package test

import (
	"go-template/service/util"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestReadSuccess() {
	givenInput := suite.makeTestReadInput()

	suite.filterString.On("MakeID", mock.Anything).Once().Return(givenIDFilter)
	suite.filterString.On("MakeDeletedAtIsNull").Once().Return(givenDeletedAtIsNullFilter)
	suite.repoStaff.On("Read", mock.Anything, []string{givenIDFilter, givenDeletedAtIsNullFilter}, givenStaff).Once().Return(nil)

	_, err := suite.service.Read(suite.ctx, givenInput)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestReadError() {
	givenInput := suite.makeTestReadInput()

	suite.filterString.On("MakeID", mock.Anything).Once().Return(givenIDFilter)
	suite.filterString.On("MakeDeletedAtIsNull").Once().Return(givenDeletedAtIsNullFilter)
	suite.repoStaff.On("Read", mock.Anything, []string{givenIDFilter, givenDeletedAtIsNullFilter}, givenStaff).Once().Return(givenError)

	_, err := suite.service.Read(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
}
