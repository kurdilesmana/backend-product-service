package userModel

type User struct {
	ID          int64  `json:"id" gorm:"column:id;type:serial;" `
	Name        string `json:"name" gorm:"column:name;type:varchar(100);null;" `
	Email       string `json:"email" gorm:"column:email;type:varchar(100);null;" `
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number;type:varchar(20);null;" `
	Password    string `json:"password" gorm:"column:password;type:text;null;" `
}

func (b *User) TableName() string {
	return "public.users"
}
