package repository

import (
	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"gorm.io/gorm"
)

type BookingRepository interface {
	Create(booking entity.Booking) (entity.Booking, error)
	FindByID(id int) (entity.Booking, error)
	FindAll() ([]entity.Booking, error)
	Update(booking entity.Booking) (entity.Booking, error)
	Delete(id int) error
	GetBookingsByUserID(userID int) ([]entity.Booking, error)
	GetBookingsByServiceID(serviceID int) ([]entity.Booking, error)
	UpdateBookingStatus(bookingID string, status string) error
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db}
}

func (r *bookingRepository) Create(booking entity.Booking) (entity.Booking, error) {
	err := r.db.Create(&booking).Error
	return booking, err
}

func (r *bookingRepository) FindByID(id int) (entity.Booking, error) {
	var booking entity.Booking
	err := r.db.Preload("User").Preload("Service").First(&booking, id).Error
	return booking, err
}

func (r *bookingRepository) FindAll() ([]entity.Booking, error) {
	var bookings []entity.Booking
	err := r.db.Preload("User").Preload("Service").Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepository) Update(booking entity.Booking) (entity.Booking, error) {
	err := r.db.Save(&booking).Error
	return booking, err
}

func (r *bookingRepository) Delete(id int) error {
	err := r.db.Delete(&entity.Booking{}, id).Error
	return err
}

func (r *bookingRepository) GetBookingsByUserID(userID int) ([]entity.Booking, error) {
	var bookings []entity.Booking
	err := r.db.Where("user_id = ?", userID).Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepository) GetBookingsByServiceID(serviceID int) ([]entity.Booking, error) {
	var bookings []entity.Booking
	err := r.db.Where("service_id = ?", serviceID).Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepository) UpdateBookingStatus(bookingID string, status string) error {
	return r.db.Model(&entity.Booking{}).Where("id = ?", bookingID).Update("status", status).Error
}

func (r *bookingRepository) CancelBooking(bookingID string) error {
	return r.db.Model(&entity.Booking{}).Where("id = ?", bookingID).Update("status", "Cancelled").Error
}
