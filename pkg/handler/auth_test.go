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
	mocks_service "github.com/kirill0909/resource-keeper-api/pkg/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mocks_service.MockAuthorization, user models.User)

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
			mockBehavior: func(s *mocks_service.MockAuthorization, user models.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Empty Field",
			inputBody:            `{"name":"John Down", "password":"JohnPass"}`,
			mockBehavior:         func(s *mocks_service.MockAuthorization, user models.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:                 "Empty Value",
			inputBody:            `{"name":"John Down", "email":"john@gmail.com", "password":" "}`,
			mockBehavior:         func(s *mocks_service.MockAuthorization, user models.User) {},
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
			mockBehavior: func(s *mocks_service.MockAuthorization, user models.User) {
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

			auth := mocks_service.NewMockAuthorization(controller)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			router := gin.New()
			router.POST("/sign-up", handler.signUp)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))

			router.ServeHTTP(recorder, request)

			assert.Equal(t, testCase.expectedResponseBody, recorder.Body.String())
			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)

		})
	}
}
