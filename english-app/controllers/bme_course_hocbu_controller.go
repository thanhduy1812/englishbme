package controllers

import (
	"EnglishApp/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type BmeCourseHocBu struct {
	ID            int64  `json:"id"  gorm:"id"`
	NgayHocBu     string `json:"ngay_hoc_bu"  gorm:"ngay_hoc_bu"`
	MaLop         string `json:"ma_lop"  gorm:"ma_lop"`
	TenFacebook   string `json:"ten_facebook"  gorm:"ten_facebook"`
	HoVaTen       string `json:"ho_va_ten"  gorm:"ho_va_ten"`
	NgaySinh      string `json:"ngay_sinh"  gorm:"ngay_sinh"`
	SoDienThoai   string `json:"so_dien_thoai"  gorm:"so_dien_thoai"`
	GiaoVienDayBu string `json:"giao_vien_day_bu"  gorm:"giao_vien_day_bu"`
	LopHocBu      string `json:"lop_hoc_bu"  gorm:"lop_hoc_bu"`
	PhiHocBu      string `json:"phi_hoc_bu"  gorm:"phi_hoc_bu"`
	BaiHocBu      string `json:"bai_hoc_bu"  gorm:"bai_hoc_bu"`
	Note          string `json:"note"  gorm:"note"`
}

func (BmeCourseHocBu) TableName() string {
	return "ds_hoc_bu"
}

// GetAllBmeCourseHocBu is a handler function to get all BME students from the database
func GetAllBmeCourseHocBu(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var courses []BmeCourseHocBu
		result := db.Find(&courses)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: result.Error.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, courses)
	}
}

// FindBmeCourseHocBuByKey is a handler function to find a specific BME class code by key
func FindBmeCourseHocBuByKey(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		filters := make(map[string]interface{})
		for key, values := range ctx.Request.URL.Query() {
			// Assume there is only one value for each key
			filters[key] = values[0]
		}

		var finalResult []BmeCourseHocBu
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
