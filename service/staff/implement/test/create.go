package test

import (
	"go-template/service/util"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestCreateSuccess() {
	givenInput := suite.makeTestCreateInput()

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.uuid.On("Generate").Once().Return(givenGenerateUUID)
	suite.repoStaff.On("Create", mock.Anything, givenStaff).Once().Return(givenID, nil)

	_, err := suite.service.Create(suite.ctx, givenInput)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestCreateValidateError() {
	givenInput := suite.makeTestCreateInput()

	suite.validator.On("Validate", givenInput).Once().Return(givenError)

	_, err := suite.service.Create(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.ValidationCreateErr))
}

func (suite *PackageTestSuite) TestCreateError() {
	givenInput := suite.makeTestCreateInput()

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.uuid.On("Generate").Once().Return(givenGenerateUUID)
	suite.repoStaff.On("Create", mock.Anything, givenStaff).Once().Return(givenID, givenError)

	_, err := suite.service.Create(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoCreateErr))
}
