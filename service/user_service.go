package service

import (
	"errors"

	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"github.com/Ayyasy123/dibimbing-capstone.git/repository"
	"github.com/Ayyasy123/dibimbing-capstone.git/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req *entity.RegisterUserReq) (*entity.UserRes, error)
	Login(req *entity.LoginUserReq) (*entity.UserRes, string, error)
	GetUserByID(id int) (*entity.UserRes, error)
	GetAllUsers(limit, offset int) ([]*entity.UserRes, error)
	UpdateUser(req *entity.UpdateUserReq) (*entity.UserRes, error)
	RegisterAsTechnician(req *entity.RegisterAsTechnicianReq) (*entity.TechnicianRes, error)
	UpdateTechnician(req *entity.UpdateTechnicianReq) (*entity.TechnicianRes, error)
	DeleteUser(id int) error
	RegisterAsAdmin(req *entity.RegisterUserReq) (*entity.UserRes, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) Register(req *entity.RegisterUserReq) (*entity.UserRes, error) {
	exists, err := s.userRepository.IsEmailExists(req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "user", // Default role
	}

	err = s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	userRes := &entity.UserRes{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userRes, nil
}

func (s *userService) Login(req *entity.LoginUserReq) (*entity.UserRes, string, error) {
	// Cari user berdasarkan email
	user, err := s.userRepository.FindUserByEmail(req.Email)
	if err != nil {
		return nil, "", errors.New("user not found")
	}

	// Bandingkan password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, "", errors.New("invalid password")
	}

	// Generate token JWT
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	userRes := &entity.UserRes{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userRes, token, nil
}

func (s *userService) GetUserByID(id int) (*entity.UserRes, error) {
	user, err := s.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	userRes := &entity.UserRes{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userRes, nil
}

func (s *userService) GetAllUsers(limit, offset int) ([]*entity.UserRes, error) {
	users, err := s.userRepository.FindAll(limit, offset)
	if err != nil {
		return nil, err
	}

	var userRes []*entity.UserRes
	for _, user := range users {
		userRes = append(userRes, &entity.UserRes{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return userRes, nil
}

func (s *userService) UpdateUser(req *entity.UpdateUserReq) (*entity.UserRes, error) {
	user, err := s.userRepository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}
	if req.Role != "" {
		user.Role = req.Role
	}
	if req.Address != "" {
		user.Address = req.Address
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}

	err = s.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	userRes := &entity.UserRes{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userRes, nil
}

func (s *userService) RegisterAsTechnician(req *entity.RegisterAsTechnicianReq) (*entity.TechnicianRes, error) {
	user, err := s.userRepository.FindByID(req.ID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Update user role and technician-specific fields
	user.Role = "technician"
	user.Address = req.Address
	user.Phone = req.Phone
	user.Expertise = req.Expertise
	user.Availability = req.Availability

	err = s.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	technicianRes := &entity.TechnicianRes{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Role:         user.Role,
		Address:      user.Address,
		Phone:        user.Phone,
		Expertise:    user.Expertise,
		Availability: user.Availability,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	return technicianRes, nil
}

func (s *userService) UpdateTechnician(req *entity.UpdateTechnicianReq) (*entity.TechnicianRes, error) {
	user, err := s.userRepository.FindByID(req.ID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}
	if req.Role != "" {
		user.Role = req.Role
	}
	if req.Address != "" {
		user.Address = req.Address
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.Expertise != "" {
		user.Expertise = req.Expertise
	}
	if req.Availability != "" {
		user.Availability = req.Availability
	}

	err = s.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	technicianRes := &entity.TechnicianRes{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Role:         user.Role,
		Address:      user.Address,
		Phone:        user.Phone,
		Expertise:    user.Expertise,
		Availability: user.Availability,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	return technicianRes, nil
}

func (s *userService) DeleteUser(id int) error {
	return s.userRepository.Delete(id)
}

func (s *userService) RegisterAsAdmin(req *entity.RegisterUserReq) (*entity.UserRes, error) {
	// Cek apakah email sudah terdaftar
	exists, err := s.userRepository.IsEmailExists(req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Buat user dengan role admin
	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "admin", // Role di-set sebagai admin
	}

	// Simpan ke database
	err = s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	// Konversi ke UserRes
	userRes := &entity.UserRes{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userRes, nil
}
