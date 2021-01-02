package middleware

import (
	"errors"
	"fmt"
	"github.com/JIeeiroSst/store/component"
	"github.com/JIeeiroSst/store/utils/jwt"
	"github.com/allegro/bigcache"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken:=c.Request.Header.Get("Authorization")
		if bearToken == ""{
			c.AbortWithStatusJSON(401,component.RestResponse{Message: "Authentication failure: Token not provided"})
			return
		}
		strArr := strings.Split(bearToken, " ")
		message,err:=jwt.ParseToken(strArr[1])
		if err!=nil{
			c.AbortWithStatusJSON(400,component.RestResponse{Message: message})
			return
		}
		sessionId, _ := c.Cookie("current_subject")
		sub, err := component.GlobalCache.Get(sessionId)
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			c.AbortWithStatusJSON(401, component.RestResponse{Message: "user hasn't logged in yet"})
			return
		}
		c.Set("current_subject", string(sub))
		c.Next()
	}
}

func Authorize(obj string, act string, adapter persist.Adapter) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, existed := c.Get("current_subject")
		if !existed {
			c.AbortWithStatusJSON(401, component.RestResponse{Message: "user hasn't logged in yet"})
			return
		}
		ok, err := enforce(val.(string), obj, act, adapter)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(500, component.RestResponse{Message: "error occurred when authorizing user"})
			return
		}
		if !ok {
			c.AbortWithStatusJSON(403, component.RestResponse{Message: "forbidden"})
			return
		}
		c.Next()
	}
}

func enforce(sub string, obj string, act string, adapter persist.Adapter) (bool, error) {
	enforcer, err := casbin.NewEnforcer("conf/rbac_model.conf", adapter)
	if err != nil {
		return false, fmt.Errorf("failed to create casbin enforcer: %w", err)
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	ok, err := enforcer.Enforce(sub, obj, act)
	return ok, err
}
