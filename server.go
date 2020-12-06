package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/jieeiro/api/commands"
	"github.com/jieeiro/api/utils/authz"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
)

func main()  {
	server:=gin.Default()

	a,err:=gormadapter.NewAdapter("postgresql","postgresql://root:@1234:5432/hospital?sslmode=disable")
	if err!=nil{
		log.Println(err)
	}

	e, _ :=casbin.NewEnforcer("authz_model.conf", a)

	server.Use(authz.NewAuthorizer(e))

	commands.InitCommand()

	server.Run()
}