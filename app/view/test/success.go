package test

import (
	"net/http"

	"go-template/app/view"
)

type ItemStruct struct {
	Title string
	Body  string
}

func (suite *PackageTestSuite) TestMakeCreateSuccessResp() {
	view.MakeCreateSuccessResp(suite.ctx, "id")
	suite.Equal(http.StatusCreated, suite.ctx.Writer.Status())
}

func (suite *PackageTestSuite) TestMakeReadSuccessResp() {
	item := ItemStruct{
		Title: "",
		Body:  "",
	}
	view.MakeReadSuccessResp(suite.ctx, item)
	suite.Equal(http.StatusOK, suite.ctx.Writer.Status())
}

func (suite *PackageTestSuite) TestMakeUpdateSuccessResp() {
	view.MakeUpdateSuccessResp(suite.ctx)
	suite.Equal(http.StatusOK, suite.ctx.Writer.Status())
}

func (suite *PackageTestSuite) TestMakeDeleteSuccessResp() {
	view.MakeDeleteSuccessResp(suite.ctx)
	suite.Equal(http.StatusOK, suite.ctx.Writer.Status())
}
