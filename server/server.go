package server

import (
	"ThirdEssentials/services"

	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {

	r := gin.Default()
	services.Router(r)

	return r
}
