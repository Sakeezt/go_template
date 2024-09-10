package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestUpdate() {
	givenInput := suite.makeTestUpdateInput()
	req, resp := suite.makeUpdateRequest(givenInput)

	suite.service.On("Update", mock.Anything, givenInput).Once().Return(nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *PackageTestSuite) TestUpdateInvalidJSON() {
	req, resp := suite.makeUpdateRequestInvalidJSON("test")

	suite.service.On("Update", mock.Anything, nil).Once()
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}

func (suite *PackageTestSuite) TestUpdateServiceError() {
	givenInput := suite.makeTestUpdateInput()
	req, resp := suite.makeUpdateRequest(givenInput)

	suite.service.On("Update", mock.Anything, givenInput).Once().Return(givenError)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}
