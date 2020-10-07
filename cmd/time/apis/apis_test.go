package apis

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type apiTestCase struct {
	tag              string
	method           string
	url              string
	parameters       string
	body             string
	function         gin.HandlerFunc
	status           int
	response         string
}

// Creates new router in testing mode
func newRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	return router
}

// Used to run single API test case. It makes HTTP request and returns its response
func testAPI(router *gin.Engine, method string, url string, parameters string, function gin.HandlerFunc, body string) *httptest.ResponseRecorder {
	router.Handle(method, url, function)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url+parameters, bytes.NewBufferString(body))
	router.ServeHTTP(res, req)
	return res
}

// Checks JSON response is same as expected data in test case file.
// All test expected test case responses are stored in `test_data/test_case_data` folder in format `<suite_name>_t<number>.json`
func runAPITests(t *testing.T, tests []apiTestCase) {
	for _, test := range tests {
		router := newRouter()
		res := testAPI(router, test.method, test.url, test.parameters, test.function, test.body)
		assert.Equal(t, test.status, res.Code, test.tag)
		assert.JSONEq(t, test.response, res.Body.String(), test.tag)
	}
}
