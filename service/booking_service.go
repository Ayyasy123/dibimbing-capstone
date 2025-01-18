package service

import (
	"errors"
	"time"

	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"github.com/Ayyasy123/dibimbing-capstone.git/repository"
)

type BookingService interface {
	CreateBooking(req entity.CreateBookingReq) (entity.Booking, error)
	GetBookingByID(id int) (entity.Booking, error)
	UpdateBooking(req entity.UpdateBookingReq) (entity.Booking, error)
	DeleteBooking(id int) error
	GetAllBookings(limit, offset int) ([]entity.Booking, error)
	GetBookingsByUserID(userID int) ([]entity.BookingRes, error)
	GetBookingsByServiceID(serviceID int) ([]entity.BookingRes, error)
	UpdateBookingStatus(bookingID string, status string) error
	GetBookingReport(startDate, endDate time.Time) (entity.BookingReport, error)
	GetAvailableDates(serviceID int, year int, month int) ([]time.Time, error)
}

type bookingService struct {
	repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) BookingService {
	return &bookingService{repo}
}

func (s *bookingService) CreateBooking(req entity.CreateBookingReq) (entity.Booking, error) {
	// Dapatkan tanggal hari ini (awal hari, 00:00:00)
	today := time.Now().UTC().Truncate(24 * time.Hour)

	// Jika tanggal booking adalah hari ini atau sebelumnya, tolak booking
	if req.Date.Before(today) || req.Date.Equal(today) {
		return entity.Booking{}, errors.New("booking cannot be accepted for today or past dates")
	}

	// Cek ketersediaan layanan pada tanggal yang diminta
	isAvailable, err := s.repo.CheckServiceAvailability(req.ServiceID, req.Date)
	if err != nil {
		return entity.Booking{}, err
	}
	if !isAvailable {
		return entity.Booking{}, errors.New("service is already booked on the requested date")
	}

	// Buat booking baru
	booking := entity.Booking{
		UserID:      req.UserID,
		ServiceID:   req.ServiceID,
		Date:        req.Date,
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
	// booking.Time = req.Time
	booking.Status = req.Status
	booking.Description = req.Description

	return s.repo.Update(booking)
}

func (s *bookingService) DeleteBooking(id int) error {
	return s.repo.Delete(id)
}

func (s *bookingService) GetAllBookings(limit, offset int) ([]entity.Booking, error) {
	return s.repo.FindAll(limit, offset)
}

func (s *bookingService) GetBookingsByUserID(userID int) ([]entity.BookingRes, error) {
	bookings, err := s.repo.GetBookingsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var bookingRes []entity.BookingRes
	for _, booking := range bookings {
		bookingRes = append(bookingRes, entity.BookingRes{
			ID:        booking.ID,
			UserID:    booking.UserID,
			ServiceID: booking.ServiceID,
			Date:      booking.Date,
			// Time:        booking.Time,
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
			ID:        booking.ID,
			UserID:    booking.UserID,
			ServiceID: booking.ServiceID,
			Date:      booking.Date,
			// Time:        booking.Time,
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
		// "Rescheduled": true,
	}

	if !allowedStatuses[status] {
		return errors.New("invalid status")
	}

	return s.repo.UpdateBookingStatus(bookingID, status)
}

func (s *bookingService) GetBookingReport(startDate, endDate time.Time) (entity.BookingReport, error) {
	// Get total bookings
	totalBooking, err := s.repo.GetTotalBookings(startDate, endDate)
	if err != nil {
		return entity.BookingReport{}, err
	}

	// Get total revenue
	totalRevenue, err := s.repo.GetTotalRevenue(startDate, endDate)
	if err != nil {
		return entity.BookingReport{}, err
	}

	// Define the statuses to query
	statuses := []string{"Pending", "In Progress", "Completed", "Canceled"}

	// Initialize the status details array
	statusDetails := []entity.BookingStatusDetail{}

	// Loop through each status and get the count and revenue
	for _, status := range statuses {
		count, revenue, err := s.repo.GetBookingsByStatus(status, startDate, endDate)
		if err != nil {
			return entity.BookingReport{}, err
		}

		// Append the status details to the array
		statusDetails = append(statusDetails, entity.BookingStatusDetail{
			BookingStatus: status,
			BookingCount:  int(count),
			Revenue:       revenue,
		})
	}

	// Create the report
	report := entity.BookingReport{
		TotalBooking: int(totalBooking),
		TotalRevenue: totalRevenue,
		Status:       statusDetails,
	}

	return report, nil
}

func (s *bookingService) GetAvailableDates(serviceID int, year int, month int) ([]time.Time, error) {
	// Dapatkan tanggal-tanggal yang sudah dipesan
	bookedDates, err := s.repo.GetBookedDates(serviceID, year, month)
	if err != nil {
		return nil, err
	}

	// Buat map untuk menyimpan tanggal yang sudah dipesan
	bookedMap := make(map[string]bool)
	for _, date := range bookedDates {
		bookedMap[date.Format("2006-01-02")] = true
	}

	// Hitung jumlah hari dalam bulan dan tahun yang diminta
	daysInMonth := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()

	// Buat slice untuk menyimpan tanggal yang tersedia
	var availableDates []time.Time

	// Dapatkan tanggal hari ini (awal hari, 00:00:00)
	today := time.Now().UTC().Truncate(24 * time.Hour)

	// Loop melalui setiap hari dalam bulan
	for day := 1; day <= daysInMonth; day++ {
		date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

		// Jika tanggal adalah hari ini atau sebelumnya, abaikan
		if date.Before(today) || date.Equal(today) {
			continue
		}

		// Jika tanggal tidak ada di bookedMap, tambahkan ke availableDates
		if !bookedMap[date.Format("2006-01-02")] {
			availableDates = append(availableDates, date)
		}
	}

	return availableDates, nil
}
