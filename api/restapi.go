package api

import (
	"github.com/gin-gonic/gin"
)

type restapi struct {
	eng *gin.Engine
}

// InitRestAPI ...
func InitRestAPI() Client {
	gin.New()
}
