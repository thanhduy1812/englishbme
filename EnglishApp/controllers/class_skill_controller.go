package controllers

import (
	"EnglishApp/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type ClassSkill struct {
	common.GTDEntity
	ClassCodeID *uint     `json:"classCodeId" gorm:"class_code_id"`
	SkillID     *uint     `json:"skillId" gorm:"skill_id"`
	StartDate   time.Time `json:"startDate" gorm:"start_date"`
	EndDate     time.Time `json:"endDate" gorm:"end_date"`
}

func (ClassSkill) TableName() string {
	return "class_skill"
}

// FindAllClassSkills retrieves all class skills from the database.
func FindAllClassSkills(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var classSkills []ClassSkill
		var response common.GTDResponse

		if err := db.Find(&classSkills).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		response.Data = classSkills
		c.JSON(http.StatusOK, &response)
	}
}

// CreateClassSkill creates a new class skill in the database.
func CreateClassSkill(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var classSkill ClassSkill
		var response common.GTDResponse

		if err := c.ShouldBindJSON(&classSkill); err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		if err := db.Create(&classSkill).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, gtdErr)
			return
		}

		response.Data = classSkill
		c.JSON(http.StatusCreated, &response)
	}
}

// UpdateClassSkill updates an existing class skill in the database.
func UpdateClassSkill(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var classSkill ClassSkill
		var response common.GTDResponse

		if err := db.First(&classSkill, id).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusNotFound),
				Message: "Class skill not found",
			}
			c.JSON(http.StatusNotFound, gtdErr)
			return
		}

		if err := c.ShouldBindJSON(&classSkill); err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		if err := db.Save(&classSkill).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, gtdErr)
			return
		}

		response.Data = classSkill
		c.JSON(http.StatusOK, &response)
	}
}

// FindClassSkillByID retrieves a class skill by ID from the database.
func FindClassSkillByID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var classSkill ClassSkill
		var response common.GTDResponse

		if err := db.First(&classSkill, id).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusNotFound),
				Message: "Class skill not found",
			}
			c.JSON(http.StatusNotFound, gtdErr)
			return
		}

		response.Data = classSkill
		c.JSON(http.StatusOK, &response)
	}
}

// DeleteClassSkill deletes a class skill from the database.
func DeleteClassSkill(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var classSkill ClassSkill
		var response common.GTDResponse

		if err := db.First(&classSkill, id).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusNotFound),
				Message: "Class skill not found",
			}
			c.JSON(http.StatusNotFound, gtdErr)
			return
		}

		if err := db.Delete(&classSkill).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, gtdErr)
			return
		}

		response.Data = "Class skill deleted successfully"
		c.JSON(http.StatusOK, &response)
	}
}
