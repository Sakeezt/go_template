package test

import (
	"context"

	"go-template/service/util/mocks"
	"go-template/service/validator"

	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	ctx       context.Context
	validator validator.Validator
	repoStaff *mocks.Repository
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.repoStaff = &mocks.Repository{}
	vRepo := validator.ValidatorRepository{
		Staff: suite.repoStaff,
	}
	suite.validator = validator.New(&vRepo)
}
