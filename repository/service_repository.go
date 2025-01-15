package repository

import (
	"strings"

	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"gorm.io/gorm"
)

type ServiceRepository interface {
	Create(service *entity.Service) error
	FindByID(id int) (*entity.Service, error)
	FindAll() ([]entity.Service, error)
	Update(service *entity.Service) error
	Delete(id int) error
	GetServicesByUserID(userID int) ([]entity.Service, error)
	SearchServices(searchQuery string, minPrice, maxPrice int) ([]entity.Service, error)
}

type serviceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) ServiceRepository {
	return &serviceRepository{db}
}

func (r *serviceRepository) Create(service *entity.Service) error {
	return r.db.Create(service).Error
}

func (r *serviceRepository) FindByID(id int) (*entity.Service, error) {
	var service entity.Service
	err := r.db.Preload("User").Preload("Bookings").First(&service, id).Error
	return &service, err
}

func (r *serviceRepository) FindAll() ([]entity.Service, error) {
	var services []entity.Service
	err := r.db.Preload("User").Preload("Bookings").Find(&services).Error
	return services, err
}

func (r *serviceRepository) Update(service *entity.Service) error {
	return r.db.Save(service).Error
}

func (r *serviceRepository) Delete(id int) error {
	return r.db.Delete(&entity.Service{}, id).Error
}

func (r *serviceRepository) GetServicesByUserID(userID int) ([]entity.Service, error) {
	var services []entity.Service
	err := r.db.Where("user_id = ?", userID).Find(&services).Error
	return services, err
}

func (r *serviceRepository) SearchServices(searchQuery string, minPrice, maxPrice int) ([]entity.Service, error) {
	var services []entity.Service
	query := r.db.Joins("JOIN users ON users.id = services.user_id")

	if searchQuery != "" {
		searchQuery = strings.ToLower(searchQuery) // Ubah ke lowercase untuk pencarian case-insensitive
		query = query.Where(
			"LOWER(users.address) LIKE ? OR LOWER(services.name) LIKE ? OR LOWER(services.description) LIKE ?",
			"%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%",
		)
	}

	query = query.Where("cost BETWEEN ? AND ?", minPrice, maxPrice)

	err := query.Preload("User").Find(&services).Error
	return services, err
}
