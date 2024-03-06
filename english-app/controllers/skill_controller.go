package controllers

import (
	"EnglishApp/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Skill struct {
	common.GTDEntity
	Name string `gorm:"type:varchar(255)" json:"name"`
	Code string `gorm:"type:varchar(255)" json:"code"`
}

func (Skill) TableName() string {
	return "skill"
}

// FindAllSkills retrieves all skills from the database.
func FindAllSkills(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var skills []Skill
		var response common.GTDResponse

		if err := db.Find(&skills).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		response.Data = skills
		c.JSON(http.StatusOK, &response)
	}
}

// CreateSkill creates a new skill in the database.
func CreateSkill(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var skill Skill
		var response common.GTDResponse

		if err := c.ShouldBindJSON(&skill); err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		if err := db.Create(&skill).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, gtdErr)
			return
		}

		response.Data = skill
		c.JSON(http.StatusCreated, &response)
	}
}

// UpdateSkill updates an existing skill in the database.
func UpdateSkill(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var skill Skill
		var response common.GTDResponse

		if err := db.First(&skill, id).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusNotFound),
				Message: "Skill not found",
			}
			c.JSON(http.StatusNotFound, gtdErr)
			return
		}

		if err := c.ShouldBindJSON(&skill); err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, gtdErr)
			return
		}

		if err := db.Save(&skill).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, gtdErr)
			return
		}

		response.Data = skill
		c.JSON(http.StatusOK, &response)
	}
}

// FindSkillByID retrieves a skill by ID from the database.
func FindSkillByID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var skill Skill
		var response common.GTDResponse

		if err := db.First(&skill, id).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusNotFound),
				Message: "Skill not found",
			}
			c.JSON(http.StatusNotFound, gtdErr)
			return
		}

		response.Data = skill
		c.JSON(http.StatusOK, &response)
	}
}

// DeleteSkill deletes a skill from the database.
func DeleteSkill(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var skill Skill

		if err := db.First(&skill, id).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusNotFound),
				Message: "Skill not found",
			}
			c.JSON(http.StatusNotFound, gtdErr)
			return
		}

		if err := db.Delete(&skill).Error; err != nil {
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
