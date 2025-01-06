package model

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID        uint           `gorm:"primaryKey"`
    TelegramID int64        `gorm:"unique;not null"`
    Username  string        `gorm:"unique;not null"`
    Name *string
    LastName *string
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}