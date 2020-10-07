package apis

import (
    "net/http"
    "testing"
    "time"
)


func TestTime(t *testing.T) {
    tz, _ := time.LoadLocation("Australia/Sydney")

    runAPITests(t, []apiTestCase{
        {"t1 - get Time with format", "GET", "/time", "?format=2006-01-02 15:04", "",
            GetTime, http.StatusOK,
            "{\"time\":\"" + time.Now().Format("2006-01-02 15:04") + "\"}"},
        {"t2 - get Time with tz", "GET", "/time", "?tz=Australia/Sydney", "",
            GetTime, http.StatusOK,
            "{\"time\":\"" + time.Now().In(tz).Format("Mon Jan 2 15:04:05 -0700 MST 2006") + "\"}"},
        {"t3 - get Time without params", "GET", "/time", "", "",
            GetTime, http.StatusOK,
            "{\"time\":\"" + time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006") + "\"}"},
    })
}
