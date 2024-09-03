package dtos

type LGPDRemovalRequestDto struct {
	Name           string `json:"name" validate:"required, regexp=^[a-zA-ZÀ-ÿ'’\- ]{1,255}$"`
	Address        string `json:"address" validate:"required, regexp=^[a-zA-ZÀ-ÿ'’\- ]{1,255}$"`
	PhoneNumber    string `json:"phone_number" validate:"required, regexp=^\+?(\d{1,3})?[-.\s]?(\(?\d{2}\)?)?[-.\s]?\d{4,5}[-.\s]?\d{4}$"`
	PaymentHistory bool   `json:"payment_history" validate:"required,oneof=true false"`
}
