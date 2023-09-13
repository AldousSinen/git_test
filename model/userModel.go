package model

type TestFields struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Address   Address
}

type Address struct {
	Street   string `json:"street"`
	Barangay string `json:"barangay"`
	Contact  Contact
}

type Contact struct {
	Telephone int `json:"telephone"`
	Mobile    int `json:"mobile"`
}
