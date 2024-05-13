package req

type NewUser struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	DocumentNumber string `json:"documentNumber"`
}
