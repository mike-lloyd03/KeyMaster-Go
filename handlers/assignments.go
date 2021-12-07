package handlers

import (
	"errors"
	"fmt"
	"keymaster_go/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAssignments(c *gin.Context) {
	var assignments []models.Assignment
	models.DB.Find(&assignments)

	c.HTML(
		http.StatusOK,
		"assignments.html",
		gin.H{
			"Assignments": assignments,
		},
	)
}

func GetAssignKey(c *gin.Context) {
	var users []models.User
	var keys []models.Key

	models.DB.Find(&users)
	models.DB.Find(&keys)

	c.HTML(
		http.StatusOK,
		"assign_key.html",
		gin.H{
			"Users": users,
			"Keys":  keys,
		},
	)
}

func PostAssignKey(c *gin.Context) {
	if c.PostForm("submit") != "" {
		dateOut, err := time.Parse("2006-01-02", c.PostForm("date_out"))
		if err != nil {
			// There's a better way to do this
			panic("Invalid date format.")
		}

		for _, user := range c.PostFormArray("user") {
			for _, key := range c.PostFormArray("key") {
				// Check if key is currently assigned to user
				var existingAssignment models.Assignment
				err := models.DB.Where("user = ? AND key = ? AND date_in = ?", user, key, time.Time{}).First(&existingAssignment).Error

				if errors.Is(err, gorm.ErrRecordNotFound) {
					models.DB.Create(&models.Assignment{
						User:    user,
						Key:     key,
						DateOut: dateOut,
					})
				} else {
					// flash(
					// 	f'Key "{key}" already assigned to {get_display_name(user)}',
					// 	"danger",
					// )
				}
			}
			// flash(f'Key "{key}" assigned to {user}')
		}
	}

	c.Redirect(http.StatusFound, "/assignments")
}

func GetEditAssignment(c *gin.Context) {
	var users []models.User
	var keys []models.Key
	var assignment models.Assignment

	models.DB.Find(&users)
	models.DB.Find(&keys)
	err := models.DB.First(&assignment, c.Query("id")).Error

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
		"edit_assignment.html",
		gin.H{
			"Assignment": assignment,
			"Keys":       keys,
			"Users":      users,
		},
	)
}

func PostEditAssignment(c *gin.Context) {
	var assignment models.Assignment
	models.DB.First(&assignment, "id = ?", c.Query("id"))

	if c.PostForm("submit") != "" {
		user := c.PostForm("user")
		key := c.PostForm("key")
		dateOut, err := time.Parse("2006-01-02", c.PostForm("date_out"))
		if err != nil {
			fmt.Println("dateOut invalid date format.", c.PostForm("date_out"))
			panic(err)
		}

		var dateIn time.Time
		dateInString := c.PostForm("date_in")

		if dateInString == "" {
			dateIn = time.Time{}
		} else {
			dateIn, err = time.Parse("2006-01-02", dateInString)
			if err != nil {
				fmt.Println("dateIn invalid date format.", dateInString)
				panic(err)
			}
		}
		fmt.Println("dateIn:", dateIn)

		models.DB.Model(&assignment).Updates(map[string]interface{}{
			"User":    user,
			"Key":     key,
			"DateOut": dateOut,
			"DateIn":  dateIn,
		})
	} else if c.PostForm("delete") != "" {
		models.DB.Delete(&assignment, "id = ?", assignment.ID)
	}

	c.Redirect(http.StatusFound, "/assignments")
}
