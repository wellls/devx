package http

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetUsers(c *gin.Context) {

	err := selectAll(h.db)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "erro"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "List of users"})
}

func selectAll(db *sql.DB) error {
	stmt := `select * from users`
	rows, err := db.Query(stmt)
	if err != nil {
		return err
	}

	fmt.Println(rows)

	return nil
}
