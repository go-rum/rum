package rum

import (
	"github.com/gin-gonic/gin"
	"github.com/go-rum/rum/errno"
	"net/http"
)

func WrongWithErrno(ctx *gin.Context, err error) {
	code, msg := errno.Decode(err)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": msg,
	})
}

func WrongWithMessage(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}
