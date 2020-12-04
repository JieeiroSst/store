package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/jieeiro/api/utils/authz"
)

func main()  {
	server:=gin.Default()

	e, _ :=casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")

	server.Use(authz.NewAuthorizer(e))

	server.Run()
}