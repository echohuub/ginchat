package models

import "gorm.io/gorm"

type GroupBasic struct {
	gorm.Model
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
