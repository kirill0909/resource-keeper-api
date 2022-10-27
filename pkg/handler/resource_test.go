package handler

import (
	"bytes"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/kirill0909/resource-keeper-api/models"
	"github.com/kirill0909/resource-keeper-api/pkg/service"
	service_mocks "github.com/kirill0909/resource-keeper-api/pkg/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHandler_createResource(t *testing.T) {
	type mockBehavior func(r *service_mocks.MockUserResource, resource models.UserResource)

	testTable := []struct {
		name                 string
		inputBody            string
		inputResource        models.UserResource
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:          "Ok",
			inputBody:     `{"user_id":1, "resource_name":"rname", "resource_login":"rlogin", "resource_password":"rpass"}`,
			inputResource: models.UserResource{UID: 1, ResourceName: "rname", ResourceLogin: "rlogin", ResourcePassword: "rpass"},
			mockBehavior: func(r *service_mocks.MockUserResource, resource models.UserResource) {
				r.EXPECT().CreateResource(resource).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Invalid input body",
			inputBody:            `{"user_id":1, "resource_login":"rlogin", "resource_password":"rpass"}`,
			inputResource:        models.UserResource{UID: 1, ResourceName: "rname", ResourceLogin: "rlogin", ResourcePassword: "rpass"},
			mockBehavior:         func(r *service_mocks.MockUserResource, resource models.UserResource) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:                 "Invalid input value",
			inputBody:            `{"user_id":1, "resource_name":"   ", "resource_login":"rlogin", "resource_password":" "}`,
			inputResource:        models.UserResource{UID: 1, ResourceName: "rname", ResourceLogin: "rlogin", ResourcePassword: "rpass"},
			mockBehavior:         func(r *service_mocks.MockUserResource, resource models.UserResource) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input value"}`,
		},
		{
			name:          "Something went wrong",
			inputBody:     `{"user_id":1, "resource_name":"rname", "resource_login":"rlogin", "resource_password":"rpass"}`,
			inputResource: models.UserResource{UID: 1, ResourceName: "rname", ResourceLogin: "rlogin", ResourcePassword: "rpass"},
			mockBehavior: func(r *service_mocks.MockUserResource, resource models.UserResource) {
				r.EXPECT().CreateResource(resource).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// init Dependencies
			controller := gomock.NewController(t)
			defer controller.Finish()

			repo := service_mocks.NewMockUserResource(controller)
			testCase.mockBehavior(repo, testCase.inputResource)

			service := &service.Service{UserResource: repo}
			handler := Handler{service}

			// Init endpoint
			router := gin.New()
			router.POST("/resource", handler.createResource)

			// Create request
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/resource",
				bytes.NewBufferString(testCase.inputBody))

			testContext, _ := gin.CreateTestContext(recorder)
			testContext.Request = request
			testContext.Set(userCtx, 1)

			// Perform request
			handler.createResource(testContext)

			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
			assert.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}

}

func TestHandler_getAllResources(t *testing.T) {
	type mockBehavior func(r *service_mocks.MockUserResource, userId int)

	testTable := []struct {
		name                 string
		UID                  int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			UID:  1,
			mockBehavior: func(r *service_mocks.MockUserResource, userId int) {
				r.EXPECT().GetAllResources(userId).Return([]models.UserResource{
					{Id: 1, UID: 1, ResourceName: "rname", ResourceLogin: "rlogin", ResourcePassword: "rpass",
						DateCreation: "2022-10-27 10:40:21.123", LastUpdate: "2022-10-27 10:40:21.123"},
				}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{"data":[{"id":1,"user_id":1,"resource_name":"rname","resource_login":"rlogin",` +
				`"resource_password":"rpass","date_creation":"2022-10-27 10:40:21.123","last_update":"2022-10-27 10:40:21.123"}]}`,
		},
		{
			name: "User id not found",
			mockBehavior: func(r *service_mocks.MockUserResource, userId int) {
				r.EXPECT().GetAllResources(userId).Return(nil, errors.New("user id not found"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"user id not found"}`,
		},
		{
			name: "Service faild",
			UID:  1,
			mockBehavior: func(r *service_mocks.MockUserResource, userId int) {
				r.EXPECT().GetAllResources(userId).Return(nil, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			controller.Finish()

			repo := service_mocks.NewMockUserResource(controller)
			testCase.mockBehavior(repo, testCase.UID)

			service := &service.Service{UserResource: repo}
			handler := Handler{service}

			router := gin.New()
			router.GET("/resource", handler.getAllResources)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/resource", nil)

			testContext, _ := gin.CreateTestContext(recorder)
			testContext.Request = request
			testContext.Set(userCtx, testCase.UID)

			handler.getAllResources(testContext)

			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
			assert.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}

func TestHandler_getResourceById(t *testing.T) {
	type mockBehavior func(r *service_mocks.MockUserResource, userId, rsourceid int)

	testTable := []struct {
		name                 string
		UID                  int
		resourceId           int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:       "Ok",
			UID:        1,
			resourceId: 1,
			mockBehavior: func(r *service_mocks.MockUserResource, userId, resourceId int) {
				r.EXPECT().GetById(userId, resourceId).Return(models.UserResource{
					Id: 1, UID: 1, ResourceName: "rname", ResourceLogin: "rlogin", ResourcePassword: "rpass",
					DateCreation: "2022-10-27 10:40:21.123", LastUpdate: "2022-10-27 10:40:21.123"}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{"resource":{"id":1,"user_id":1,"resource_name":"rname","resource_login":"rlogin",` +
				`"resource_password":"rpass","date_creation":"2022-10-27 10:40:21.123","last_update":"2022-10-27 10:40:21.123"}}`,
		},
		{
			name: "Invalid id param",
			mockBehavior: func(r *service_mocks.MockUserResource, userId, resourceId int) {
				r.EXPECT().GetById(userId, resourceId).Return(models.UserResource{}, errors.New("invalid id param"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"invalid id param"}`,
		},
		{
			name:       "Service faild",
			UID:        1,
			resourceId: 1,
			mockBehavior: func(r *service_mocks.MockUserResource, userId, resourceId int) {
				r.EXPECT().GetById(userId, resourceId).Return(models.UserResource{}, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			controller.Finish()

			repo := service_mocks.NewMockUserResource(controller)
			testCase.mockBehavior(repo, testCase.UID, testCase.resourceId)

			service := &service.Service{UserResource: repo}
			handler := Handler{service}

			router := gin.New()
			router.GET("/resource", handler.getResourceById)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/resource", nil)

			testContext, _ := gin.CreateTestContext(recorder)
			testContext.Request = request
			testContext.Set(userCtx, testCase.UID)
			testContext.AddParam("id", fmt.Sprintf("%d", testCase.resourceId))

			handler.getResourceById(testContext)

			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
			assert.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}

func TestHandler_updateResource(t *testing.T) {
	type mockBehavior func(r *service_mocks.MockUserResource, userId, rsourceid int,
		input models.UserResourceUpdate)

	args := struct {
		ResName  string
		ResLogin string
		ResPass  string
	}{
		ResName:  "rname",
		ResLogin: "rlogin",
		ResPass:  "rpass",
	}

	testTable := []struct {
		name                 string
		UID                  int
		resourceId           int
		inputBody            string
		inputResource        models.UserResourceUpdate
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:       "Ok",
			UID:        1,
			resourceId: 1,
			inputBody:  `{"resource_name":"rname", "resource_login":"rlogin", "resource_password":"rpass"}`,
			inputResource: models.UserResourceUpdate{
				ResourceName:     &args.ResName,
				ResourceLogin:    &args.ResLogin,
				ResourcePassword: &args.ResPass,
			},
			mockBehavior: func(r *service_mocks.MockUserResource, userId, resourceId int, input models.UserResourceUpdate) {
				r.EXPECT().UpdateResource(userId, resourceId, input).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"status":"ok"}`,
		},
		{
			name:      "Invalid id param",
			inputBody: `{"resource_name":"rname", "resource_login":"rlogin", "resource_password":"rpass"}`,
			inputResource: models.UserResourceUpdate{
				ResourceName:     &args.ResName,
				ResourceLogin:    &args.ResLogin,
				ResourcePassword: &args.ResPass,
			},
			mockBehavior: func(r *service_mocks.MockUserResource, userId, resourceId int, input models.UserResourceUpdate) {
				r.EXPECT().UpdateResource(userId, resourceId, input).Return(errors.New("invalid id param"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"invalid id param"}`,
		},
		{
			name:      "Empty value",
			inputBody: `{"resource_name":" ", "resource_login":"rlogin", "resource_password":"rpass"}`,
			inputResource: models.UserResourceUpdate{
				ResourceName:     pointerString(" "),
				ResourceLogin:    &args.ResLogin,
				ResourcePassword: &args.ResPass,
			},
			mockBehavior: func(r *service_mocks.MockUserResource, userId, resourceId int, input models.UserResourceUpdate) {
				r.EXPECT().UpdateResource(userId, resourceId, input).Return(errors.New("update structure has no values"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"update structure has no values"}`,
		},
		{
			name:      "All values is empty",
			inputBody: `{"resource_name":" ", "resource_login":" ", "resource_password":" "}`,
			inputResource: models.UserResourceUpdate{
				ResourceName:     pointerString(" "),
				ResourceLogin:    pointerString(" "),
				ResourcePassword: pointerString(" "),
			},
			mockBehavior: func(r *service_mocks.MockUserResource, userId, resourceId int, input models.UserResourceUpdate) {
				r.EXPECT().UpdateResource(userId, resourceId, input).Return(errors.New("update structure has no values"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"update structure has no values"}`,
		},
		{
			name:          "Empty body",
			inputBody:     `{}`,
			inputResource: models.UserResourceUpdate{},
			mockBehavior: func(r *service_mocks.MockUserResource, userId, resourceId int, input models.UserResourceUpdate) {
				r.EXPECT().UpdateResource(userId, resourceId, input).Return(errors.New("no new value for set"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"no new value for set"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			controller.Finish()

			repo := service_mocks.NewMockUserResource(controller)
			testCase.mockBehavior(repo, testCase.UID, testCase.resourceId, testCase.inputResource)

			service := &service.Service{UserResource: repo}
			handler := Handler{service}

			router := gin.New()
			router.PUT("/resource", handler.updateResource)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("PUT", "/resource",
				bytes.NewBufferString(testCase.inputBody))

			testContext, _ := gin.CreateTestContext(recorder)
			testContext.Request = request
			testContext.Set(userCtx, testCase.UID)
			testContext.AddParam("id", fmt.Sprintf("%d", testCase.resourceId))

			handler.updateResource(testContext)

			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
			assert.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}

func pointerString(str string) *string {
	return &str
}
