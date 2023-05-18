package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	Course struct {
		ID          uint           `gorm:"primaryKey" json:"id"`
		Title       string         `json:"title"`
		Description string         `json:"desc"`
		Price       float64        `json:"price"`
		Banner      string         `json:"banner"`
		CategoryID  uint           `json:"category_id"`
		Enroll      uint           `json:"enroll"`
		CreatedAt   time.Time      `gorm:"autoCreateTime" json:"createdAt"`
		UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
		DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
		Category    *Category      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category,omitempty"`
	}

	Category struct {
		ID        uint      `gorm:"primaryKey" json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
		UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	}
)
