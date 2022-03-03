package rum

import (
	"github.com/gin-gonic/gin"
	"github.com/go-rum/rum/errno"
	"net/http"
)

func Wrong(ctx *gin.Context, err *errno.Errno) {
	code, msg := errno.Decode(err)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": msg,
	})
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}
