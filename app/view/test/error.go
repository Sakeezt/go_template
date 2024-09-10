package test

import (
	"errors"
	"net/http"

	"go-template/app/view"
	"go-template/service/util"
)

func (suite *PackageTestSuite) TestMakeErrResp() {
	err := util.ConvertInputToDomainErr(errors.New("test"))
	view.MakeErrResp(suite.ctx, err)
	suite.Equal(http.StatusBadRequest, suite.ctx.Writer.Status())
}
