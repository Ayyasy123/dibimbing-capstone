package service

import (
	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"github.com/Ayyasy123/dibimbing-capstone.git/repository"
)

type ReviewService interface {
	CreateReview(req entity.CreateReviewReq) (entity.Review, error)
	GetReviewByID(id int) (entity.Review, error)
	UpdateReview(req entity.UpdateReviewReq) (entity.Review, error)
	DeleteReview(id int) error
	GetAllReviews() ([]entity.Review, error)
}

type reviewService struct {
	repo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) ReviewService {
	return &reviewService{repo}
}

func (s *reviewService) CreateReview(req entity.CreateReviewReq) (entity.Review, error) {
	review := entity.Review{
		BookingID: req.BookingID,
		Rating:    req.Rating,
		Comment:   req.Comment,
	}
	return s.repo.Create(review)
}

func (s *reviewService) GetReviewByID(id int) (entity.Review, error) {
	return s.repo.FindByID(id)
}

func (s *reviewService) UpdateReview(req entity.UpdateReviewReq) (entity.Review, error) {
	review, err := s.repo.FindByID(req.ID)
	if err != nil {
		return review, err
	}

	review.BookingID = req.BookingID
	review.Rating = req.Rating
	review.Comment = req.Comment

	return s.repo.Update(review)
}

func (s *reviewService) DeleteReview(id int) error {
	return s.repo.Delete(id)
}

func (s *reviewService) GetAllReviews() ([]entity.Review, error) {
	return s.repo.FindAll()
}
