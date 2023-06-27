package models

type Channel struct {
  ID string `gorm:"primaryKey"`
  
  Name string
  Description string
  Price int64
  IsIndefinite bool
  IsEnabled bool

  MaxMembers int64
}

