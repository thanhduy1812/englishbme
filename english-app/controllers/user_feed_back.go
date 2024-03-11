package controllers

import (
	"EnglishApp/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type UserFeedback struct {
	common.GTDEntity
	UserName        string `json:"user_name" gorm:"user_name"`
	LessonRoadmapID uint
	CancelDate      time.Time
	FeedbackID      uint
	FeedbackAnswer  string
	FeedbackDate    time.Time `json:"feedback_date" gorm:"feedback_date"`
	CatchUpDate     time.Time `json:"catch_up_date" gorm:"catch_up_date"`
}

func (UserFeedback) TableName() string {
	return "user_feedback"
}

// FindAllUserFeedbacks retrieves all user feedbacks from the database.
func FindAllUserFeedbacks(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var userFeedbacks []UserFeedback
		var response gin.H

		if err := db.Find(&userFeedbacks).Error; err != nil {
			response = gin.H{
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response = gin.H{
			"data": userFeedbacks,
		}
		c.JSON(http.StatusOK, response)
	}
}

// CreateUserFeedback creates a new user feedback in the database.
func CreateUserFeedback(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var userFeedback UserFeedback
		var response gin.H

		if err := c.ShouldBindJSON(&userFeedback); err != nil {
			response = gin.H{
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		if err := db.Create(&userFeedback).Error; err != nil {
			response = gin.H{
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response = gin.H{
			"data": userFeedback,
		}
		c.JSON(http.StatusOK, response)
	}
}

// UpdateUserFeedback updates an existing user feedback in the database.
func UpdateUserFeedback(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := gin.H{
				"error": "Invalid ID",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		var userFeedback UserFeedback
		if err := db.First(&userFeedback, id).Error; err != nil {
			response := gin.H{
				"error": "User feedback not found",
			}
			c.JSON(http.StatusNotFound, response)
			return
		}

		if err := c.ShouldBindJSON(&userFeedback); err != nil {
			response := gin.H{
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		if err := db.Save(&userFeedback).Error; err != nil {
			response := gin.H{
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := gin.H{
			"data": userFeedback,
		}
		c.JSON(http.StatusOK, response)
	}
}

// DeleteUserFeedback deletes a user feedback from the database.
func DeleteUserFeedback(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := gin.H{
				"error": "Invalid ID",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		var userFeedback UserFeedback
		if err := db.First(&userFeedback, id).Error; err != nil {
			response := gin.H{
				"error": "User feedback not found",
			}
			c.JSON(http.StatusNotFound, response)
			return
		}

		if err := db.Delete(&userFeedback).Error; err != nil {
			response := gin.H{
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := gin.H{
			"message": "User feedback deleted successfully",
		}
		c.JSON(http.StatusOK, response)
	}
}
