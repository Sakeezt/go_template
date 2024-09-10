package test

import (
	"context"

	"go-template/config"
	"go-template/service/util"
	"go-template/service/validator"

	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	ctx       context.Context
	uuid      util.UUID
	validator validator.Validator
	datetime  util.DateTime
}

func (suite *PackageTestSuite) SetupSuite() {
	var err error
	suite.ctx = context.Background()
	suite.uuid, err = util.NewUUID()
	suite.NoError(err)
	suite.datetime = util.NewDateTime(&config.Config{AppTimeZone: "Asia/Bangkok"})
}
