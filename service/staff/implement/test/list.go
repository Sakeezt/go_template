package test

import (
	"go-template/domain"
	"go-template/service/util"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestListSuccess() {
	givenPageOption := &domain.PageOption{}
	givenItems = []interface{}{
		&domain.Staff{},
	}

	suite.validator.On("Validate", givenPageOption).Once().Return(nil)
	suite.repoStaff.On("List", mock.Anything, givenPageOption, &domain.Staff{}).Once().Return(givenTotal, givenItems, nil)

	_, _, err := suite.service.List(suite.ctx, givenPageOption)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestListValidateError() {
	givenPageOption := &domain.PageOption{}

	suite.validator.On("Validate", givenPageOption).Once().Return(givenError)

	_, _, err := suite.service.List(suite.ctx, givenPageOption)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.ValidationParamOptionErr))
}

func (suite *PackageTestSuite) TestListError() {
	givenPageOption := &domain.PageOption{}
	givenItems = []interface{}{
		&domain.Staff{},
	}

	suite.validator.On("Validate", givenPageOption).Once().Return(nil)
	suite.repoStaff.On("List", mock.Anything, givenPageOption, &domain.Staff{}).Once().Return(0, nil, givenError)

	_, _, err := suite.service.List(suite.ctx, givenPageOption)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoListErr))
}
