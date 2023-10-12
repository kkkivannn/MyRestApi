package api

type Client struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required" db:"name"`
	Email    string `json:"email" bd:"email"`
	Age      int    `json:"age" db:"age"`
	Phone    string `json:"phone" db:"phone_number"`
	Password string `json:"password" binding:"required"`
}

type Profile struct {
	Name  string `json:"name" db:"name"`
	Age   string `json:"age" db:"age"`
	Phone string `json:"phone" db:"phone_number"`
}
