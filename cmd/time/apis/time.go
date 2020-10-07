package apis

import (
	"github.com/gin-gonic/gin"
	"time"
)

// GetTime godoc
// @Summary Retrieves time
// @Produce json
// @Param format path string true "Format"
// @Param tz path string true "TimeZone"
// @Success 200
// @Router /time/{id} [get]
func GetTime(c *gin.Context) {
	formatStr := c.DefaultQuery("format", "Mon Jan 2 15:04:05 -0700 MST 2006")
	tzStr := c.Query("tz")

	tz, err := time.LoadLocation(tzStr)
	if err != nil {
		panic(err)
	}
	curTime := time.Now().Round(0).In(tz).Format(formatStr)
	c.JSON(200, gin.H{
		"time": curTime,
	})
}