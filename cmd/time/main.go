package main

import (
    "github.com/JulyBelost/nlmk-timezone/cmd/time/apis"
    _ "github.com/JulyBelost/nlmk-timezone/cmd/time/docs"
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
)

// @title Time Swagger API
// @version 1.0
// @description Swagger API for Golang Project Time.
// @termsOfService http://swagger.io/terms/

// @BasePath /api/v1
func main() {
    r := gin.New()

    // gin.DefaultWriter = os.Stdout
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    v1 := r.Group("/api/v1")
    {
        v1.GET("/time", apis.GetTime)
    }

    r.Run(":8080") // listen and serve on localhost
}
