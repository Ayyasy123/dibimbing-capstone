package repository

import (
	"strings"

	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"gorm.io/gorm"
)

type ServiceRepository interface {
	Create(service *entity.Service) error
	FindByID(id int) (*entity.Service, error)
	FindAll(limit, offset int) ([]entity.Service, error)
	Update(service *entity.Service) error
	Delete(id int) error
	GetServicesByUserID(userID int) ([]entity.Service, error)
	SearchServices(searchQuery string, minPrice, maxPrice int) ([]entity.Service, error)
	GetServiceCostDistribution(startDate, endDate string) (map[string]int, error)
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

func (r *serviceRepository) FindAll(limit, offset int) ([]entity.Service, error) {
	var services []entity.Service
	err := r.db.Limit(limit).Offset(offset).Find(&services).Error
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

func (r *serviceRepository) GetServiceCostDistribution(startDate, endDate string) (map[string]int, error) {
	var costDistribution []struct {
		CostRange string
		Count     int
	}

	query := r.db.Model(&entity.Service{})
	if startDate != "" && endDate != "" {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	err := query.Select("CASE " +
		"WHEN cost BETWEEN 50000 AND 100000 THEN '50000-100000' " +
		"WHEN cost BETWEEN 100001 AND 300000 THEN '100001-300000' " +
		"WHEN cost BETWEEN 300001 AND 500000 THEN '300001-500000' " +
		"WHEN cost BETWEEN 500001 AND 700000 THEN '500001-700000' " +
		"WHEN cost BETWEEN 700001 AND 1000000 THEN '700001-1000000' " +
		"ELSE '1000001+' END as cost_range, count(*) as count").
		Group("cost_range").
		Scan(&costDistribution).Error
	if err != nil {
		return nil, err
	}

	distributionMap := make(map[string]int)
	for _, cd := range costDistribution {
		distributionMap[cd.CostRange] = cd.Count
	}

	return distributionMap, nil
}
