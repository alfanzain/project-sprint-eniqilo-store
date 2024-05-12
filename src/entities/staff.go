package entities

type Staff struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"-"`
	Token       string `json:"token"`
}
