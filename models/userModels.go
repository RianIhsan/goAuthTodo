package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}

type UserRes struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (UserRes) TableName() string {
	return "users"
}
