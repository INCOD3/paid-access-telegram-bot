package models

import (
	"time"
)

type Subscription struct {
  ID uint `gorm:"primaryKey"`

  IsIndefinite bool
  ExpiresAt int64
  ExpiresAtParsed time.Time `gorm:"-"`
  
  UserId int64
  User User `gorm:"foreignKey:UserId"`
}
