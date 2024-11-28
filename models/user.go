package models

type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Nickname  string `gorm:"not null"`
	Password  string `gorm:"not null"`
}
