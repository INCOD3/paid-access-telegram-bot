package models

import "gorm.io/gorm"

type Channel struct {
  gorm.Model

  Id string `gorm:"unique;"`
  
  Name string
  Description string
  Price int64
  IsIndefinite bool
  IsEnabled bool

  MaxMembers int64
}

func NewChannel(id, name, description string, price int64, isIndefinite bool, maxMembers int64) *Channel {
  return &Channel{
    Id: id,
    Name: name,
    Description: description,
    Price: price,
    IsIndefinite: isIndefinite,
    IsEnabled: true,
    MaxMembers: maxMembers,
  }
}
