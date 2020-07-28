package controller

import (
	"ftpLeaf/controller/downloadController"
	"github.com/gin-gonic/gin"
)

func Register(eng *gin.Engine) {
	downloadController.Register(eng)
}
