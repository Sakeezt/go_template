package test

import (
	"go-template/service/util/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	datetime *mocks.DateTime
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.datetime = &mocks.DateTime{}
}

func (suite *PackageTestSuite) SetupTest() {
	var now int64
	var dtString string
	suite.datetime.On("GetUnixNow").Return(now)
	suite.datetime.On("ConvertUnixToDateTime", mock.Anything).Return(dtString)
}
