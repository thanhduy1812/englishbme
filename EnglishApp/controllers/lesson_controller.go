package controllers

import (
	"EnglishApp/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Lesson struct {
	common.GTDEntity
	SkillID      uint   // Assuming SkillID is a foreign key to another table
	ShortContent string `json:"shortContent" gorm:"short_content"`
	Content      string `gorm:"type:text"`
	HTMLContent  string `json:"htmlContent" gorm:"html_content"`
	Thumbnail    []byte `json:"thumbnail" gorm:"type:longblob"`
}

func (Lesson) TableName() string {
	return "lesson"
}

// FindAllLessons retrieves all lessons from the database.
func FindAllLessons(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var lessons []Lesson
		var response common.GTDResponse

		if err := db.Find(&lessons).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		response.Data = lessons
		c.JSON(http.StatusOK, &response)
	}
}

// CreateLesson creates a new lesson in the database.
func CreateLesson(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var lesson Lesson
		var response common.GTDResponse

		if err := c.ShouldBindJSON(&lesson); err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		if err := db.Create(&lesson).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, gtdErr)
			return
		}

		response.Data = lesson
		c.JSON(http.StatusCreated, &response)
	}
}

// UpdateLesson updates an existing lesson in the database.
func UpdateLesson(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var lesson Lesson
		var response common.GTDResponse

		if err := db.First(&lesson, id).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusNotFound),
				Message: "Lesson not found",
			}
			c.JSON(http.StatusNotFound, gtdErr)
			return
		}

		if err := c.ShouldBindJSON(&lesson); err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		if err := db.Save(&lesson).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, gtdErr)
			return
		}

		response.Data = lesson
		c.JSON(http.StatusOK, &response)
	}
}

// FindLessonByID retrieves a lesson by ID from the database.
func FindLessonByID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var lesson Lesson
		var response common.GTDResponse

		if err := db.First(&lesson, id).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusNotFound),
				Message: "Lesson not found",
			}
			c.JSON(http.StatusNotFound, gtdErr)
			return
		}

		response.Data = lesson
		c.JSON(http.StatusOK, &response)
	}
}

// DeleteLesson deletes a lesson from the database.
func DeleteLesson(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var lesson Lesson

		if err := db.First(&lesson, id).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusNotFound),
				Message: "Lesson not found",
			}
			c.JSON(http.StatusNotFound, gtdErr)
			return
		}

		if err := db.Delete(&lesson).Error; err != nil {
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
