package main

import (
	"os"

	"github.com/forana/stuff-tracker/api/db"
	"github.com/forana/stuff-tracker/api/routes"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

func main() {
	setupLogger()
	migrate()
	serve()
}

func setupLogger() {
	logrus.SetLevel(logrus.DebugLevel)
	// gin's debug mode is utterly worthless
	gin.SetMode(gin.ReleaseMode)
}

func migrate() {
	err := db.Migrate()
	if err != nil {
		panic(err)
	}
}

func serve() {
	r := gin.New()
	r.Use(ginlogrus.Logger(logrus.StandardLogger()), gin.Recovery())
	r.Use(static.Serve("/", static.LocalFile("./static", true)))
	logrus.Info("Starting server")
	routes.Bind(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logrus.Info("Listening on port " + port)
	r.Run(":" + port) // listen and serve on 0.0.0.0:${PORT}
}
