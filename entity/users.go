package entity

import "time"

type User struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Role         string    `gorm:"type:ENUM('admin', 'user', 'technician');default:'user'" json:"role"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	Expertise    string    `json:"expertise"`
	Availability string    `json:"availability"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Services     []Service `json:"services,omitempty" gorm:"foreignKey:UserID"` // Relasi: User has many Services
	Bookings     []Booking `json:"bookings,omitempty" gorm:"foreignKey:UserID"` // Relasi: User has many Bookings
}

type RegisterUserReq struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUserReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterAsTechnicianReq struct {
	ID           int    `json:"id" validate:"required"`
	Role         string `json:"role" validate:"required"`
	Address      string `json:"address" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	Expertise    string `json:"expertise" validate:"required"`
	Availability string `json:"availability" validate:"required"`
}

type UpdateUserReq struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type UpdateTechnicianReq struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	Expertise    string `json:"expertise"`
	Availability string `json:"availability"`
}

type UserRes struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TechnicianRes struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	Expertise    string    `json:"expertise"`
	Availability string    `json:"availability"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
