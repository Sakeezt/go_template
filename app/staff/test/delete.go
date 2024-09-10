package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestDelete() {
	givenInput := suite.makeTestDeleteInput()
	req, resp := suite.makeDeleteRequest(givenInput)

	suite.service.On("Delete", mock.Anything, givenInput).Once().Return(nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *PackageTestSuite) TestDeleteServiceError() {
	givenInput := suite.makeTestDeleteInput()
	req, resp := suite.makeDeleteRequest(givenInput)

	suite.service.On("Delete", mock.Anything, givenInput).Once().Return(givenError)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}
