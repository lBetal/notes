package notes

type User struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password" binding:"required"`
}
