package entity

import "time"

type Review struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	BookingID int       `json:"booking_id" gorm:"not null"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Booking   Booking   `json:"booking,omitempty" gorm:"foreignKey:BookingID"` // Relasi: Review belongs to Booking
}

type CreateReviewReq struct {
	BookingID int       `json:"booking_id" validate:"required"`
	Rating    int       `json:"rating" validate:"required"`
	Comment   string    `json:"comment" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateReviewReq struct {
	ID        int       `json:"id" validate:"required"`
	BookingID int       `json:"booking_id" validate:"required"`
	Rating    int       `json:"rating" validate:"required"`
	Comment   string    `json:"comment" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ReviewRes struct {
	ID        int       `json:"id"`
	BookingID int       `json:"booking_id"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
