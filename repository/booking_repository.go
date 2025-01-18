package repository

import (
	"time"

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
	GetTotalBookings(startDate, endDate time.Time) (int64, error)
	GetTotalRevenue(startDate, endDate time.Time) (float64, error)
	GetBookingsByStatus(status string, startDate, endDate time.Time) (int64, float64, error)
	CheckServiceAvailability(serviceID int, date time.Time) (bool, error)
	GetBookedDates(serviceID int, year int, month int) ([]time.Time, error)
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

func (r *bookingRepository) GetTotalBookings(startDate, endDate time.Time) (int64, error) {
	var total int64
	query := r.db.Model(&entity.Booking{})

	// Tambahkan filter tanggal jika startDate dan endDate tidak kosong
	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("date BETWEEN ? AND ?", startDate, endDate)
	}

	err := query.Count(&total).Error
	return total, err
}

func (r *bookingRepository) GetTotalRevenue(startDate, endDate time.Time) (float64, error) {
	var totalRevenue float64
	query := r.db.Model(&entity.Booking{}).Joins("JOIN payments ON payments.booking_id = bookings.id").
		Select("COALESCE(SUM(payments.amount), 0)")

	// Tambahkan filter tanggal jika startDate dan endDate tidak kosong
	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("bookings.date BETWEEN ? AND ?", startDate, endDate)
	}

	err := query.Scan(&totalRevenue).Error
	return totalRevenue, err
}

func (r *bookingRepository) GetBookingsByStatus(status string, startDate, endDate time.Time) (int64, float64, error) {
	var count int64
	var totalRevenue float64

	// Query to count bookings by status
	query := r.db.Model(&entity.Booking{}).Where("bookings.status = ?", status)

	// Add date filter if provided
	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("bookings.date BETWEEN ? AND ?", startDate, endDate)
	}

	// Count the number of bookings
	err := query.Count(&count).Error
	if err != nil {
		return 0, 0, err
	}

	// Query to calculate total revenue for the given status
	revenueQuery := r.db.Model(&entity.Booking{}).
		Joins("JOIN payments ON payments.booking_id = bookings.id").
		Where("bookings.status = ?", status)

	// Add date filter if provided
	if !startDate.IsZero() && !endDate.IsZero() {
		revenueQuery = revenueQuery.Where("bookings.date BETWEEN ? AND ?", startDate, endDate)
	}

	// Sum the payment amounts
	err = revenueQuery.Select("COALESCE(SUM(payments.amount), 0)").Scan(&totalRevenue).Error
	if err != nil {
		return 0, 0, err
	}

	return count, totalRevenue, nil
}

func (r *bookingRepository) CheckServiceAvailability(serviceID int, date time.Time) (bool, error) {
	var count int64
	err := r.db.Model(&entity.Booking{}).
		Where("service_id = ? AND DATE(date) = ?", serviceID, date.Format("2006-01-02")).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (r *bookingRepository) GetBookedDates(serviceID int, year int, month int) ([]time.Time, error) {
	var bookedDates []time.Time

	// Query untuk mendapatkan tanggal-tanggal yang sudah dipesan
	err := r.db.Model(&entity.Booking{}).
		Where("service_id = ? AND YEAR(date) = ? AND MONTH(date) = ?", serviceID, year, month).
		Pluck("date", &bookedDates).Error

	if err != nil {
		return nil, err
	}

	return bookedDates, nil
}
