package entity

import (
	"time"
)

type Booking struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int       `json:"user_id" gorm:"not null"`
	ServiceID int       `json:"service_id" gorm:"not null"`
	Date      time.Time `json:"date" gorm:"type:date"`
	// Time        time.Time `json:"time" gorm:"type:time"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	User        User      `json:"user,omitempty" gorm:"foreignKey:UserID"`       // Relasi: Booking belongs to User
	Service     Service   `json:"service,omitempty" gorm:"foreignKey:ServiceID"` // Relasi: Booking belongs to Service
}

type CreateBookingReq struct {
	UserID    int       `json:"user_id" validate:"required"`
	ServiceID int       `json:"service_id" validate:"required"`
	Date      time.Time `json:"date" validate:"required"`
	// Time        time.Time `json:"time" validate:"required"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

type UpdateBookingReq struct {
	ID        int       `json:"id" validate:"required"`
	UserID    int       `json:"user_id" validate:"required"`
	ServiceID int       `json:"service_id" validate:"required"`
	Date      time.Time `json:"date" validate:"required"`
	// Time        time.Time `json:"time" validate:"required"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

type BookingRes struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id" `
	ServiceID int       `json:"service_id" `
	Date      time.Time `json:"date"`
	// Time        time.Time `json:"time"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookingReport struct {
	TotalBooking int                   `json:"total_booking"`
	TotalRevenue float64               `json:"total_revenue"`
	Status       []BookingStatusDetail `json:"status"`
}

type BookingStatusDetail struct {
	BookingPending    int     `json:"booking_pending"`
	RevenuePending    float64 `json:"revenue_pending"`
	BookingInProgress int     `json:"booking_in_progress"`
	RevenueInProgress float64 `json:"revenue_in_progress"`
	BookingCompleted  int     `json:"booking_completed"`
	RevenueCompleted  float64 `json:"revenue_completed"`
	BookingCanceled   int     `json:"booking_canceled"`
	RevenueCanceled   float64 `json:"revenue_canceled"`
}
