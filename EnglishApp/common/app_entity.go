package common

import "time"

type GTDEntity struct {
	Id        int        `json:"id" gorm:"primary_key;auto_increment"`
	CreatedAt *time.Time `json:"createdAt" gorm:"created_at"`
	UpdateAt  *time.Time `json:"updateAt,omitempty" gorm:"update_at"`
	CreatedBy *string    `json:"created_by"`
	UpdatedBy *string    `json:"updated_by"`
	IsDeleted bool       `json:"is_deleted"`
}
