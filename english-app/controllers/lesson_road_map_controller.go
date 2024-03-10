package controllers

import (
	"EnglishApp/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type LessonRoadmap struct {
	common.GTDEntity
	ClassCode    string     `json:"classCode" gorm:"class_code"`
	SkillCode    string     `json:"skillCode" gorm:"skill_code"`
	LessonName   string     `json:"lessonName" gorm:"lesson_name"`
	LessonStatus *string    `json:"lessonStatus" gorm:"lesson_status"`
	StartDate    *time.Time `json:"startDate" gorm:"start_date"`
	EndDate      *time.Time `json:"endDate" gorm:"end_date"`
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

// FindLessonRoadMapByKey retrieves a lesson roadmap by ID from the database.
func FindLessonRoadMapByKey(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		filters := make(map[string]interface{})
		for key, values := range ctx.Request.URL.Query() {
			// Assume there is only one value for each key
			filters[key] = values[0]
		}

		var finalResult []LessonRoadmap
		result := db.Where(filters).Find(&finalResult)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: result.Error.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, finalResult)
	}
}
