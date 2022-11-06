package controller

import (
	"net/http"
	"net/http/httptest"
)

// global vars
type TestCase struct {
	writer               *httptest.ResponseRecorder
	request              *http.Request
	expectedResponseCode int
	expectedResponseBody []byte
	testMessage          string
}

// *************************************************************************************************************
// 				TESTS
// *************************************************************************************************************
