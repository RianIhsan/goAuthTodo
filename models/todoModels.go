package models

type Todo struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Complete bool   `json:"complete" gorm:"default:false;not null"`
	UserID   int    `json:"user_id" gorm:"not null"`
	User     User   `json:"user" gorm:"foreignkey:UserID"`
}

type TodoRes struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
	UserID   int    `json:"-"`
}

func (TodoRes) TableName() string {
	return "todos"
}
