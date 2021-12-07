package handlers

import (
	"keymaster_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.HTML(
		http.StatusOK,
		"users.html",
		gin.H{
			"Users": users,
		},
	)
}

func GetAddUser(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"add_user.html",
		gin.H{},
	)
}

func PostAddUser(c *gin.Context) {
	if c.PostForm("submit") != "" {
		username := c.PostForm("username")
		email := c.PostForm("email")
		displayName := c.PostForm("display_name")
		password := c.PostForm("password")
		canLogin := c.PostForm("can_login") == "y"

		models.DB.Create(&models.User{
			Username:     username,
			Email:        email,
			DisplayName:  displayName,
			PasswordHash: password,
			CanLogin:     canLogin,
		})
	}

	c.Redirect(http.StatusFound, "/users")
}

func GetEditUser(c *gin.Context) {
	var user models.User
	err := models.DB.First(&user, c.Query("id")).Error
	if err != nil {
		c.HTML(
			http.StatusNotFound,
			"404.html",
			gin.H{},
		)
		return
	}

	c.HTML(
		http.StatusOK,
		"edit_user.html",
		gin.H{
			"ID":          user.ID,
			"Username":    user.Username,
			"Email":       user.Email,
			"DisplayName": user.DisplayName,
			"CanLogin":    user.CanLogin,
		},
	)
}

func PostEditUser(c *gin.Context) {
	var user models.User
	models.DB.First(&user, c.Query("id"))

	if c.PostForm("submit") != "" {
		username := c.PostForm("username")
		email := c.PostForm("email")
		displayName := c.PostForm("display_name")
		canLogin := c.PostForm("can_login") == "y"

		models.DB.Model(&user).Updates(models.User{
			Username:    username,
			Email:       email,
			DisplayName: displayName,
			CanLogin:    canLogin,
		})
	} else if c.PostForm("delete") != "" {
		models.DB.Delete(&user, user.ID)
	}

	c.Redirect(http.StatusFound, "/users")
}
