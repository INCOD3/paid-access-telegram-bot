package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
  gorm.Model

  ID int `gorm:"primaryKey"`

  IsIndefinite bool
  ExpiresAt int64
  ExpiresAtParsed time.Time `gorm:"-"`
  
  UserId int64
  User User `gorm:"foreignKey:UserId"`

  ChannelId string 
  Channel Channel `gorm:"foreignKey:ChannelId"`
}
