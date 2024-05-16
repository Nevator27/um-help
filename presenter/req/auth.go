package req

type LoginRequest struct {
	DocumentNumber string `json:"documentNumber"`
	Password       string `jason:"password"`
}
