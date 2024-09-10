package view

import (
	"net/http"

	"go-template/service/util"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type ErrorResp struct {
	Code   int        `json:"code"`
	Status string     `json:"status"`
	Errors []*ErrItem `json:"errors"`
}

type ErrItem struct {
	Cause   string `json:"cause" example:"Some Error Message (REPOSITORY:READ)"`
	Code    string `json:"code" example:"REPOSITORY"`
	SubCode string `json:"subCode" example:"READ"`
}

func MakeErrResp(c *gin.Context, err error) {
	code := getHTTPStatusCode(err)
	errResponse := &ErrorResp{}
	if code == http.StatusNotFound {
		errResponse = &ErrorResp{
			Code:   getHTTPStatusCode(err),
			Status: MsgNotFoundData(c),
			Errors: getRespErrors(err),
		}
		c.JSON(errResponse.Code, errResponse)
	} else {
		errResponse = &ErrorResp{
			Code:   getHTTPStatusCode(err),
			Status: MsgAttemptError(c),
			Errors: getRespErrors(err),
		}
		c.JSON(errResponse.Code, errResponse)
	}
}

func getHTTPStatusCode(err error) (code int) {
	switch err := util.TypeOfErr(err); {
	case err.IsType(util.ConvertInputToDomainErr):
		return http.StatusBadRequest
	case err.IsType(util.JsonDecoderErr):
		return http.StatusBadRequest
	case err.IsType(util.QueryStrDecoderErr):
		return http.StatusBadRequest
	case err.IsType(util.ValidationCreateErr):
		return http.StatusUnprocessableEntity
	case err.IsType(util.ValidationUpdateErr):
		return http.StatusUnprocessableEntity
	case err.IsType(util.ValidationParamOptionErr):
		return http.StatusUnprocessableEntity
	case err.IsType(util.ValidationHeaderErr):
		return http.StatusUnprocessableEntity
	case err.IsType(util.RepoCreateErr):
		return http.StatusUnprocessableEntity
	case err.IsType(util.RepoUpdateErr):
		return http.StatusUnprocessableEntity
	case err.IsType(util.RepoReadErr):
		return http.StatusNoContent
	case err.IsType(util.RepoListErr):
		return http.StatusNoContent
	default:
		return http.StatusInternalServerError
	}
}

func getRespErrors(err error) (errs []*ErrItem) {
	switch err.(type) {
	case *util.Error:
		return utilToResp(err.(*util.Error))
	default:
		ukErr := util.UnknownErr(err)
		return []*ErrItem{
			{
				Cause:   ukErr.Error(),
				Code:    ukErr.Code,
				SubCode: ukErr.SubCode,
			},
		}
	}
}

func utilToResp(err *util.Error) (errs []*ErrItem) {
	switch err := util.TypeOfErr(err); {
	case err.IsType(util.ValidationCreateErr):
		return validateToResp(err)
	case err.IsType(util.ValidationUpdateErr):
		return validateToResp(err)
	case err.IsType(util.ValidationParamOptionErr):
		return validateToResp(err)
	default:
		return []*ErrItem{
			{
				Cause:   err.Error(),
				Code:    err.Code,
				SubCode: err.SubCode,
			},
		}
	}
}

func validateToResp(err *util.Error) (errs []*ErrItem) {
	vErrs := err.Cause.(validator.ValidationErrors)
	errs = make([]*ErrItem, len(vErrs))
	for i, vErr := range vErrs {
		errs[i] = &ErrItem{
			Cause:   vErr.Translate(nil),
			Code:    vErr.Tag(),
			SubCode: vErr.Field(),
		}
	}
	return errs
}
