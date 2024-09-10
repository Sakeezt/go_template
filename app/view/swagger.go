package view

type SuccessReadResp struct {
	Code   int         `json:"code" example:"200"`
	Status string      `json:"status" example:"Get data success"`
	Data   interface{} `json:"data" swaggertype:"object" swaggerignore:"true"`
}

type SuccessCreateResp struct {
	Code   int         `json:"code" example:"201"`
	Status string      `json:"status" example:"Create data success"`
	Data   interface{} `json:"data" swaggertype:"object" swaggerignore:"true"`
}

type SuccessUpdateResp struct {
	Code   int         `json:"code" example:"200"`
	Status string      `json:"status" example:"Update data success"`
	Data   interface{} `json:"data" swaggertype:"object" swaggerignore:"true"`
}

type SuccessDeleteResp struct {
	Code   int         `json:"code" example:"200"`
	Status string      `json:"status" example:"Delete data success"`
	Data   interface{} `json:"data" swaggertype:"object" swaggerignore:"true"`
}

type Error204Resp struct {
	Code   int    `json:"code" example:"204"`
	Status string `json:"status" example:"No content"`
}

type Error400Resp struct {
	Code   int    `json:"code" example:"400"`
	Status string `json:"status" example:"Bad request"`
}

type Error404Resp struct {
	Code   int    `json:"code" example:"404"`
	Status string `json:"status" example:"Not found"`
}

type Error422Resp struct {
	Code   int    `json:"code" example:"422"`
	Status string `json:"status" example:"Attempt error"`
}

type Error500Resp struct {
	Code   int    `json:"code" example:"500"`
	Status string `json:"status" example:"Internal server error"`
}

type ErrorCreateItem struct {
	Cause   string `json:"cause" example:"Some Error Message (REPOSITORY:CREATE)"`
	Code    string `json:"code" example:"REPOSITORY"`
	SubCode string `json:"subCode" example:"CREATE"`
}

type ErrorReadItem struct {
	Cause   string `json:"cause" example:"Some Error Message (REPOSITORY:READ)"`
	Code    string `json:"code" example:"REPOSITORY"`
	SubCode string `json:"subCode" example:"READ"`
}

type ErrorUpdateItem struct {
	Cause   string `json:"cause" example:"Some Error Message (REPOSITORY:UPDATE)"`
	Code    string `json:"code" example:"REPOSITORY"`
	SubCode string `json:"subCode" example:"UPDATE"`
}

type ErrorDeleteItem struct {
	Cause   string `json:"cause" example:"Some Error Message (REPOSITORY:DELETE)"`
	Code    string `json:"code" example:"REPOSITORY"`
	SubCode string `json:"subCode" example:"DELETE"`
}
