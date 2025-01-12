package repository

import (
	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"gorm.io/gorm"
)

type ServiceRepository interface {
	Create(service *entity.Service) error
	FindByID(id int) (*entity.Service, error)
	FindAll() ([]entity.Service, error)
	Update(service *entity.Service) error
	Delete(id int) error
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
