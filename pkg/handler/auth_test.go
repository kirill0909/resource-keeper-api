package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/kirill0909/resource-keeper-api/models"
	"github.com/kirill0909/resource-keeper-api/pkg/service"
	service_mocks "github.com/kirill0909/resource-keeper-api/pkg/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *service_mocks.MockAuthorization, user models.User)

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            models.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"name":"John Down", "email":"john@gmail.com", "password":"JohnPass"}`,
			inputUser: models.User{
				Name:     "John Down",
				Email:    "john@gmail.com",
				Password: "JohnPass",
			},
			mockBehavior: func(s *service_mocks.MockAuthorization, user models.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Empty Field",
			inputBody:            `{"name":"John Down", "password":"JohnPass"}`,
			mockBehavior:         func(s *service_mocks.MockAuthorization, user models.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:                 "Empty Value",
			inputBody:            `{"name":"John Down", "email":"john@gmail.com", "password":" "}`,
			mockBehavior:         func(s *service_mocks.MockAuthorization, user models.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Failure",
			inputBody: `{"name":"John Down", "email":"john@gmail.com", "password":"JohnPass"}`,
			inputUser: models.User{
				Name:     "John Down",
				Email:    "john@gmail.com",
				Password: "JohnPass",
			},
			mockBehavior: func(s *service_mocks.MockAuthorization, user models.User) {
				s.EXPECT().CreateUser(user).Return(1, errors.New("somethig went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"somethig went wrong"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			auth := service_mocks.NewMockAuthorization(controller)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			router := gin.New()
			router.POST("/sign-up", handler.signUp)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(testCase.inputBody))

			router.ServeHTTP(recorder, request)

			assert.Equal(t, testCase.expectedResponseBody, recorder.Body.String())
			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)

		})
	}
}

func TestHandler_signIn(t *testing.T) {
	type mockBehavior func(s *service_mocks.MockAuthorization, input signInInput)

	testTable := []struct {
		name                 string
		inputBody            string
		inputSignInInput     signInInput
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:             "Ok",
			inputBody:        `{"email":"john@gmail.com", "password":"JohnPass"}`,
			inputSignInInput: signInInput{Email: "john@gmail.com", Password: "JohnPass"},
			mockBehavior: func(s *service_mocks.MockAuthorization, input signInInput) {
				s.EXPECT().GenerateToken(input.Email, input.Password).Return("token", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"token":"token"}`,
		},
		{
			name:                 "Empty Field",
			inputBody:            `{"email":"john@gmail.com"}`,
			inputSignInInput:     signInInput{Email: "john@gmail.com", Password: "JohnPass"},
			mockBehavior:         func(s *service_mocks.MockAuthorization, input signInInput) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:                 "Empty Value",
			inputBody:            `{"email":"john@gmail.com", "password":" "}`,
			inputSignInInput:     signInInput{Email: "john@gmail.com", Password: "John Down"},
			mockBehavior:         func(s *service_mocks.MockAuthorization, input signInInput) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:             "Service Failure",
			inputBody:        `{"email":"john@gmail.com", "password":"JohnPass"}`,
			inputSignInInput: signInInput{Email: "john@gmail.com", Password: "JohnPass"},
			mockBehavior: func(s *service_mocks.MockAuthorization, input signInInput) {
				s.EXPECT().GenerateToken(input.Email, input.Password).Return("", errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			auth := service_mocks.NewMockAuthorization(controller)
			testCase.mockBehavior(auth, testCase.inputSignInInput)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			router := gin.New()
			router.POST("/sign-in", handler.signIn)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/sign-in",
				bytes.NewBufferString(testCase.inputBody))

			router.ServeHTTP(recorder, request)

			assert.Equal(t, testCase.expectedResponseBody, recorder.Body.String())
			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)

		})
	}

}
