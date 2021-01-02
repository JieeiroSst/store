package handler

import (
	"github.com/JIeeiroSst/store/component"
	"github.com/gin-gonic/gin"
)

func ReadResource(c *gin.Context) {
	c.JSON(200, component.RestResponse{Code: 200, Message: "read resource successfully", Data: "resource"})
}

func WriteResource(c *gin.Context) {
	c.JSON(200, component.RestResponse{Code: 200, Message: "write resource successfully", Data: "resource"})
}