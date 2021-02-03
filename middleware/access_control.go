package middleware

import (
	"github.com/JIeeiroSst/store/component"
	"github.com/JIeeiroSst/store/utils/jwt"
	"github.com/gin-gonic/gin"
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
		if err != nil {
			c.AbortWithStatusJSON(401, component.RestResponse{Message: "user hasn't logged in yet"})
			return
		}
		c.Set("current_subject", string(sub))
		c.Next()
	}
}

func AuthorizeAdmin(parameters  string) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, existed := c.Get("current_subject")
		if !existed {
			c.AbortWithStatusJSON(401, component.RestResponse{Message: "user hasn't logged in yet"})
			return
		}
		token:=c.Request.Header.Get("Authorization")
		strArr := strings.Split(token, " ")
		username,permission :=jwt.GetPermission(strArr[1])
		if username == "admin" {
			c.Next()
		}
		if parameters!=permission && permission != "private" && username != "admin"  {
			c.AbortWithStatusJSON(401, component.RestResponse{Message: "the user does not have admin access to access "})
			return
		}

		c.Next()
	}
}

func AuthorizeClient(parameters  string) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, existed := c.Get("current_subject")
		if !existed {
			c.AbortWithStatusJSON(401, component.RestResponse{Message: "user hasn't logged in yet"})
			return
		}
		token:=c.Request.Header.Get("Authorization")
		strArr := strings.Split(token, " ")
		_,permission :=jwt.GetPermission(strArr[1])
		if parameters != permission && permission != "public" {
			c.AbortWithStatusJSON(401, component.RestResponse{Message: "the user does not have access to access "})
			return
		}

		c.Next()
	}
}