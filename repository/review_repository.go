package repository

import (
	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	Create(review entity.Review) (entity.Review, error)
	FindByID(id int) (entity.Review, error)
	Update(review entity.Review) (entity.Review, error)
	Delete(id int) error
	FindAll() ([]entity.Review, error)
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
