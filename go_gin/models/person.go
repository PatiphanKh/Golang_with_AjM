package models

type Person struct {
	ID          int
	FirstName   string
	LastName    string
	Age         int
	Email       string
	Password    string
	PostAddress Address
}

type Address struct {
	HouseNo string
	City    string
}
