package dtos

type LGPDRemovalRequestDto struct {
	Name           string `json:"name" validate:"nonzero,min=10,max=255"`
	Address        string `json:"address" validate:"nonzero,min=10,max=255"`
	PhoneNumber    string `json:"phone_number" validate:"len=11"`
	PaymentHistory bool   `json:"payment_history" validate:"nonnil"`
}
