package handlers

import (
	"keymaster_go/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	assignments = []models.Assignment{
		{ID: 1, User: "mike", Key: "key1", DateOut: time.Date(2021, 11, 10, 0, 0, 0, 0, time.UTC)},
		{ID: 2, User: "aaron", Key: "key2", DateOut: time.Date(2021, 11, 10, 0, 0, 0, 0, time.UTC)},
	}
)

func ShowAssignments(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"assignments.html",
		gin.H{
			"Assignments": assignments,
		},
	)
}

func ShowAssignKey(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"assign_key.html",
		gin.H{
			"Assignments": assignments,
		},
	)
}
