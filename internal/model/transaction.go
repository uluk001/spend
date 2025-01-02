package model

import (
	"time"
	"gorm.io/gorm"
)


type TransactionCategory struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string `gorm:"not null"`
    Transactions []Transaction `gorm:"foreignKey:CategoryID"` // Явно указываем связь
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"` // Soft delete
}


type Transaction struct {
    ID         uint           `gorm:"primaryKey"`
    UserID     uint          `gorm:"not null"`
    Amount     int           `gorm:"not null"`
    Comment    string
    CategoryID uint          `gorm:"not null"` // Переместим выше для группировки
    Category   TransactionCategory `gorm:"foreignKey:CategoryID"` // Явно указываем связь
    CreatedAt  time.Time
    UpdatedAt  time.Time
    DeletedAt  gorm.DeletedAt `gorm:"index"`
}