package entity

import "time"

type Payment struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement" `
	BookingID int       `json:"booking_id" gorm:"not null"`
	Amount    string    `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Booking   Booking   `json:"booking,omitempty" gorm:"foreignKey:BookingID"` // Relasi: Payment belongs to Booking
}

type CreatePaymentReq struct {
	BookingID int       `json:"booking_id" validate:"required"`
	Amount    string    `json:"amount" validate:"required"`
	Status    string    `json:"status" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdatePaymentReq struct {
	ID        int       `json:"id" validate:"required"`
	BookingID int       `json:"booking_id" validate:"required"`
	Amount    string    `json:"amount" validate:"required"`
	Status    string    `json:"status" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PaymentRes struct {
	ID        int       `json:"id"`
	BookingID int       `json:"booking_id"`
	Amount    string    `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PaymentReport struct {
	TotalPayment int                   `json:"total_payment"`
	TotalAmount  float64               `json:"total_amount"`
	Status       []PaymentStatusDetail `json:"status"`
}

type PaymentStatusDetail struct {
	PaymentPaid     int     `json:"payment_paid"`
	AmountPaid      float64 `json:"amount_paid"`
	PaymentPending  int     `json:"payment_pending"`
	AmountPending   float64 `json:"amount_pending"`
	PaymentRefunded int     `json:"payment_refunded"`
	AmountRefunded  float64 `json:"amount_refunded"`
	PaymentFailed   int     `json:"payment_failed"`
	AmountFailed    float64 `json:"amount_failed"`
}
