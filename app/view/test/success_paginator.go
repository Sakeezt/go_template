package test

import (
	"net/http"

	"go-template/app/view"
	"go-template/domain"
)

func (suite *PackageTestSuite) TestMakePaginatorResp() {
	items := []ItemStruct{
		{
			Title: "Test",
			Body:  "Test",
		},
	}
	opt := &domain.PageOption{
		Page:    1,
		PerPage: 1,
		Filters: nil,
		Sorts:   nil,
	}
	view.MakePaginatorResp(suite.ctx, len(items), opt, items)
	suite.Equal(http.StatusOK, suite.ctx.Writer.Status())
}

func (suite *PackageTestSuite) TestMakePaginatorNoContentResp() {
	items := []ItemStruct{}
	opt := &domain.PageOption{
		Page:    1,
		PerPage: 1,
		Filters: nil,
		Sorts:   nil,
	}
	view.MakePaginatorResp(suite.ctx, len(items), opt, items)
	suite.Equal(http.StatusNoContent, suite.ctx.Writer.Status())
}
