package controllers

import (
	"EnglishApp/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type BMeStudentEntity struct {
	ID          int64  `json:"id"  gorm:"id"`
	ClassCode   string `json:"class_code"  gorm:"class_code"`
	TenFacebook string `json:"ten_facebook"  gorm:"ten_facebook"`
	HoVaTen     string `json:"ho_va_ten"  gorm:"ho_va_ten"`
	NgaySinh    string `json:"ngay_sinh"  gorm:"ngay_sinh"`
	SoDienThoai string `json:"so_dien_thoai"  gorm:"so_dien_thoai"`
	NgayVaoHoc  string `json:"ngay_vao_hoc"  gorm:"ngay_vao_hoc"`
	NgayKetThuc string `json:"ngay_ket_thuc"  gorm:"ngay_ket_thuc"`
	Note        string `json:"note"  gorm:"note"`
	Content     string `json:"content"  gorm:"content"`
}

func (BMeStudentEntity) TableName() string {
	return "ds_hv_tong"
}

func GetAllBmeStudents(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var result []BMeStudentEntity
		var response common.GTDResponse

		if err := db.Find(&result).Error; err != nil {
			gtdErr := common.GTDError{
				Code:    "400",
				Message: "test error",
			}
			response.Error = &gtdErr
			c.JSON(http.StatusBadGateway, response)
			return
		}

		response.Data = result

		c.JSON(http.StatusOK, response)
	}
}

// FindBmeStudentByKey is a handler function to find a specific BME class code by key
func FindBmeStudentByKey(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		filters := make(map[string]interface{})
		for key, values := range ctx.Request.URL.Query() {
			// Assume there is only one value for each key
			filters[key] = values[0]
		}

		var finalResult []BMeStudentEntity
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

// SearchBmeStudentByColumn is a handler function to find a specific BME class code by key
func SearchBmeStudentByColumn(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// Get the column and substring from the query parameters
		column := ctx.Query("column")
		substring := ctx.Query("substring")

		var finalResult []BMeStudentEntity
		result := db.Where(column+" LIKE ?", "%"+substring+"%").Find(&finalResult)
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
