package main

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	adminGraph "github.com/JIeeiroSst/store/graph/admin/graph"
	adminGenerated "github.com/JIeeiroSst/store/graph/admin/graph/generated"
	clientGraph"github.com/JIeeiroSst/store/graph/client/graph"
	clientGenerated"github.com/JIeeiroSst/store/graph/client/graph/generated"
	api "github.com/JIeeiroSst/store/handler"
	"github.com/JIeeiroSst/store/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router *gin.Engine
)

func graphqlHandlerAdmin() gin.HandlerFunc {
	h := handler.NewDefaultServer(adminGenerated.NewExecutableSchema(adminGenerated.Config{Resolvers: &adminGraph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func graphqlHandlerClient() gin.HandlerFunc {
	h := handler.NewDefaultServer(clientGenerated.NewExecutableSchema(clientGenerated.Config{Resolvers: &clientGraph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func init() {
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
		resource.POST("/admin/graphql",middleware.AuthorizeAdmin("private"), graphqlHandlerAdmin())
		resource.POST("/graphql",middleware.AuthorizeClient("public"), graphqlHandlerClient())
	}

}

func main(){
	err := router.Run()
	if err != nil {
		log.Fatalln(fmt.Errorf("faild to start Gin application: %w", err))
	}
}