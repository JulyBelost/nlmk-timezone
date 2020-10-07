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

func newRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	return router
}

func testAPI(router *gin.Engine, method string, url string, parameters string, function gin.HandlerFunc, body string) *httptest.ResponseRecorder {
	router.Handle(method, url, function)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url+parameters, bytes.NewBufferString(body))
	router.ServeHTTP(res, req)
	return res
}

func runAPITests(t *testing.T, tests []apiTestCase) {
	for _, test := range tests {
		router := newRouter()
		res := testAPI(router, test.method, test.url, test.parameters, test.function, test.body)
		assert.Equal(t, test.status, res.Code, test.tag)
		assert.JSONEq(t, test.response, res.Body.String(), test.tag)
	}
}
