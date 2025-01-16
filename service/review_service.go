package service

import (
	"time"

	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"github.com/Ayyasy123/dibimbing-capstone.git/repository"
)

type ReviewService interface {
	CreateReview(req entity.CreateReviewReq) (entity.Review, error)
	GetReviewByID(id int) (entity.Review, error)
	UpdateReview(req entity.UpdateReviewReq) (entity.Review, error)
	DeleteReview(id int) error
	GetAllReviews() ([]entity.Review, error)
	GetReviewReport(startDate, endDate time.Time, serviceID int) (entity.ReviewReport, error)
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

func (s *reviewService) GetReviewReport(startDate, endDate time.Time, serviceID int) (entity.ReviewReport, error) {
	// Ambil total review (dengan atau tanpa filter tanggal dan service_id)
	totalReviews, err := s.repo.GetTotalReviews(startDate, endDate, serviceID)
	if err != nil {
		return entity.ReviewReport{}, err
	}

	// Ambil rata-rata rating (dengan atau tanpa filter tanggal dan service_id)
	averageRating, err := s.repo.GetAverageRating(startDate, endDate, serviceID)
	if err != nil {
		return entity.ReviewReport{}, err
	}

	// Ambil jumlah review untuk setiap rating (1-5)
	rating1Count, err := s.repo.GetReviewsByRating(1, startDate, endDate, serviceID)
	if err != nil {
		return entity.ReviewReport{}, err
	}

	rating2Count, err := s.repo.GetReviewsByRating(2, startDate, endDate, serviceID)
	if err != nil {
		return entity.ReviewReport{}, err
	}

	rating3Count, err := s.repo.GetReviewsByRating(3, startDate, endDate, serviceID)
	if err != nil {
		return entity.ReviewReport{}, err
	}

	rating4Count, err := s.repo.GetReviewsByRating(4, startDate, endDate, serviceID)
	if err != nil {
		return entity.ReviewReport{}, err
	}

	rating5Count, err := s.repo.GetReviewsByRating(5, startDate, endDate, serviceID)
	if err != nil {
		return entity.ReviewReport{}, err
	}

	// Buat response
	report := entity.ReviewReport{
		TotalReviews:  int(totalReviews),
		AverageRating: averageRating,
		RatingDistribution: []entity.RatingDistribution{
			{
				Rating: 1,
				Count:  int(rating1Count),
			},
			{
				Rating: 2,
				Count:  int(rating2Count),
			},
			{
				Rating: 3,
				Count:  int(rating3Count),
			},
			{
				Rating: 4,
				Count:  int(rating4Count),
			},
			{
				Rating: 5,
				Count:  int(rating5Count),
			},
		},
	}

	return report, nil
}
