package controllers

import (
	"EnglishApp/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

// User entity struct
type User struct {
	common.GTDEntity
	FullName    string    `json:"fullName" gorm:"full_name"`
	SocialName  string    `json:"socialName" gorm:"social_name"`
	DOB         time.Time `json:"dob" gorm:"dob"`
	Avatar      string    `json:"avatar" gorm:"avatar"`
	Username    string    `json:"username" gorm:"username"`
	Password    string    `json:"password" gorm:"password"`
	Role        string    `json:"role" gorm:"role"`
	Email       string    `json:"email" gorm:"email"`
	PhoneNumber string    `json:"phoneNumber" gorm:"phone_number"`
	Address     string    `json:"address" gorm:"address"`
	Tag         string    `json:"tag" gorm:"tag"`
}

type UserCreation struct {
	FullName    string    `json:"fullName" gorm:"full_name"`
	SocialName  string    `json:"socialName" gorm:"social_name"`
	DOB         time.Time `json:"dob" gorm:"dob"`
	Username    string    `json:"username" gorm:"username"`
	Role        string    `json:"role" gorm:"role"` //admin | mentor | user
	Email       string    `json:"email" gorm:"email"`
	PhoneNumber string    `json:"phoneNumber" gorm:"phone_number"`
	Address     string    `json:"address" gorm:"address"`
}

func (User) TableName() string {
	return "user"
}
func (UserCreation) TableName() string {
	return "user"
}

// CreateUser creates a new user
func CreateUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var userCreation UserCreation
		var response common.GTDResponse

		// Bind the user data from the request body
		err := c.ShouldBindJSON(&userCreation)
		if err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, &gtdErr)
			return
		}

		user := User{
			FullName:    userCreation.FullName,
			SocialName:  userCreation.SocialName,
			DOB:         userCreation.DOB,
			Username:    userCreation.PhoneNumber,
			Password:    "123456",
			Role:        userCreation.Role,
			Email:       userCreation.Email,
			PhoneNumber: userCreation.PhoneNumber,
			Address:     userCreation.Address,
		}

		if err := db.Create(&user).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, &gtdErr)
			return
		}

		response.Data = user
		c.JSON(http.StatusOK, &response)
	}

}

// FindUserByID Find user by ID
func FindUserByID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Get the user ID from the URL parameter
		userID := c.Param("id")

		var user User
		var response common.GTDResponse

		if err := db.First(&user, userID).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, &gtdErr)
			return
		}

		response.Data = user
		c.JSON(200, &response)
	}
}

// FindAllUsers Find all users
func FindAllUsers(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var users []User
		var response common.GTDResponse
		if err := db.Find(&users).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, &gtdErr)
			return
		}

		response.Data = users
		c.JSON(200, &response)
	}
}

// UpdateUser Update user
func UpdateUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Get the user ID from the URL parameter
		userID := c.Param("id")
		var response common.GTDResponse

		var user User
		if err := db.First(&user, userID).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, &gtdErr)
			return
		}

		// Bind the updated data from the request body
		err := c.ShouldBindJSON(&user)
		if err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, &gtdErr)
			return
		}

		if err := db.Save(&user).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, &gtdErr)
			return
		}

		response.Data = user
		c.JSON(200, response)
	}
}

// DeleteUser Delete user
func DeleteUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Get the user ID from the URL parameter
		userID := c.Param("id")
		var response common.GTDResponse

		var user User
		if err := db.First(&user, userID).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, &gtdErr)
			return
		}

		if err := db.Delete(&user).Error; err != nil {
			gtdErr := &common.GTDError{
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, &gtdErr)
			return
		}

		response.Data = user
		c.JSON(200, response)
	}
}
