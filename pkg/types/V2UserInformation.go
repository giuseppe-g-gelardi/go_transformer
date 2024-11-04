package types

type V2UserInformation struct {
	ID                 string             `json:"id"`
	AccountInformation AccountInformation `json:"accountInformation"`
	UserInformation    UserInformation    `json:"userInformation"`
	ContactInformation ContactInformation `json:"contact"`
	Tags               []string           `json:"tags"`
	Profile            string             `json:"profile"`
}

type AccountInformation struct {
	IsActive   bool   `json:"isActive"`
	Registered string `json:"registered"` // date
	Balance    string `json:"balance"`    // float64
}

type UserInformation struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	EyeColor  string `json:"eyeColor"`
	Picture   string `json:"picture"`
	Company   string `json:"company"`
}

type ContactInformation struct {
	Email   string  `json:"email"` // email .. must include '@'?
	Phone   string  `json:"phone"` // phone
	Address Address `json:"address"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    int    `json:"zip"`
}
