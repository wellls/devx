package http

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetUsers(c *gin.Context) {
	selectAll(h.db)
	c.JSON(http.StatusOK, gin.H{"message": "List of users"})
}

func selectAll(db *sql.DB) error {
	stmt := `select * from users`
	rows, err := db.Query(stmt)
	if err != nil {
		return err
	}

	log.Println(rows)

	return nil
}
