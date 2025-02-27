package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"` // 'id' should be 'ID' (Go convention), and extra space in `gorm:"primaryKey index "` is removed
	Name     string `json:"name"` // Missing closing quote fixed
	Email    string `gorm:"unique" json:"email"` // Corrected misplaced gorm tag
	Password string `json:"password"` // Removed extra comma
}
