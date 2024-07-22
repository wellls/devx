package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	userHttp "github.com/wellls/devx/internal/api/users/http"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	userHttp.UserRoutes(router, db)
}
