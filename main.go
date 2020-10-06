package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/time", func(c *gin.Context) {
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
	})
	r.Run(":80") // listen and serve on localhost
}