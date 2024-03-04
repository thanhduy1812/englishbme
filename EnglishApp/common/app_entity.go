package common

import "time"

type GTDEntity struct {
	Id        int        `json:"id" gorm:"primary_key;auto_increment"`
	CreatedAt *time.Time `json:"created_at" gorm:"default:current_timestamp()"`
	UpdateAt  *time.Time `json:"update_at,omitempty" gorm:"default:current_timestamp()"`
	CreatedBy *string    `json:"created_by"`
	UpdatedBy *string    `json:"updated_by"`
	IsDeleted bool       `json:"is_deleted"`
}
