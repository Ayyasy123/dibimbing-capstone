package repository

import (
	"time"

	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	Create(review entity.Review) (entity.Review, error)
	FindByID(id int) (entity.Review, error)
	Update(review entity.Review) (entity.Review, error)
	Delete(id int) error
	FindAll() ([]entity.Review, error)
	GetTotalReviews(startDate, endDate time.Time, serviceID int) (int64, error)
	GetAverageRating(startDate, endDate time.Time, serviceID int) (float64, error)
	GetReviewsByRating(rating int, startDate, endDate time.Time, serviceID int) (int64, error)
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db}
}

func (r *reviewRepository) Create(review entity.Review) (entity.Review, error) {
	err := r.db.Create(&review).Error
	return review, err
}

func (r *reviewRepository) FindByID(id int) (entity.Review, error) {
	var review entity.Review
	err := r.db.Preload("Booking").First(&review, id).Error
	return review, err
}

func (r *reviewRepository) Update(review entity.Review) (entity.Review, error) {
	err := r.db.Save(&review).Error
	return review, err
}

func (r *reviewRepository) Delete(id int) error {
	err := r.db.Delete(&entity.Review{}, id).Error
	return err
}

func (r *reviewRepository) FindAll() ([]entity.Review, error) {
	var reviews []entity.Review
	err := r.db.Preload("Booking").Find(&reviews).Error
	return reviews, err
}

func (r *reviewRepository) GetTotalReviews(startDate, endDate time.Time, serviceID int) (int64, error) {
	var total int64
	query := r.db.Model(&entity.Review{})

	// Tambahkan filter tanggal jika startDate dan endDate tidak kosong
	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	// Tambahkan filter booking_id jika diberikan
	if serviceID > 0 {
		query = query.Joins("JOIN bookings ON bookings.id = reviews.booking_id").
			Where("bookings.service_id = ?", serviceID)
	}

	err := query.Count(&total).Error
	return total, err
}

func (r *reviewRepository) GetAverageRating(startDate, endDate time.Time, serviceID int) (float64, error) {
	var averageRating float64
	query := r.db.Model(&entity.Review{}).Select("COALESCE(AVG(rating), 0)")

	// Tambahkan filter tanggal jika startDate dan endDate tidak kosong
	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	// Tambahkan filter booking_id jika diberikan
	if serviceID > 0 {
		query = query.Joins("JOIN bookings ON bookings.id = reviews.booking_id").
			Where("bookings.service_id = ?", serviceID)
	}

	err := query.Scan(&averageRating).Error
	return averageRating, err
}

func (r *reviewRepository) GetReviewsByRating(rating int, startDate, endDate time.Time, serviceID int) (int64, error) {
	var count int64
	query := r.db.Model(&entity.Review{}).Where("rating = ?", rating)

	// Tambahkan filter tanggal jika startDate dan endDate tidak kosong
	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	// Tambahkan filter booking_id jika diberikan
	if serviceID > 0 {
		query = query.Joins("JOIN bookings ON bookings.id = reviews.booking_id").
			Where("bookings.service_id = ?", serviceID)
	}

	err := query.Count(&count).Error
	return count, err
}
