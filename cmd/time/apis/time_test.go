package apis

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func init() {
	timeNow = func() time.Time {
		t, _ := time.Parse("2006-01-02 15:04:05", "2017-01-20 01:02:03")
		return t
	}
}

func TestTime(t *testing.T) {
	fmt.Println(timeNow())

	runAPITests(t, []apiTestCase{
		{"t1 - get Time", "GET", "/time", "/time", "", GetTime, http.StatusOK, ""},
	})
}
