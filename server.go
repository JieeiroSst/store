package main

import (
	"fmt"
	"github.com/JIeeiroSst/store/component"
	"github.com/JIeeiroSst/store/handler"
	"github.com/JIeeiroSst/store/middleware"
	gormadapter"github.com/casbin/gorm-adapter/v3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router *gin.Engine
)

func init() {
	adapter,_:=gormadapter.NewAdapterByDB(component.DB)

	router = gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))
	router.POST("/user/signup",handler.SingUp)
	router.POST("/user/login", handler.Login)
	resource := router.Group("/api")
	resource.Use(middleware.Authenticate())
	{
		resource.GET("/resource", middleware.Authorize("resource", "read", adapter), handler.ReadResource)
		resource.POST("/resource", middleware.Authorize("resource", "write", adapter), handler.WriteResource)
	}
}

func main(){
	err := router.Run()
	if err != nil {
		log.Fatalln(fmt.Errorf("faild to start Gin application: %w", err))
	}
}