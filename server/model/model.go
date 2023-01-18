package model

import "gorm.io/gorm"

type BaseModel struct {
	gorm.Model
}

func (a BaseModel) GetStringID() string {
	return string(a.ID)
}
