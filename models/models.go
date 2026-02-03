package models

import "gorm.io/gorm"

type Task struct {
    gorm.Model
    Description string `json:"description" gorm:"type:text"`
    Note        string `json:"note" gorm:"type:text"`
}
