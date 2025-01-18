package service

import (
	"errors"
	"time"

	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"github.com/Ayyasy123/dibimbing-capstone.git/repository"
)

type PaymentService interface {
	CreatePayment(req entity.CreatePaymentReq) (entity.Payment, error)
	GetPaymentByID(id int) (entity.Payment, error)
	UpdatePayment(req entity.UpdatePaymentReq) (entity.Payment, error)
	DeletePayment(id int) error
	GetAllPayments(limit, offset int) ([]entity.Payment, error)
	UpdatePaymentStatus(paymentID string, status string) error
	GetPaymentReport(startDate, endDate time.Time, serviceID int) (entity.PaymentReport, error)
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

func (s *paymentService) GetAllPayments(limit, offset int) ([]entity.Payment, error) {
	return s.repo.FindAll(limit, offset)
}

func (s *paymentService) UpdatePaymentStatus(paymentID string, status string) error {
	// Validasi status yang diperbolehkan
	allowedStatuses := map[string]bool{
		"Paid":     true,
		"Failed":   true,
		"Refunded": true,
		"Expired":  true,
	}

	if !allowedStatuses[status] {
		return errors.New("invalid status")
	}

	return s.repo.UpdatePaymentStatus(paymentID, status)
}

func (s *paymentService) GetPaymentReport(startDate, endDate time.Time, serviceID int) (entity.PaymentReport, error) {
	// Ambil total pembayaran
	totalPayment, err := s.repo.GetTotalPayments(startDate, endDate, serviceID)
	if err != nil {
		return entity.PaymentReport{}, err
	}

	// Ambil total jumlah uang
	totalAmount, err := s.repo.GetTotalAmount(startDate, endDate, serviceID)
	if err != nil {
		return entity.PaymentReport{}, err
	}

	// Ambil jumlah dan total uang untuk setiap status
	paidCount, paidAmount, err := s.repo.GetPaymentsByStatus("paid", startDate, endDate, serviceID)
	if err != nil {
		return entity.PaymentReport{}, err
	}

	pendingCount, pendingAmount, err := s.repo.GetPaymentsByStatus("pending", startDate, endDate, serviceID)
	if err != nil {
		return entity.PaymentReport{}, err
	}

	refundedCount, refundedAmount, err := s.repo.GetPaymentsByStatus("refunded", startDate, endDate, serviceID)
	if err != nil {
		return entity.PaymentReport{}, err
	}

	failedCount, failedAmount, err := s.repo.GetPaymentsByStatus("failed", startDate, endDate, serviceID)
	if err != nil {
		return entity.PaymentReport{}, err
	}

	// Buat response
	report := entity.PaymentReport{
		TotalPayment: int(totalPayment),
		TotalAmount:  totalAmount,
		Status: []entity.PaymentStatusDetail{
			{
				PaymentStatus: "Paid",
				PaymentCount:  int(paidCount),
				Amount:        paidAmount,
			},
			{
				PaymentStatus: "Pending",
				PaymentCount:  int(pendingCount),
				Amount:        pendingAmount,
			},
			{
				PaymentStatus: "Refunded",
				PaymentCount:  int(refundedCount),
				Amount:        refundedAmount,
			},
			{
				PaymentStatus: "Failed",
				PaymentCount:  int(failedCount),
				Amount:        failedAmount,
			},
		},
	}

	return report, nil
}
