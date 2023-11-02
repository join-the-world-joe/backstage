package global

import (
	"backstage/common/middleware"
	"github.com/gin-gonic/gin"
)

var _router *gin.Engine

func SetRouter(router *gin.Engine) {
	_router = router
	_router.Use(middleware.Cors())
	//_router.Use(middleware.Permission())
}

func Router() *gin.Engine {
	return _router
}
