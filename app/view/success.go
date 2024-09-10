package view

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

const (
	CONTENT_LOCATION = "Content-Location"
)

type SuccessResp struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func makeSuccessResp(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, SuccessResp{
		Code:   code,
		Status: msg,
		Data:   data,
	})
}

func MakeCreateSuccessResp(c *gin.Context, id string) {
	c.Header(CONTENT_LOCATION, id)
	makeSuccessResp(c, http.StatusCreated, MsgCreateDataSuccess(c), id)
}

func MakeReadSuccessResp(c *gin.Context, data interface{}) {
	makeSuccessResp(c, http.StatusOK, MsgGetDataSuccess(c), data)
}

func MakeUpdateSuccessResp(c *gin.Context) {
	makeSuccessResp(c, http.StatusOK, MsgUpdateDataSuccess(c), reflect.TypeOf(struct{}{}))
}

func MakeDeleteSuccessResp(c *gin.Context) {
	makeSuccessResp(c, http.StatusOK, MsgDeleteDataSuccess(c), reflect.TypeOf(struct{}{}))
}
