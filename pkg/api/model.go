package api

type Alert struct {
	Money    float32 `json:"money"`
	Currency string  `json:"currency"`
	Operator string  `json:"threshold"`
	Email    string  `json:"email"`
}
