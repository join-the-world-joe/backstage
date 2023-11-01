package api

import (
	"backstage/global"
	"github.com/gin-gonic/gin"
)

const (
	Input = "Input"
)

func echo() {
	global.Router().POST(Echo, func(c *gin.Context) {
		c.JSON(200, c.PostForm(Input))
	})
}
