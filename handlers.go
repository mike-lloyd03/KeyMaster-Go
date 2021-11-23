package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	headings = []string{"User", "Assigned Keys"}
	rows     = [][]string{{"Aaron Sum", "Key1, Key2"}, {"Mike Lloyd", "Key2, Key3"}}
	keys     = []Key{
		{"key1", "Opens things", "Active"},
		{"key2", "Opens stuff", "Active"},
	}
	users = []User{
		{1, "mike", "mike@mail.com", "Mike Lloyd", "blah", true},
		{2, "aaron", "aaron@mail.com", "Aaron Sum", "blah", false},
	}
)

func showIndex(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"Headings": headings,
			"Rows":     rows,
		},
	)
}

func showKeys(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"keys.html",
		gin.H{
			"Keys": keys,
		},
	)
}

func showUsers(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"users.html",
		gin.H{
			"Users": users,
		},
	)
}
