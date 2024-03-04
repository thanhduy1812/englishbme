package controllers

import (
	"EnglishApp/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Feedback struct {
	common.GTDEntity
	Question       string `gorm:"column:fb_question" json:"fb_question"`
	DefaultAnswers string `gorm:"column:fb_default_answers" json:"fb_default_answers"`
}

func (Feedback) TableName() string {
	return "feed_back"
}

// FindAllFeedbacks retrieves all feedbacks from the database.
func FindAllFeedbacks(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var feedbacks []Feedback
		var response common.GTDResponse

		if err := db.Find(&feedbacks).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		response.Data = feedbacks
		c.JSON(http.StatusOK, &response)
	}
}

// CreateFeedback creates a new feedback in the database.
func CreateFeedback(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var feedback Feedback
		var response common.GTDResponse

		if err := c.ShouldBindJSON(&feedback); err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		if err := db.Create(&feedback).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, gtdErr)
			return
		}

		response.Data = feedback
		c.JSON(http.StatusCreated, &response)
	}
}

// UpdateFeedback updates an existing feedback in the database.
func UpdateFeedback(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var feedback Feedback
		var response common.GTDResponse

		if err := db.First(&feedback, id).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusNotFound),
				Message: "Feedback not found",
			}
			c.JSON(http.StatusNotFound, gtdErr)
			return
		}

		if err := c.ShouldBindJSON(&feedback); err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		if err := db.Save(&feedback).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, gtdErr)
			return
		}

		response.Data = feedback
		c.JSON(http.StatusOK, &response)
	}
}

// FindFeedbackByID retrieves a feedback by ID from the database.
func FindFeedbackByID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var feedback Feedback
		var response common.GTDResponse

		if err := db.First(&feedback, id).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusNotFound),
				Message: "Feedback not found",
			}
			c.JSON(http.StatusNotFound, gtdErr)
			return
		}

		response.Data = feedback
		c.JSON(http.StatusOK, &response)
	}
}

// DeleteFeedback deletes a feedback from the database.
func DeleteFeedback(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var feedback Feedback

		if err := db.First(&feedback, id).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusNotFound),
				Message: "Feedback not found",
			}
			c.JSON(http.StatusNotFound, gtdErr)
			return
		}

		if err := db.Delete(&feedback).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, gtdErr)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
