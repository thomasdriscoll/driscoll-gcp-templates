package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thomasdriscoll/driscoll-gcp-templates/driscoll-gcp-golang-template/enums"
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

func TestPing(t *testing.T) {
	// constants
	route := enums.ADMIN_URL + enums.PING_URL

	//mocks
	adminController := AdminControllerImpl{}
	engine := http.NewServeMux()
	engine.HandleFunc(route, adminController.Ping)

	// *************************************************************************************************************
	// 				Requests and Responses
	// *************************************************************************************************************

	// 200
	okRequest, _ := http.NewRequest(http.MethodGet, route, nil)
	okResponse := enums.PING_MESSAGE

	testCases := []TestCase{
		{
			writer:               httptest.NewRecorder(),
			request:              okRequest,
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: []byte(okResponse),
			testMessage:          "/admin/ping is working and tested!",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			engine.ServeHTTP(testCase.writer, testCase.request)
			//assertions
			if testCase.expectedResponseCode != testCase.writer.Code {
				t.Errorf("Response code does not match expected status code")
			}
			fmt.Printf("%s", testCase.expectedResponseBody)
			fmt.Println()
			fmt.Printf("%s", testCase.writer.Body.Bytes())
			fmt.Println()
			if !bytes.Equal(testCase.expectedResponseBody, testCase.writer.Body.Bytes()) {
				t.Errorf("Response body does not match expected response")
			}
		})
	}
}
