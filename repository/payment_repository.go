package repository

import (
	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment entity.Payment) (entity.Payment, error)
	FindByID(id int) (entity.Payment, error)
	Update(payment entity.Payment) (entity.Payment, error)
	Delete(id int) error
	FindAll() ([]entity.Payment, error)
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

func (r *paymentRepository) FindAll() ([]entity.Payment, error) {
	var payments []entity.Payment
	err := r.db.Preload("Booking").Find(&payments).Error
	return payments, err
}
