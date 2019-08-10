package model

type User struct {
	ID uint `gorm:"primary_key"`
	//gorm.Model
	Username  string `gorm:"column:username"`
	Password  string
	Phone     string
	CreatedAt string `gorm:"column:created_at"`
	UpdatedAt string `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "user"
}
