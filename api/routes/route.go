package routes

import (
	"github.com/gin-gonic/gin"
)

type route struct {
	Verb string
	Path string
	Func func(*gin.Context)
}
