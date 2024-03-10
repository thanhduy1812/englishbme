package controllers

import (
	"EnglishApp/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type BmeCourse struct {
	ID              uint   `gorm:"primaryKey"`
	OwnerID         string `json:"owner_id" gorm:"column:owner_id"`
	MaLop           string `json:"ma_lop" gorm:"column:ma_lop"`
	ShvDanhSach     string `json:"shv_danh_sach" gorm:"column:shv_danh_sach"`
	ShvToiDa        string `json:"shv_toi_da" gorm:"column:shv_toi_da"`
	ShvBoSung       string `json:"shv_bo_sung" gorm:"column:shv_bo_sung"`
	TongHVDangHoc   string `json:"tong_hv_dang_hoc" gorm:"column:tong_hv_dang_hoc"`
	GiaoVienHienTai string `json:"giao_vien_hien_tai" gorm:"column:giao_vien_hien_tai"`
	NgayKhaiGiang   string `json:"ngay_khai_giang" gorm:"column:ngay_khai_giang"`
	DinhHuong       string `json:"dinh_huong" gorm:"column:dinh_huong"`
	PhatAm          string `json:"phat_am" gorm:"column:phat_am"`
	NguPhap         string `json:"ngu_phap" gorm:"column:ngu_phap"`
	Nghe            string `json:"nghe" gorm:"column:nghe"`
	Noi             string `json:"noi" gorm:"column:noi"`
	NkhtGV          string `json:"nkht_gv" gorm:"column:nkht_gv"`
	NkhtHV          string `json:"nkht_hv" gorm:"column:nkht_hv"`
	Mau             string `json:"mau" gorm:"column:mau"`
	Note            string `json:"note" gorm:"column:note"`
	Content         string `json:"content" gorm:"column:content"`
}

func (BmeCourse) TableName() string {
	return "ds_ma_lop"
}

// GetAllBmeCourse is a handler function to get all BME students from the database
func GetAllBmeCourse(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var students []BmeCourse
		result := db.Find(&students)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: result.Error.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, students)
	}
}

// GetBmeCourse is a handler function to get a specific BME class code by ID
func GetBmeCourse(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var classCode BmeCourse
		result := db.First(&classCode, id)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: result.Error.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, classCode)
	}
}

// FindBmeCourseByKey is a handler function to find a specific BME class code by key
func FindBmeCourseByKey(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		filters := make(map[string]interface{})
		for key, values := range ctx.Request.URL.Query() {
			// Assume there is only one value for each key
			filters[key] = values[0]
		}

		var classCodes []BmeCourse
		result := db.Where(filters).Find(&classCodes)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: result.Error.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, classCodes)
	}
}

// CreateBmeCourse is a handler function to create a new BME class code
func CreateBmeCourse(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var classCode BmeCourse
		if err := ctx.ShouldBindJSON(&classCode); err != nil {
			ctx.JSON(http.StatusInternalServerError, common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: err.Error(),
			})
			return
		}

		result := db.Create(&classCode)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: result.Error.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, classCode)
	}
}

// UpdateBmeCourse is a handler function to update an existing BME class code
func UpdateBmeCourse(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var classCode BmeCourse
		result := db.First(&classCode, id)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: result.Error.Error(),
			})
			return
		}

		if err := ctx.ShouldBindJSON(&classCode); err != nil {
			ctx.JSON(http.StatusInternalServerError, common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: result.Error.Error(),
			})
			return
		}

		result = db.Save(&classCode)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: result.Error.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, classCode)
	}
}

// DeleteBmeCourse is a handler function to delete a specific BME class code by ID
func DeleteBmeCourse(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var classCode BmeCourse
		result := db.Delete(&classCode, id)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, common.GTDError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: result.Error.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "BME class code deleted successfully"})
	}
}
