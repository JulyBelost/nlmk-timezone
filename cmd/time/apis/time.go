package apis

import (
	"github.com/gin-gonic/gin"
	"time"
)

// GetTime godoc
// @Summary Retrieves time
// @Produce json
// @Param format path string false "Format"
// @Param tz path string false "TimeZone"
// @Success 200
// @Router /time [get]
func GetTime(c *gin.Context) {
	var tz *time.Location
	var err error

	formatStr := c.DefaultQuery("format", "Mon Jan 2 15:04:05 -0700 MST 2006")
	tzStr := c.Query("tz")

	t := time.Now()
	if tzStr == "" {
		tz = t.Location()
	} else {
		tz, err = time.LoadLocation(tzStr)
		if err != nil {
			panic(err)
		}
	}

	curTime := time.Now().Round(0).In(tz).Format(formatStr)
	c.JSON(200, gin.H{
		"time": curTime,
	})
}