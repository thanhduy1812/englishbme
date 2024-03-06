package controllers

import (
	"EnglishApp/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ClassCode struct {
	common.GTDEntity
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Code        string `gorm:"type:varchar(255)" json:"code"`
	RoomID      *uint  `json:"roomId" gorm:"room_id"`
	Capacity    int    `json:"capacity"`
	Type        string `gorm:"type:varchar(255)" json:"type"`
	Status      string `gorm:"type:varchar(255)" json:"status"`
	TotalLesson *uint  `json:"totalLesson" gorm:"total_lesson"`
	Color       string `gorm:"type:varchar(255)" json:"color"`
	CheckerID   *uint  `json:"checkerId" gorm:"checker_id"`
	AssistantID *uint  `json:"assistantId" gorm:"assistant_id"`
	TeacherID   *uint  `json:"teacherId" gorm:"teacher_id"`
}

func (ClassCode) TableName() string {
	return "class_code"
}

// FindAllClassCodes retrieves all class codes from the database.
func FindAllClassCodes(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var classCodes []ClassCode
		var response common.GTDResponse

		if err := db.Find(&classCodes).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		response.Data = classCodes
		c.JSON(http.StatusOK, &response)
	}
}

// CreateClassCode creates a new class code in the database.
func CreateClassCode(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var classCode ClassCode
		var response common.GTDResponse

		if err := c.ShouldBindJSON(&classCode); err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		if err := db.Create(&classCode).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, gtdErr)
			return
		}

		response.Data = classCode
		c.JSON(http.StatusCreated, &response)
	}
}

// UpdateClassCode updates an existing class code in the database.
func UpdateClassCode(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var classCode ClassCode
		var response common.GTDResponse

		if err := db.First(&classCode, id).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusNotFound),
				Message: "Class code not found",
			}
			c.JSON(http.StatusNotFound, gtdErr)
			return
		}

		if err := c.ShouldBindJSON(&classCode); err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		if err := db.Save(&classCode).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, gtdErr)
			return
		}

		response.Data = classCode
		c.JSON(http.StatusOK, &response)
	}
}

// FindClassCodeByID retrieves a class code by ID from the database.
func FindClassCodeByID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var classCode ClassCode
		var response common.GTDResponse

		if err := db.First(&classCode, id).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusNotFound),
				Message: "Class code not found",
			}
			c.JSON(http.StatusNotFound, gtdErr)
			return
		}

		response.Data = classCode
		c.JSON(http.StatusOK, &response)
	}
}

// DeleteClassCode deletes a class code from the database.
func DeleteClassCode(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var classCode ClassCode
		var response common.GTDResponse

		if err := db.First(&classCode, id).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusNotFound),
				Message: "Class code not found",
			}
			c.JSON(http.StatusNotFound, gtdErr)
			return
		}

		if err := db.Delete(&classCode).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, gtdErr)
			return
		}

		response.Data = classCode
		c.JSON(http.StatusOK, &response)
	}
}
