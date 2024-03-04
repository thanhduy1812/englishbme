package controllers

import (
	"EnglishApp/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type ClassUser struct {
	common.GTDEntity
	ClassCodeID uint
	UserID      uint
	StartDate   time.Time
	EndDate     time.Time
	CancelDate  time.Time
}

// FindAllClassUsers retrieves all class users from the database.
func FindAllClassUsers(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var classUsers []ClassUser
		var response gin.H

		if err := db.Find(&classUsers).Error; err != nil {
			response = gin.H{
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response = gin.H{
			"data": classUsers,
		}
		c.JSON(http.StatusOK, response)
	}
}

// CreateClassUser creates a new class user in the database.
func CreateClassUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var classUser ClassUser
		var response gin.H

		if err := c.ShouldBindJSON(&classUser); err != nil {
			response = gin.H{
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		if err := db.Create(&classUser).Error; err != nil {
			response = gin.H{
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response = gin.H{
			"data": classUser,
		}
		c.JSON(http.StatusOK, response)
	}
}

// UpdateClassUser updates an existing class user in the database.
func UpdateClassUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := gin.H{
				"error": "Invalid ID",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		var classUser ClassUser
		if err := db.First(&classUser, id).Error; err != nil {
			response := gin.H{
				"error": "Class user not found",
			}
			c.JSON(http.StatusNotFound, response)
			return
		}

		if err := c.ShouldBindJSON(&classUser); err != nil {
			response := gin.H{
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		if err := db.Save(&classUser).Error; err != nil {
			response := gin.H{
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := gin.H{
			"data": classUser,
		}
		c.JSON(http.StatusOK, response)
	}
}

// DeleteClassUser deletes a class user from the database.
func DeleteClassUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := gin.H{
				"error": "Invalid ID",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		var classUser ClassUser
		if err := db.First(&classUser, id).Error; err != nil {
			response := gin.H{
				"error": "Class user not found",
			}
			c.JSON(http.StatusNotFound, response)
			return
		}

		if err := db.Delete(&classUser).Error; err != nil {
			response := gin.H{
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := gin.H{
			"message": "Class user deleted successfully",
		}
		c.JSON(http.StatusOK, response)
	}
}
