package test

import (
	"errors"
	"net/http/httptest"

	"go-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"go-template/app/staff"
	"go-template/service/staff/mocks"
)

var (
	givenError error
	opt        *domain.PageOption
)

type PackageTestSuite struct {
	suite.Suite
	router  *gin.Engine
	ctx     *gin.Context
	ctrl    *staff.Controller
	service *mocks.Service
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.service = &mocks.Service{}
	suite.ctrl = staff.New(suite.service)
	suite.ctx, suite.router = gin.CreateTestContext(httptest.NewRecorder())

	suite.router.GET("/staffs", suite.ctrl.List)
	suite.router.POST("/staffs", suite.ctrl.Create)
	suite.router.GET("/staffs/:id", suite.ctrl.Read)
	suite.router.PUT("/staffs/:id", suite.ctrl.Update)
	suite.router.DELETE("/staffs/:id", suite.ctrl.Delete)
}

func (suite *PackageTestSuite) SetupTest() {
	givenError = errors.New("some error message")
	opt = &domain.PageOption{
		Page:    1,
		PerPage: 15,
		Filters: []string{"deleted_at:isNull"},
		Sorts:   []string{"created_at:desc"},
	}
}
