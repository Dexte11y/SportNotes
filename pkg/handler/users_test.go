package handler

import (
	"bytes"
	"net/http/httptest"
	"sportnotes/pkg/schemas"
	"sportnotes/pkg/service"
	"testing"

	mockService "sportnotes/pkg/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestCreateUserHandler(t *testing.T) {
	type mockBehavior func(s *mockService.MockUserList, user schemas.User)
	testTable := []struct {
		name                string
		inputBody           string
		inputUser           schemas.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test", "surname":"test", "login":"testLogin", "password":"qwerty", "email":"testEmail"}`,
			inputUser: schemas.User{
				Name:     "Test",
				Surname:  "test",
				Login:    "testLogin",
				Password: "qwerty",
				Email:    "testEmail",
			},
			mockBehavior: func(s *mockService.MockUserList, user schemas.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		// {
		// 	name:                "Empty Fields",
		// 	inputBody:           `{"name":"Test", "surname":"test", "login":"testLogin", "password":"qwerty", "email":"testEmail"}`,
		// 	mockBehavior:        func(s *mockService.MockUserList, user schemas.User) {},
		// 	expectedStatusCode:  400,
		// 	expectedRequestBody: `{"message":"invalid input body"}`,
		// },
		// {
		// 	name:      "Service Failure",
		// 	inputBody: `{"name":"Test", "surname":"test", "login":"testLogin", "password":"qwerty",  "email":"testEmail"}`,
		// 	inputUser: schemas.User{
		// 		Name:     "Test",
		// 		Surname:  "test",
		// 		Login:    "testLogin",
		// 		Password: "qwerty",
		// 		Email:    "Email",
		// 	},
		// 	mockBehavior: func(s *mockService.MockUserList, user schemas.User) {
		// 		s.EXPECT().CreateUser(user).Return(0, errors.New("service failure"))
		// 	},
		// 	expectedStatusCode:  500,
		// 	expectedRequestBody: `{"message":"service failure"}`,
		// },
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			user := mockService.NewMockUserList(c)
			testCase.mockBehavior(user, testCase.inputUser)

			services := &service.Service{UserList: user}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/api/users/", handler.createUser)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/users/", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
