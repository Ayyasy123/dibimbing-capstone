package repository

import (
	"time"

	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment entity.Payment) (entity.Payment, error)
	FindByID(id int) (entity.Payment, error)
	Update(payment entity.Payment) (entity.Payment, error)
	Delete(id int) error
	FindAll(limit, offset int) ([]entity.Payment, error)
	UpdatePaymentStatus(paymentID string, status string) error
	GetTotalPayments(startDate, endDate time.Time, serviceID int) (int64, error)
	GetTotalAmount(startDate, endDate time.Time, serviceID int) (float64, error)
	GetPaymentsByStatus(status string, startDate, endDate time.Time, serviceID int) (int64, float64, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) Create(payment entity.Payment) (entity.Payment, error) {
	err := r.db.Create(&payment).Error
	return payment, err
}

func (r *paymentRepository) FindByID(id int) (entity.Payment, error) {
	var payment entity.Payment
	err := r.db.Preload("Booking").First(&payment, id).Error
	return payment, err
}

func (r *paymentRepository) Update(payment entity.Payment) (entity.Payment, error) {
	err := r.db.Save(&payment).Error
	return payment, err
}

func (r *paymentRepository) Delete(id int) error {
	err := r.db.Delete(&entity.Payment{}, id).Error
	return err
}

func (r *paymentRepository) FindAll(limit, offset int) ([]entity.Payment, error) {
	var payments []entity.Payment
	err := r.db.Limit(limit).Offset(offset).Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) UpdatePaymentStatus(paymentID string, status string) error {
	return r.db.Model(&entity.Payment{}).Where("id = ?", paymentID).Update("status", status).Error
}

func (r *paymentRepository) GetTotalPayments(startDate, endDate time.Time, serviceID int) (int64, error) {
	var total int64
	query := r.db.Model(&entity.Payment{})

	// Tambahkan filter tanggal jika startDate dan endDate tidak kosong
	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	// Tambahkan filter service_id jika diberikan
	if serviceID > 0 {
		query = query.Joins("JOIN bookings ON bookings.id = payments.booking_id").
			Where("bookings.service_id = ?", serviceID)
	}

	err := query.Count(&total).Error
	return total, err
}

func (r *paymentRepository) GetTotalAmount(startDate, endDate time.Time, serviceID int) (float64, error) {
	var totalAmount float64
	query := r.db.Model(&entity.Payment{}).Select("COALESCE(SUM(payments.amount), 0)")

	// Tambahkan filter tanggal jika startDate dan endDate tidak kosong
	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("payments.created_at BETWEEN ? AND ?", startDate, endDate)
	}

	// Tambahkan filter service_id jika diberikan
	if serviceID > 0 {
		query = query.Joins("JOIN bookings ON bookings.id = payments.booking_id").
			Where("bookings.service_id = ?", serviceID)
	}

	err := query.Scan(&totalAmount).Error
	return totalAmount, err
}

func (r *paymentRepository) GetPaymentsByStatus(status string, startDate, endDate time.Time, serviceID int) (int64, float64, error) {
	var count int64
	var totalAmount float64
	query := r.db.Model(&entity.Payment{}).Where("payments.status = ?", status)

	// Tambahkan filter tanggal jika startDate dan endDate tidak kosong
	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("payments.created_at BETWEEN ? AND ?", startDate, endDate)
	}

	// Tambahkan filter service_id jika diberikan
	if serviceID > 0 {
		query = query.Joins("JOIN bookings ON bookings.id = payments.booking_id").
			Where("bookings.service_id = ?", serviceID)
	}

	// Hitung jumlah pembayaran dan total uang
	err := query.Count(&count).Error
	if err != nil {
		return 0, 0, err
	}

	err = query.Select("COALESCE(SUM(payments.amount), 0)").Scan(&totalAmount).Error
	if err != nil {
		return 0, 0, err
	}

	return count, totalAmount, nil
}
