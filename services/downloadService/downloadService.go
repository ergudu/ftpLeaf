package downloadService

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func GetFileName() string {
	return "测试文件.txt"
}

func GetFile(c *gin.Context) {
	localName := filepath.Join("filehouse", "leaf.txt")
	fi, err := os.Stat(localName)
	if err != nil {
		c.String(http.StatusOK, "没有找到文件")
		return
	}
	c.FileAttachment(localName, fi.Name())
}

func GetMaxFile(c *gin.Context) {
	var (
		err    error
		fh     *os.File
		nr, sz int64
	)
	localName := filepath.Join("filehouse", "protobuf-go-1.25.0.zip")
	if fh, err = os.Open(localName); err != nil {
		c.String(http.StatusOK, "没有找到文件")
	}
	fhinfo, err := fh.Stat()
	if err != nil {
		c.String(http.StatusOK, "没有找到文件")
	}
	for {
		if nr, err = io.Copy(c.Writer, fh); err != nil {
			break
		}
		if nr == 0 {
			time.Sleep(5 * time.Second)
		}

		if sz += nr; sz >= fhinfo.Size() {
			return
		}
	}
	c.Writer.Header().Set("Content-Type", "application/octet-stream")

	c.AbortWithStatus(http.StatusServiceUnavailable)
}
