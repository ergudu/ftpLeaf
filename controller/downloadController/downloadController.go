package downloadController

import (
	"ftpLeaf/services/downloadService"
	"github.com/gin-gonic/gin"
)

func Register(r gin.IRouter) {
	r.GET("/download", func(c *gin.Context) {
		downloadService.GetFile(c)
	})

	r.GET("/max", func(c *gin.Context) {
		downloadService.GetMaxFile(c)
	})
}
