package model

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Title    string `json:"title"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Banner   string `json:"banner"`
}
