package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	headings = []string{"User", "Assigned Keys"}
	rows     = [][]string{{"Aaron Sum", "Key1, Key2"}, {"Mike Lloyd", "Key2, Key3"}}
)

func ShowIndex(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"Headings": headings,
			"Rows":     rows,
		},
	)
}
