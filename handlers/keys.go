package handlers

import (
	"keymaster_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetKeys(c *gin.Context) {
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

func GetAddKey(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"add_key.html",
		gin.H{},
	)
}

func PostAddKey(c *gin.Context) {
	if c.PostForm("submit") != "" {
		name := c.PostForm("name")
		description := c.PostForm("description")

		models.DB.Create(&models.Key{
			Name:        name,
			Description: description,
		})
	}

	c.Redirect(http.StatusFound, "/keys")
}

func GetEditKey(c *gin.Context) {
	var key models.Key
	models.DB.First(&key, "name = ?", c.Query("name"))

	c.HTML(
		http.StatusOK,
		"edit_key.html",
		gin.H{
			"Name":        key.Name,
			"Description": key.Description,
			"Status":      key.Status,
		},
	)
}

func PostEditKey(c *gin.Context) {
	var key models.Key
	models.DB.First(&key, "name = ?", c.Query("name"))

	if c.PostForm("submit") != "" {
		description := c.PostForm("description")
		status := c.PostForm("status")

		models.DB.Model(&key).Updates(models.Key{
			Description: description,
			Status:      status,
		})
	} else if c.PostForm("delete") != "" {
		models.DB.Delete(&key, "name = ?", key.Name)
	}

	c.Redirect(http.StatusFound, "/keys")
}
