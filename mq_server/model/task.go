package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Uid       uint   `gorm:"bot null"`
	Title     string `gorm:"index:not null"`
	Status    int    `gorm:"default:'0'"`
	Context   string `gorm:"type:longtext"`
	StartTime int64
	EnfTime   int64 `gorm:"default:'0'"`
}