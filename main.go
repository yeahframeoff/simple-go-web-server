package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

const DefaultDb = "db.sqlite3"
const DefaultPort = 8080

func main() {
	setupLoggers()
	var (
		dbName string
		host   string
		port   int
	)
	flag.StringVar(&dbName, "db", "", "path to sqlite database")
	flag.StringVar(&host, "host", "localhost", "host to listen")
	flag.IntVar(&port, "port", DefaultPort, "host to listen")
	flag.Parse()
	if dbName == "" {
		dbName = DefaultDb
		infoLog.Warnf("Database not specified. Using `%s` to store data", dbName)
	} else {
		infoLog.Infof("Using `%s` to store data", dbName)
	}
	db := setupDb(dbName)

	app := NewApp(db)
	defer closeDb(db)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/albums", app.getAlbums)
	router.GET("/albums/:id", app.getAlbumById)
	router.POST("/albums", app.postAlbum)
	router.GET("/health", app.healthCheck)

	address := fmt.Sprintf("%s:%d", host, port)
	infoLog.Infof("Running app on %s", address)
	err := router.Run(address)
	if err != nil {
		errorLog.Fatal(err)
	}
}
