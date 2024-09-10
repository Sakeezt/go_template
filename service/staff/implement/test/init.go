package test

import (
	"context"
	"errors"

	"go-template/domain"
	"go-template/service/staff"
	"go-template/service/staff/implement"
	logMocks "go-template/service/util/logs/mocks"
	"go-template/service/util/mocks"
	validatorMocks "go-template/service/validator/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	givenStaff *domain.Staff

	givenError                 error
	givenIDFilters             []string
	givenIDFilter              string
	givenDeletedAtIsNullFilter string
	givenTotal                 int
	givenItems                 interface{}
	givenID                    string
	givenGenerateUUID          string
)

type PackageTestSuite struct {
	suite.Suite
	ctx          context.Context
	validator    *validatorMocks.Validator
	repoStaff    *mocks.Repository
	uuid         *mocks.UUID
	datetime     *mocks.DateTime
	filterString *mocks.FilterString
	log          *logMocks.Log
	service      staff.Service
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.validator = &validatorMocks.Validator{}
	suite.repoStaff = &mocks.Repository{}

	suite.uuid = &mocks.UUID{}
	suite.datetime = &mocks.DateTime{}
	suite.filterString = &mocks.FilterString{}
	suite.log = &logMocks.Log{}

	suite.service = implement.New(&implement.StaffServiceConfig{
		Validator:    suite.validator,
		RepoStaff:    suite.repoStaff,
		UUID:         suite.uuid,
		DateTime:     suite.datetime,
		FilterString: suite.filterString,
		Log:          suite.log,
	})
}

func (suite *PackageTestSuite) SetupTest() {
	givenStaff = &domain.Staff{}
	givenTotal = 1
	givenError = errors.New("some error message")

	suite.datetime.On("GetUnixNow").Return(int64(0))
	suite.datetime.On("ConvertUnixToDateTime", mock.Anything).Return("")
}
