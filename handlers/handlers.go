package handlers

import (
	"keymaster_go/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	headings = []string{"User", "Assigned Keys"}
	rows     = [][]string{{"Aaron Sum", "Key1, Key2"}, {"Mike Lloyd", "Key2, Key3"}}
	users    = []models.User{
		{ID: 1, Username: "mike", Email: "mike@mail.com", DisplayName: "Mike Lloyd", PasswordHash: "blah", CanLogin: true},
		{ID: 2, Username: "aaron", Email: "aaron@mail.com", DisplayName: "Aaron Sum", PasswordHash: "blah", CanLogin: false},
	}
	assignments = []models.Assignment{
		{ID: 1, User: "mike", Key: "key1", DateOut: time.Date(2021, 11, 10, 0, 0, 0, 0, time.UTC)},
		{ID: 2, User: "aaron", Key: "key2", DateOut: time.Date(2021, 11, 10, 0, 0, 0, 0, time.UTC)},
	}
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

func ShowKeys(c *gin.Context) {
	var keys []models.Key
	models.DB.Find(&keys)

	c.HTML(
		http.StatusOK,
		"keys.html",
		gin.H{
			"Keys": keys,
		},
	)
}

func ShowAddKey(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"add_key.html",
		gin.H{},
	)
}

func PostAddKey(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")

	models.DB.Create(&models.Key{
		Name:        name,
		Description: description,
	})

	c.Redirect(http.StatusFound, "/keys")
}

func ShowUsers(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"users.html",
		gin.H{
			"Users": users,
		},
	)
}

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
