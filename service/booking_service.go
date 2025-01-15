package service

import (
	"errors"

	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"github.com/Ayyasy123/dibimbing-capstone.git/repository"
)

type BookingService interface {
	CreateBooking(req entity.CreateBookingReq) (entity.Booking, error)
	GetBookingByID(id int) (entity.Booking, error)
	UpdateBooking(req entity.UpdateBookingReq) (entity.Booking, error)
	DeleteBooking(id int) error
	GetAllBookings() ([]entity.Booking, error)
	GetBookingsByUserID(userID int) ([]entity.BookingRes, error)
	GetBookingsByServiceID(serviceID int) ([]entity.BookingRes, error)
	UpdateBookingStatus(bookingID string, status string) error
}

type bookingService struct {
	repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) BookingService {
	return &bookingService{repo}
}

func (s *bookingService) CreateBooking(req entity.CreateBookingReq) (entity.Booking, error) {
	booking := entity.Booking{
		UserID:      req.UserID,
		ServiceID:   req.ServiceID,
		Date:        req.Date,
		Time:        req.Time,
		Status:      "Pending", // Default status
		Description: req.Description,
	}
	return s.repo.Create(booking)
}

func (s *bookingService) GetBookingByID(id int) (entity.Booking, error) {
	return s.repo.FindByID(id)
}

func (s *bookingService) UpdateBooking(req entity.UpdateBookingReq) (entity.Booking, error) {
	booking, err := s.repo.FindByID(req.ID)
	if err != nil {
		return booking, err
	}

	booking.UserID = req.UserID
	booking.ServiceID = req.ServiceID
	booking.Date = req.Date
	booking.Time = req.Time
	booking.Status = req.Status
	booking.Description = req.Description

	return s.repo.Update(booking)
}

func (s *bookingService) DeleteBooking(id int) error {
	return s.repo.Delete(id)
}

func (s *bookingService) GetAllBookings() ([]entity.Booking, error) {
	return s.repo.FindAll()
}

func (s *bookingService) GetBookingsByUserID(userID int) ([]entity.BookingRes, error) {
	bookings, err := s.repo.GetBookingsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var bookingRes []entity.BookingRes
	for _, booking := range bookings {
		bookingRes = append(bookingRes, entity.BookingRes{
			ID:          booking.ID,
			UserID:      booking.UserID,
			ServiceID:   booking.ServiceID,
			Date:        booking.Date,
			Time:        booking.Time,
			Status:      booking.Status,
			Description: booking.Description,
			CreatedAt:   booking.CreatedAt,
			UpdatedAt:   booking.UpdatedAt,
		})
	}

	return bookingRes, nil
}

func (s *bookingService) GetBookingsByServiceID(serviceID int) ([]entity.BookingRes, error) {
	bookings, err := s.repo.GetBookingsByServiceID(serviceID)
	if err != nil {
		return nil, err
	}

	var bookingRes []entity.BookingRes
	for _, booking := range bookings {
		bookingRes = append(bookingRes, entity.BookingRes{
			ID:          booking.ID,
			UserID:      booking.UserID,
			ServiceID:   booking.ServiceID,
			Date:        booking.Date,
			Time:        booking.Time,
			Status:      booking.Status,
			Description: booking.Description,
			CreatedAt:   booking.CreatedAt,
			UpdatedAt:   booking.UpdatedAt,
		})
	}

	return bookingRes, nil
}

func (s *bookingService) UpdateBookingStatus(bookingID string, status string) error {
	// Validasi status yang diperbolehkan
	allowedStatuses := map[string]bool{
		"Confirmed":   true,
		"In Progress": true,
		"Completed":   true,
		"Cancelled":   true,
		"Rescheduled": true,
	}

	if !allowedStatuses[status] {
		return errors.New("invalid status")
	}

	return s.repo.UpdateBookingStatus(bookingID, status)
}
