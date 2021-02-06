package handler

import (
	"fmt"
	"github.com/JIeeiroSst/store/component"
	"github.com/JIeeiroSst/store/utils/jwt"
	"github.com/JIeeiroSst/store/utils/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"strings"
)

func Login(c *gin.Context) {
	username ,password := c.Query("username"), c.Query("password")
	for iter := component.GlobalCache.Iterator(); iter.SetNext(); {
		info, err := iter.Value()
		if err != nil {
			continue
		}
		if string(info.Value()) == username {
			component.GlobalCache.Delete(info.Key())
			log.Printf("forced %s to log out\n", username)
			break
		}
	}
	message,token,permission,err:=user.Login(username,password)
	if err != nil  {
		c.JSON(200, component.RestResponse{Code: 200,Message: message,Data: token,Permission:permission})
		return
	}
	u, err := uuid.NewRandom()
	if err != nil {
		log.Println(fmt.Errorf("failed to generate UUID: %w", err))
	}
	sessionId := fmt.Sprintf("%s-%s", u.String(), username)
	err = component.GlobalCache.Set(sessionId, []byte(username))
	if err != nil {
		c.JSON(200, component.RestResponse{Code: 200,Message: "failed to store current subject in cache"})
		return
	}
	c.SetCookie("current_subject", sessionId, 30*60, "/api", "", false, true)
	c.JSON(200, component.RestResponse{Code: 200, Message: message,Data: token,Permission:permission})
}

func SingUp(c *gin.Context){
	username ,password := c.Query("username"), c.Query("password")
	info:= user.SignUp(username,password)
	c.JSON(200,component.RestResponse{Code:200,Message: "signup in successfully",Data:info})
}

func Geranate(c *gin.Context) {
	token:=c.Request.Header.Get("Authorization")
	strArr := strings.Split(token, " ")
	_,permission :=jwt.GetPermission(strArr[0])
	c.JSON(200,component.RestResponse{Code:200,Message: "permission in successfully",Data:permission})
}