package entity

import "time"

type Service struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      int       `json:"user_id"` // Foreign key ke User
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Cost        int       `json:"cost"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	User        User      `json:"user,omitempty" gorm:"foreignKey:UserID"`        // Relasi: Service belongs to User
	Bookings    []Booking `json:"bookings,omitempty" gorm:"foreignKey:ServiceID"` // Relasi: Service has many Bookings
}

type CreateServiceReq struct {
	UserID      int    `json:"user_id" validate:"required"` // Foreign key ke User
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Cost        int    `json:"cost" validate:"required"`
}

type UpdateServiceReq struct {
	ID          int    `json:"id" validate:"required"`
	UserID      int    `json:"user_id" validate:"required"` // Foreign key ke User
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Cost        int    `json:"cost" validate:"required"`
}

type ServiceRes struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"` // Foreign key ke User
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Cost        int       `json:"cost"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
