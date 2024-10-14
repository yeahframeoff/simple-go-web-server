package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type App struct {
	AlbumService *AlbumService
	DB           *sql.DB
}

func NewApp(db *sql.DB) App {
	return App{DB: db, AlbumService: &AlbumService{DB: db}}
}

func (app *App) getAlbums(ctx *gin.Context) {
	var pagination Pagination

	hasLimit := ctx.Query("limit") != ""
	hasOffset := ctx.Query("offset") != ""

	if hasLimit && hasOffset {
		if ctx.BindQuery(&pagination) != nil {
			ctx.IndentedJSON(http.StatusBadRequest, "pagination could not parse")
			return
		}

		albums, err := app.AlbumService.fetchAlbums(&pagination)

		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
		} else {
			ctx.IndentedJSON(http.StatusOK, albums)
		}

	} else if hasLimit && !hasOffset {
		ctx.String(http.StatusBadRequest, "should also provide `offset` when providing `limit`")
	} else if !hasLimit && hasOffset {
		ctx.String(http.StatusBadRequest, "should also provide `limit` when providing `offset`")
	} else {

		albums, err := app.AlbumService.fetchAlbums(nil)

		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
		} else {
			ctx.IndentedJSON(http.StatusOK, albums)
		}
	}

}

func (app *App) postAlbum(ctx *gin.Context) {
	var body CreateAlbumBody

	if err := ctx.BindJSON(&body); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, errors.Join(errors.New("could not parse body"), err))
	}

	alb, err := app.AlbumService.createAlbum(body)

	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		if alb == nil {
			ctx.String(http.StatusCreated, "New Post succesfully created")
		} else {
			ctx.IndentedJSON(http.StatusCreated, alb)
		}
	}
}

func (app *App) getAlbumById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil || id <= 0 {
		ctx.String(http.StatusBadRequest, "bad album id")
	}

	album, err := app.AlbumService.fetchAlbumById(id)

	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else if album != nil {
		ctx.IndentedJSON(http.StatusOK, album)
	} else {
		ctx.String(http.StatusNotFound, fmt.Sprintf("No album with id:%d", id))
	}

}

func (app *App) healthCheck(ctx *gin.Context) {
	if app.DB.Ping() != nil {
		ctx.String(http.StatusServiceUnavailable, "down")
	} else {
		ctx.String(http.StatusOK, "up")
	}
}
