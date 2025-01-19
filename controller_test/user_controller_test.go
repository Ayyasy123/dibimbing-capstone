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
