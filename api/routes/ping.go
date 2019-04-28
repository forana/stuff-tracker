package routes

import (
	"github.com/gin-gonic/gin"
)

var ping = route{
	Verb: "GET",
	Path: "/ping",
	Func: func(c *gin.Context) {
		c.String(200, "pong")
	},
}
