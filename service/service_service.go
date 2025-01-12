package service

import (
	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"github.com/Ayyasy123/dibimbing-capstone.git/repository"
)

type ServiceService interface {
	CreateService(req entity.CreateServiceReq) (*entity.Service, error)
	GetServiceByID(id int) (*entity.Service, error)
	UpdateService(req entity.UpdateServiceReq) (*entity.Service, error)
	DeleteService(id int) error
	GetAllServices() ([]entity.Service, error)
}

type serviceService struct {
	serviceRepo repository.ServiceRepository
}

func NewServiceService(serviceRepo repository.ServiceRepository) ServiceService {
	return &serviceService{serviceRepo}
}

func (s *serviceService) CreateService(req entity.CreateServiceReq) (*entity.Service, error) {
	service := &entity.Service{
		UserID:      req.UserID,
		Name:        req.Name,
		Description: req.Description,
		Cost:        req.Cost,
	}
	err := s.serviceRepo.Create(service)
	return service, err
}

func (s *serviceService) GetServiceByID(id int) (*entity.Service, error) {
	return s.serviceRepo.FindByID(id)
}

func (s *serviceService) UpdateService(req entity.UpdateServiceReq) (*entity.Service, error) {
	service, err := s.serviceRepo.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	service.UserID = req.UserID
	service.Name = req.Name
	service.Description = req.Description
	service.Cost = req.Cost

	err = s.serviceRepo.Update(service)
	return service, err
}

func (s *serviceService) DeleteService(id int) error {
	return s.serviceRepo.Delete(id)
}

func (s *serviceService) GetAllServices() ([]entity.Service, error) {
	return s.serviceRepo.FindAll()
}