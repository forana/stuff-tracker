package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Bind(e *gin.Engine) {
	for _, r := range []route{
		getRecentItems,
		ping,
	} {
		e.Handle(r.Verb, r.Path, r.Func)
		logrus.Infof("Bound %s %s", r.Verb, r.Path)
	}
}
