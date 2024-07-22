package http

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type handler struct {
	db *sql.DB
}

func UserRoutes(router *gin.Engine, db *sql.DB) {
	h := handler{db}

	router.GET("/users", h.GetUsers)
}
