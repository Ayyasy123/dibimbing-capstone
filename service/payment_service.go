package service

import (
	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"github.com/Ayyasy123/dibimbing-capstone.git/repository"
)

type PaymentService interface {
	CreatePayment(req entity.CreatePaymentReq) (entity.Payment, error)
	GetPaymentByID(id int) (entity.Payment, error)
	UpdatePayment(req entity.UpdatePaymentReq) (entity.Payment, error)
	DeletePayment(id int) error
	GetAllPayments() ([]entity.Payment, error)
}

type paymentService struct {
	repo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) PaymentService {
	return &paymentService{repo}
}

func (s *paymentService) CreatePayment(req entity.CreatePaymentReq) (entity.Payment, error) {
	payment := entity.Payment{
		BookingID: req.BookingID,
		Amount:    req.Amount,
		Status:    "Pending", // Default status
	}
	return s.repo.Create(payment)
}

func (s *paymentService) GetPaymentByID(id int) (entity.Payment, error) {
	return s.repo.FindByID(id)
}

func (s *paymentService) UpdatePayment(req entity.UpdatePaymentReq) (entity.Payment, error) {
	payment, err := s.repo.FindByID(req.ID)
	if err != nil {
		return payment, err
	}

	payment.BookingID = req.BookingID
	payment.Amount = req.Amount
	payment.Status = req.Status

	return s.repo.Update(payment)
}

func (s *paymentService) DeletePayment(id int) error {
	return s.repo.Delete(id)
}

func (s *paymentService) GetAllPayments() ([]entity.Payment, error) {
	return s.repo.FindAll()
}
