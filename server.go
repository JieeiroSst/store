package main

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/jieeiro/api/commands"
	"github.com/jieeiro/api/utils/authz"
	"log"
)

func main()  {
	server:=gin.Default()

	adapter,err:=gormadapter.NewAdapter("postgresql","postgresql://root:@1234:5432/hospital?sslmode=disable")
	if err!=nil{
		log.Println(err)
	}

	enforcer, _ :=casbin.NewEnforcer("authz_model.conf", adapter)

	server.Use(authz.NewAuthorizer(enforcer))

	_ = enforcer.LoadPolicy()

	commands.InitCommand()
	_ = enforcer.SavePolicy()

	// Do permission checking
	//enforcer.Enforce("alice", "data1", "write")
	// Do some mutations
	//enforcer.AddPolicy("alice", "data2", "write")
	//enforcer.RemovePolicy("alice", "data1", "write")

	_ = server.Run()
}