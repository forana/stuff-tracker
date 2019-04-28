package routes

import (
	"strconv"

	"github.com/forana/stuff-tracker/api/db"
	"github.com/gin-gonic/gin"
)

var getRecentItems = route{
	Verb: "GET",
	Path: "/items/recent",
	Func: func(c *gin.Context) {
		countVal := c.Query("count")
		count := 10
		if countVal != "" {
			c, err := strconv.Atoi(countVal)
			if err == nil {
				count = c
			}
		}

		conn, err := db.Open()
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		c.JSON(200, map[string]int{"count": count})
	},
}
