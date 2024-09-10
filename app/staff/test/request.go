package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"go-template/service/staff/inout"
)

const (
	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_PUT    = "PUT"
	METHOD_DELETE = "DELETE"
	CONTENT_TYPE  = "Content-Type"
	CONTENT_VALUE = "application/json"
)

func (suite *PackageTestSuite) makeCreateRequest(input *inout.StaffCreateInput) (req *http.Request, w *httptest.ResponseRecorder) {
	jsonBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest(METHOD_POST, "/staffs", bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder()
}

func (suite *PackageTestSuite) makeCreateRequestInvalidJSON() (req *http.Request, w *httptest.ResponseRecorder) {
	jsonBytes := []byte("{{{}}}")
	req, _ = http.NewRequest(METHOD_POST, "/staffs", bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder()
}

func (suite *PackageTestSuite) makeReadRequest(input *inout.StaffReadInput) (req *http.Request, w *httptest.ResponseRecorder) {
	req, _ = http.NewRequest(METHOD_GET, fmt.Sprintf("/staffs/%s", input.ID), nil)
	return req, httptest.NewRecorder()
}

func (suite *PackageTestSuite) makeUpdateRequest(input *inout.StaffUpdateInput) (req *http.Request, w *httptest.ResponseRecorder) {
	jsonBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest(METHOD_PUT, fmt.Sprintf("/staffs/%s", input.ID), bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder()
}

func (suite *PackageTestSuite) makeUpdateRequestInvalidJSON(id string) (req *http.Request, w *httptest.ResponseRecorder) {
	jsonBytes := []byte("{{{}}}")
	req, _ = http.NewRequest(METHOD_PUT, fmt.Sprintf("/staffs/%s", id), bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder()
}

func (suite *PackageTestSuite) makeDeleteRequest(input *inout.StaffDeleteInput) (req *http.Request, w *httptest.ResponseRecorder) {
	jsonBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest(METHOD_DELETE, fmt.Sprintf("/staffs/%s", input.ID), bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder()
}

func (suite *PackageTestSuite) makeListRequest() (req *http.Request, w *httptest.ResponseRecorder) {
	req, _ = http.NewRequest(METHOD_GET, "/staffs?page=1&per_page=15", nil)
	return req, httptest.NewRecorder()
}

func (suite *PackageTestSuite) makeListRequestInvalidQueryParam() (req *http.Request, w *httptest.ResponseRecorder) {
	req, _ = http.NewRequest(METHOD_GET, "/staffs?page={{}}&per_page={{}}", nil)
	return req, httptest.NewRecorder()
}
