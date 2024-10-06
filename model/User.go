package model

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
	Age       uint   `json:"age"`
	Image     string `json:"image"`
	Country   string `json:"country"`
}
