package handler

import (
	"fmt"
	"github.com/JIeeiroSst/store/component"
	"github.com/JIeeiroSst/store/utils/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

func Login(c *gin.Context) {
	username, password := c.PostForm("username"), c.PostForm("password")
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
	token:=user.Login(username,password)
	if token=="" {
		c.JSON(200, component.RestResponse{Message: "no such account"})
		return
	}

	u, err := uuid.NewRandom()
	if err != nil {
		log.Println(fmt.Errorf("failed to generate UUID: %w", err))
	}
	sessionId := fmt.Sprintf("%s-%s", u.String(), username)
	err = component.GlobalCache.Set(sessionId, []byte(username))
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to store current subject in cache: %w", err))
		return
	}
	c.SetCookie("current_subject", sessionId, 30*60, "/api", "", false, true)
	c.JSON(200, component.RestResponse{Code: 1, Message: username +"Token:"+ token + " logged in successfully"})
}

func SingUp(c *gin.Context){
	username,password:=c.PostForm("username"),c.PostForm("password")
	info:= user.SignUp(username,password)
	c.JSON(200,component.RestResponse{Code:1,Message: info+"signup in successfully"})
}