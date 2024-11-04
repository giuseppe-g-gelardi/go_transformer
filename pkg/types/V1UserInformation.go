package types

type V1UserInformation struct {
	ID         string    `json:"id"`
	IsActive   bool      `json:"isActive"`
	Balance    string    `json:"balance"` // float64
	Picture    string    `json:"picture"`
	Age        int       `json:"age"`
	EyeColor   string    `json:"eyeColor"` // enum (blue, brown, green)
	Name       string    `json:"name"`
	Gender     string    `json:"gender"`
	Company    string    `json:"company"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	About      string    `json:"about"`
	Registered string    `json:"registered"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Tags       []string  `json:"tags"`
	Friends    []friends `json:"friends"`
}

type friends struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
