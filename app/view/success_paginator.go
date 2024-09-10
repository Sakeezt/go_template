package view

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"go-template/domain"

	"github.com/gin-gonic/gin"
)

const xContentLength = "X-Content-Length"

type SuccessPaginatorResp struct {
	Code   int                `json:"code" example:"200"`
	Status string             `json:"status" example:"Get data success"`
	Data   dataListWithOption `json:"data" swaggerignore:"true"`
}

type dataListWithOption struct {
	Option   optionResp  `json:"option"`
	DataList interface{} `json:"lists" swaggertype:"object" swaggerignore:"true"`
}

type optionResp struct {
	Total       int           `json:"total" example:"1"`
	PerPage     int           `json:"per_page" example:"15"`
	CurrentPage int           `json:"current_page" example:"1"`
	TotalPage   int           `json:"total_page" example:"1"`
	Filters     []*filterData `json:"filters"`
	Search      []*searchData `json:"search"`
	Sorts       []*sortData   `json:"sorts"`
}

type filterData struct {
	Key       string `json:"key" example:"status"`
	Operation string `json:"operation" example:"eq"`
	Value     string `json:"value" example:"active"`
	Raw       string `json:"raw" example:"status:eq:active"`
}

type searchData struct {
	Key   string `json:"key" example:"name"`
	Value string `json:"value" example:"john"`
	Raw   string `json:"raw" example:"name:john"`
}

type sortData struct {
	Key   string `json:"key" example:"created_at"`
	Value string `json:"value" example:"desc"`
	Raw   string `json:"raw" example:"created_at:desc"`
}

func MakePaginatorResp(c *gin.Context, total int, opt *domain.PageOption, items interface{}) {
	status := http.StatusOK
	msg := MsgGetDataSuccess(c)
	totalPage := int(math.Ceil(float64(total) / float64(opt.PerPage)))

	if opt.Page > totalPage {
		status = http.StatusNoContent
		msg = MsgNotFoundData(c)
	}

	filters := make([]*filterData, 0)
	if len(opt.Filters) > 0 {
		for _, raw := range opt.Filters {
			if raw == "deleted_at:isNull" {
				continue
			}
			keyVal := strings.Split(raw, ":")
			var data filterData
			if len(keyVal) > 2 {
				data = filterData{
					Key:       keyVal[0],
					Operation: keyVal[1],
					Value:     keyVal[2],
					Raw:       raw,
				}
			} else {
				data = filterData{
					Key:       keyVal[0],
					Operation: keyVal[1],
					Raw:       raw,
				}
			}
			filters = append(filters, &data)
		}
	}
	search := make([]*searchData, 0)
	if len(opt.Search) > 0 {
		for _, raw := range opt.Search {
			keyVal := strings.Split(raw, ":")
			data := searchData{
				Key:   keyVal[0],
				Value: keyVal[1],
				Raw:   raw,
			}
			search = append(search, &data)
		}
	}
	sorts := make([]*sortData, 0)
	if len(opt.Sorts) > 0 {
		for _, raw := range opt.Sorts {
			keyVal := strings.Split(raw, ":")
			data := &sortData{
				Key:   keyVal[0],
				Value: keyVal[1],
				Raw:   raw,
			}
			sorts = append(sorts, data)
		}
	}

	c.Header(xContentLength, strconv.Itoa(total))
	c.JSON(status, SuccessPaginatorResp{
		Code:   status,
		Status: msg,
		Data: dataListWithOption{
			Option: optionResp{
				Total:       total,
				PerPage:     opt.PerPage,
				CurrentPage: opt.Page,
				TotalPage:   totalPage,
				Filters:     filters,
				Search:      search,
				Sorts:       sorts,
			},
			DataList: items,
		},
	})
}
