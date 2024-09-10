package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"

	"go-template/service/staff/inout"
)

func (suite *PackageTestSuite) TestRead() {
	givenInput := suite.makeTestReadInput()
	req, resp := suite.makeReadRequest(givenInput)

	suite.service.On("Read", mock.Anything, givenInput).Once().Return(&inout.StaffView{}, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *PackageTestSuite) TestReadServiceError() {
	givenInput := suite.makeTestReadInput()
	req, resp := suite.makeReadRequest(givenInput)

	suite.service.On("Read", mock.Anything, givenInput).Once().Return(nil, givenError)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}
