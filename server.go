package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/JIeeiroSst/store/component"
	"github.com/JIeeiroSst/store/graph"
	"github.com/JIeeiroSst/store/graph/generated"
	api "github.com/JIeeiroSst/store/handler"
	"github.com/JIeeiroSst/store/middleware"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router *gin.Engine
)

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func init() {
	adapter,_:=gormadapter.NewAdapterByDB(component.DB)

	router = gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))
	router.POST("/user/signup",api.SingUp)
	router.POST("/user/login", api.Login)
	resource := router.Group("/api")
	resource.Use(middleware.Authenticate())
	{
		//api rest full
		resource.GET("/resource", middleware.Authorize("resource", "read", adapter), api.ReadResource)
		resource.POST("/resource", middleware.Authorize("resource", "write", adapter), api.WriteResource)

		//api graphql
		router.POST("/query",middleware.Authorize("","",adapter), graphqlHandler())
		router.GET("/", middleware.Authorize("","",adapter),playgroundHandler())
	}

}

func main(){
	err := router.Run()
	if err != nil {
		log.Fatalln(fmt.Errorf("faild to start Gin application: %w", err))
	}
}