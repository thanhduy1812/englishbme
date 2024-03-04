package controllers

import (
	"EnglishApp/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type LessonRoadmap struct {
	common.GTDEntity
	ClassRoadmapID uint   // Assuming ClassRoadmapID is a foreign key to another table
	SkillID        uint   // Assuming SkillID is a foreign key to another table
	LessonID       uint   // Assuming LessonID is a foreign key to another table
	LessonStatus   string `gorm:"type:varchar(255)"`
	StartDate      string `gorm:"type:varchar(255)"`
	EndDate        string `gorm:"type:varchar(255)"`
}

func (LessonRoadmap) TableName() string {
	return "lesson_roadmap"
}

// FindAllLessonRoadmaps retrieves all lesson roadmaps from the database.
func FindAllLessonRoadmaps(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var lessonRoadmaps []LessonRoadmap
		var response common.GTDResponse

		if err := db.Find(&lessonRoadmaps).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		response.Data = lessonRoadmaps
		c.JSON(http.StatusOK, &response)
	}
}

// CreateLessonRoadmap creates a new lesson roadmap in the database.
func CreateLessonRoadmap(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var lessonRoadmap LessonRoadmap
		var response common.GTDResponse

		if err := c.ShouldBindJSON(&lessonRoadmap); err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		if err := db.Create(&lessonRoadmap).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, gtdErr)
			return
		}

		response.Data = lessonRoadmap
		c.JSON(http.StatusCreated, &response)
	}
}

// UpdateLessonRoadmap updates an existing lesson roadmap in the database.
func UpdateLessonRoadmap(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var lessonRoadmap LessonRoadmap
		var response common.GTDResponse

		if err := db.First(&lessonRoadmap, id).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusNotFound),
				Message: "Lesson roadmap not found",
			}
			c.JSON(http.StatusNotFound, gtdErr)
			return
		}

		if err := c.ShouldBindJSON(&lessonRoadmap); err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		if err := db.Save(&lessonRoadmap).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, gtdErr)
			return
		}

		response.Data = lessonRoadmap
		c.JSON(http.StatusOK, &response)
	}
}

// FindLessonRoadmapByID retrieves a lesson roadmap by ID from the database.
func FindLessonRoadmapByID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var lessonRoadmap LessonRoadmap
		var response common.GTDResponse

		if err := db.First(&lessonRoadmap, id).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusNotFound),
				Message: "Lesson roadmap not found",
			}
			c.JSON(http.StatusNotFound, gtdErr)
			return
		}

		response.Data = lessonRoadmap
		c.JSON(http.StatusOK, &response)
	}
}
