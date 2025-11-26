package dto

type Person struct {
	ID          int
	FirstName   string `json:"f_name"`
	LastName    string `json:"l_name"`
	Age         int
	Email       string
	Password    string  `json:"-"`
	PostAddress Address `json:"Address"`
}

type Address struct {
	HouseNo string
	City    string
}
