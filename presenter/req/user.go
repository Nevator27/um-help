package presenter_req

type CreateUserAccount struct {
	FirstName string `json:"firstName"` //annotation
	LastName  string `json:"lastName"`
	TaxID     string `json:"taxId"`
}
