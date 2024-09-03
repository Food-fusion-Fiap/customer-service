package entities

type LgpdRemovalRequest struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	PhoneNumber    string `json:"phone_number"`
	PaymentHistory bool   `json:"payment_history"`
	CreatedAt      string `json:"createdAt"`
}
