package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"

	"go-template/service/staff/inout"
)

func (suite *PackageTestSuite) TestList() {
	req, resp := suite.makeListRequest()

	suite.service.On("List", mock.Anything, opt).Once().Return(0, []*inout.StaffView{}, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusNoContent, resp.Code)
}

func (suite *PackageTestSuite) TestListInvalidQueryParam() {
	req, resp := suite.makeListRequestInvalidQueryParam()

	suite.service.On("List", mock.Anything, nil).Once()
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}

func (suite *PackageTestSuite) TestListServiceError() {
	req, resp := suite.makeListRequest()

	suite.service.On("List", mock.Anything, opt).Once().Return(0, nil, givenError)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}
