package restapi

import (
	"github.com/gin-gonic/gin"
)

func responseFailed(ctx *gin.Context, err error) {
	ctx.JSON(404, gin.H{
		"CODE":    404,
		"MESSAGE": err.Error(),
	})
}

func responseSuccess(ctx *gin.Context, v interface{}) {
	ctx.JSON(200, gin.H{
		"CODE":    200,
		"MESSAGE": "SUCCESS",
		"DATA":    v,
	})
}

func formatterResponse(ctx *gin.Context, err error, v interface{}) {
	if err != nil {
		responseFailed(ctx, err)
		return
	}
	responseSuccess(ctx, v)
}
