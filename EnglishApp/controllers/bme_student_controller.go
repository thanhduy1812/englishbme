package controllers

import (
	"EnglishApp/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type BMeStudentEntity struct {
	ID          int64  `json:"id"  gorm:"id"`
	Status      string `json:"status"  gorm:"status"`
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
		//if err := c.ShouldBind(&data); err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{
		//		"error": common.GTDError{
		//			Code:    "400",
		//			Message: err.Error(),
		//		},
		//	})
		//	return
		//}

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
