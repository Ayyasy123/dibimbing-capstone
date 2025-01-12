package service

import (
	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"github.com/Ayyasy123/dibimbing-capstone.git/repository"
)

type BookingService interface {
	CreateBooking(req entity.CreateBookingReq) (entity.Booking, error)
	GetBookingByID(id int) (entity.Booking, error)
	UpdateBooking(req entity.UpdateBookingReq) (entity.Booking, error)
	DeleteBooking(id int) error
	GetAllBookings() ([]entity.Booking, error)
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
