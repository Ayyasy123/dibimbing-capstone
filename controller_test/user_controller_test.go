package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Ayyasy123/dibimbing-capstone.git/controller"
	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"github.com/Ayyasy123/dibimbing-capstone.git/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserController_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	userController := controller.NewUserController(mockUserService)

	testCases := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func()
		expectedStatus int
		expectedBody   *entity.UserRes
	}{
		{
			name: "Success - User registered successfully",
			requestBody: entity.RegisterUserReq{
				Name:     "John Doe",
				Email:    "john@example.com",
				Password: "password123",
			},
			mockSetup: func() {
				mockUserService.EXPECT().Register(gomock.Any()).Return(&entity.UserRes{
					ID:        1,
					Name:      "John Doe",
					Email:     "john@example.com",
					Role:      "user",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody: &entity.UserRes{
				ID:        1,
				Name:      "John Doe",
				Email:     "john@example.com",
				Role:      "user",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
		{
			name: "Failure - Email already registered",
			requestBody: entity.RegisterUserReq{
				Name:     "Jane Doe",
				Email:    "jane@example.com",
				Password: "password123",
			},
			mockSetup: func() {
				mockUserService.EXPECT().Register(gomock.Any()).Return(nil, errors.New("email already registered"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			reqBody, err := json.Marshal(tc.requestBody)
			assert.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)
			ctx.Request = req

			userController.Register(ctx)

			assert.Equal(t, tc.expectedStatus, res.Code)

			if tc.expectedBody != nil {
				var responseBody entity.UserRes
				err = json.Unmarshal(res.Body.Bytes(), &responseBody)
				assert.NoError(t, err)

				assert.Equal(t, tc.expectedBody.ID, responseBody.ID)
				assert.Equal(t, tc.expectedBody.Name, responseBody.Name)
				assert.Equal(t, tc.expectedBody.Email, responseBody.Email)
				assert.Equal(t, tc.expectedBody.Role, responseBody.Role)
				assert.WithinDuration(t, tc.expectedBody.CreatedAt, responseBody.CreatedAt, time.Second)
				assert.WithinDuration(t, tc.expectedBody.UpdatedAt, responseBody.UpdatedAt, time.Second)
			}
		})
	}
}

func TestUserController_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	userController := controller.NewUserController(mockUserService)

	testCases := []struct {
		name           string
		userID         string
		mockSetup      func()
		expectedStatus int
		expectedBody   *entity.UserRes
	}{
		{
			name:   "Success - User found",
			userID: "1",
			mockSetup: func() {
				mockUserService.EXPECT().GetUserByID(1).Return(&entity.UserRes{
					ID:        1,
					Name:      "John Doe",
					Email:     "john@example.com",
					Role:      "user",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &entity.UserRes{
				ID:        1,
				Name:      "John Doe",
				Email:     "john@example.com",
				Role:      "user",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
		{
			name:   "Failure - User not found",
			userID: "999",
			mockSetup: func() {
				mockUserService.EXPECT().GetUserByID(999).Return(nil, errors.New("user not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			req, err := http.NewRequest(http.MethodGet, "/users/"+tc.userID, nil)
			assert.NoError(t, err)

			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)
			ctx.Request = req
			ctx.Params = gin.Params{gin.Param{Key: "id", Value: tc.userID}}

			userController.GetUserByID(ctx)

			assert.Equal(t, tc.expectedStatus, res.Code)

			if tc.expectedBody != nil {
				var responseBody entity.UserRes
				err = json.Unmarshal(res.Body.Bytes(), &responseBody)
				assert.NoError(t, err)

				assert.Equal(t, tc.expectedBody.ID, responseBody.ID)
				assert.Equal(t, tc.expectedBody.Name, responseBody.Name)
				assert.Equal(t, tc.expectedBody.Email, responseBody.Email)
				assert.Equal(t, tc.expectedBody.Role, responseBody.Role)
				assert.WithinDuration(t, tc.expectedBody.CreatedAt, responseBody.CreatedAt, time.Second)
				assert.WithinDuration(t, tc.expectedBody.UpdatedAt, responseBody.UpdatedAt, time.Second)
			}
		})
	}
}

func TestUserController_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	userController := controller.NewUserController(mockUserService)

	testCases := []struct {
		name           string
		mockSetup      func()
		expectedStatus int
		expectedBody   []*entity.UserRes
	}{
		{
			name: "Success - Get all users",
			mockSetup: func() {
				mockUserService.EXPECT().GetAllUsers().Return([]*entity.UserRes{
					{
						ID:        1,
						Name:      "John Doe",
						Email:     "john@example.com",
						Role:      "user",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
					{
						ID:        2,
						Name:      "Jane Doe",
						Email:     "jane@example.com",
						Role:      "admin",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []*entity.UserRes{
				{
					ID:        1,
					Name:      "John Doe",
					Email:     "john@example.com",
					Role:      "user",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					ID:        2,
					Name:      "Jane Doe",
					Email:     "jane@example.com",
					Role:      "admin",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
		},
		{
			name: "Failure - Internal server error",
			mockSetup: func() {
				mockUserService.EXPECT().GetAllUsers().Return(nil, errors.New("internal server error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			req, err := http.NewRequest(http.MethodGet, "/users", nil)
			assert.NoError(t, err)

			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)
			ctx.Request = req

			userController.GetAllUsers(ctx)

			assert.Equal(t, tc.expectedStatus, res.Code)

			if tc.expectedBody != nil {
				var responseBody []*entity.UserRes
				err = json.Unmarshal(res.Body.Bytes(), &responseBody)
				assert.NoError(t, err)

				assert.Equal(t, len(tc.expectedBody), len(responseBody))
				for i, expectedUser := range tc.expectedBody {
					assert.Equal(t, expectedUser.ID, responseBody[i].ID)
					assert.Equal(t, expectedUser.Name, responseBody[i].Name)
					assert.Equal(t, expectedUser.Email, responseBody[i].Email)
					assert.Equal(t, expectedUser.Role, responseBody[i].Role)
					assert.WithinDuration(t, expectedUser.CreatedAt, responseBody[i].CreatedAt, time.Second)
					assert.WithinDuration(t, expectedUser.UpdatedAt, responseBody[i].UpdatedAt, time.Second)
				}
			}
		})
	}
}

func TestUserController_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	userController := controller.NewUserController(mockUserService)

	testCases := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func()
		expectedStatus int
		expectedBody   *entity.UserRes
	}{
		{
			name: "Success - User updated successfully",
			requestBody: entity.UpdateUserReq{
				ID:       1,
				Name:     "John Doe Updated",
				Email:    "john.updated@example.com",
				Password: "newpassword123",
				Role:     "admin",
				Address:  "123 Updated St",
				Phone:    "1234567890",
			},
			mockSetup: func() {
				mockUserService.EXPECT().UpdateUser(gomock.Any()).Return(&entity.UserRes{
					ID:        1,
					Name:      "John Doe Updated",
					Email:     "john.updated@example.com",
					Role:      "admin",
					Address:   "123 Updated St",
					Phone:     "1234567890",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &entity.UserRes{
				ID:        1,
				Name:      "John Doe Updated",
				Email:     "john.updated@example.com",
				Role:      "admin",
				Address:   "123 Updated St",
				Phone:     "1234567890",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
		{
			name: "Failure - User not found",
			requestBody: entity.UpdateUserReq{
				ID: 999,
			},
			mockSetup: func() {
				mockUserService.EXPECT().UpdateUser(gomock.Any()).Return(nil, errors.New("user not found"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			reqBody, err := json.Marshal(tc.requestBody)
			assert.NoError(t, err)

			req, err := http.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)
			ctx.Request = req

			userController.UpdateUser(ctx)

			assert.Equal(t, tc.expectedStatus, res.Code)

			if tc.expectedBody != nil {
				var responseBody entity.UserRes
				err = json.Unmarshal(res.Body.Bytes(), &responseBody)
				assert.NoError(t, err)

				assert.Equal(t, tc.expectedBody.ID, responseBody.ID)
				assert.Equal(t, tc.expectedBody.Name, responseBody.Name)
				assert.Equal(t, tc.expectedBody.Email, responseBody.Email)
				assert.Equal(t, tc.expectedBody.Role, responseBody.Role)
				assert.Equal(t, tc.expectedBody.Address, responseBody.Address)
				assert.Equal(t, tc.expectedBody.Phone, responseBody.Phone)
				assert.WithinDuration(t, tc.expectedBody.CreatedAt, responseBody.CreatedAt, time.Second)
				assert.WithinDuration(t, tc.expectedBody.UpdatedAt, responseBody.UpdatedAt, time.Second)
			}
		})
	}
}

func TestUserController_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	userController := controller.NewUserController(mockUserService)

	testCases := []struct {
		name           string
		userID         string
		mockSetup      func()
		expectedStatus int
		expectedBody   gin.H
	}{
		{
			name:   "Success - User deleted successfully",
			userID: "1",
			mockSetup: func() {
				mockUserService.EXPECT().DeleteUser(1).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: gin.H{
				"message": "user deleted successfully",
			},
		},
		{
			name:   "Failure - User not found",
			userID: "999",
			mockSetup: func() {
				mockUserService.EXPECT().DeleteUser(999).Return(errors.New("user not found"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: gin.H{
				"error": "user not found",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			req, err := http.NewRequest(http.MethodDelete, "/users/"+tc.userID, nil)
			assert.NoError(t, err)

			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)
			ctx.Request = req
			ctx.Params = gin.Params{gin.Param{Key: "id", Value: tc.userID}}

			userController.DeleteUser(ctx)

			assert.Equal(t, tc.expectedStatus, res.Code)

			var responseBody gin.H
			err = json.Unmarshal(res.Body.Bytes(), &responseBody)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedBody, responseBody)
		})
	}
}
