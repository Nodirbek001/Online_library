package entity

type User struct {
	ID        string `json:"id" gorm:"id"`
	Firstname string `json:"firstname" gorm:"firstname"`
	Lastname  string `json:"lastname" gorm:"lastname"`
	Username  string `json:"username" gorm:"username"`
	Password  string `json:"password" gorm:"password"`
	Passport  string `json:"passport" gorm:"passport"`
	Phone     string `json:"phone" gorm:"phone"`
	Role      string `json:"role" gorm:"phone"`
	Age       int16  `json:"age" gorm:"age"`
}
