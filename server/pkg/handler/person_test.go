package handler

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	server "github.com/deuuus/bmstu-rsoi"
	"github.com/deuuus/bmstu-rsoi/pkg/service"
	mock_service "github.com/deuuus/bmstu-rsoi/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
)

func TestHandler_getAllPersons(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	logrus.SetOutput(ioutil.Discard)

	type mockBehaviour func(r *mock_service.MockPerson, outputPersons []server.Person)

	tests := []struct {
		name                 string
		outputPersons        []server.Person
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			outputPersons: []server.Person{
				{
					Name:    "Test Name",
					Age:     22,
					Address: "Moscow",
					Work:    "Sleep",
				},
				{
					Name:    "Test Name2",
					Age:     23,
					Address: "Moscow",
					Work:    "Eat",
				},
				{
					Name:    "Test Name3",
					Age:     24,
					Address: "Moscow",
					Work:    "Work",
				},
			},
			mockBehaviour: func(r *mock_service.MockPerson, outputPersons []server.Person) {
				r.EXPECT().GetAllPersons().Return(outputPersons, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `[{"id":0,"name":"Test Name","age":22,"address":"Moscow","work":"Sleep"},{"id":0,"name":"Test Name2","age":23,"address":"Moscow","work":"Eat"},{"id":0,"name":"Test Name3","age":24,"address":"Moscow","work":"Work"}]`,
		},
		{
			name: "Internal Server Error",
			outputPersons: []server.Person{
				{
					Name:    "Test Name",
					Age:     22,
					Address: "Moscow",
					Work:    "Sleep",
				},
			},
			mockBehaviour: func(r *mock_service.MockPerson, outputPersons []server.Person) {
				r.EXPECT().GetAllPersons().Return(outputPersons, errors.New("some internal error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"some internal error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockPerson(c)
			test.mockBehaviour(repo, test.outputPersons)

			services := &service.Service{Person: repo}
			handler := Handler{services}

			r := gin.New()
			r.GET("/api/v1/persons", handler.getAllPersons)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/persons"), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_getPersonById(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	logrus.SetOutput(ioutil.Discard)

	type mockBehaviour func(r *mock_service.MockPerson, id int, outputPerson server.Person)

	tests := []struct {
		name                 string
		id                   int
		outputPerson         server.Person
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			id:   1,
			outputPerson: server.Person{
				Name:    "Test Name",
				Age:     22,
				Address: "Moscow",
				Work:    "Sleep",
			},
			mockBehaviour: func(r *mock_service.MockPerson, id int, outputPerson server.Person) {
				r.EXPECT().GetPersonById(gomock.Eq(id)).Return(outputPerson, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"id":0,"name":"Test Name","age":22,"address":"Moscow","work":"Sleep"}`,
		},
		{
			name: "Not Found",
			id:   100,
			outputPerson: server.Person{
				Name:    "Test Name",
				Age:     22,
				Address: "Moscow",
				Work:    "Sleep",
			},
			mockBehaviour: func(r *mock_service.MockPerson, id int, outputPerson server.Person) {
				r.EXPECT().GetPersonById(gomock.Eq(id)).Return(outputPerson, errors.New("not found"))
			},
			expectedStatusCode:   http.StatusNotFound,
			expectedResponseBody: `{"message":"not found"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockPerson(c)
			test.mockBehaviour(repo, test.id, test.outputPerson)

			services := &service.Service{Person: repo}
			handler := Handler{services}

			r := gin.New()
			r.GET("/api/v1/persons/:id", handler.getPerson)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/persons/%d", test.id), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_deletePerson(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	logrus.SetOutput(ioutil.Discard)

	type mockBehaviour func(r *mock_service.MockPerson, id int)

	tests := []struct {
		name                 string
		id                   int
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			id:   1,
			mockBehaviour: func(r *mock_service.MockPerson, id int) {
				r.EXPECT().DeletePersonById(gomock.Eq(id)).Return(nil)
			},
			expectedStatusCode: http.StatusNoContent,
			//expectedResponseBody: `""`,
		},
		{
			name: "Not found",
			id:   100,
			mockBehaviour: func(r *mock_service.MockPerson, id int) {
				r.EXPECT().DeletePersonById(gomock.Eq(id)).Return(errors.New("not found"))
			},
			expectedStatusCode: http.StatusInternalServerError,
			//expectedResponseBody: `""`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockPerson(c)
			test.mockBehaviour(repo, test.id)

			services := &service.Service{Person: repo}
			handler := Handler{services}

			r := gin.New()
			r.DELETE("/api/v1/persons/:id", handler.deletePerson)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/v1/persons/%d", test.id), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			//assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_createPerson(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	logrus.SetOutput(ioutil.Discard)

	type mockBehaviour func(r *mock_service.MockPerson, person server.Person)

	tests := []struct {
		name                 string
		inputBody            string
		inputPerson          server.Person
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"name": "Test Name", "age": 22, "address": "Moscow", "work": "Sleep"}`,
			inputPerson: server.Person{
				Name:    "Test Name",
				Age:     22,
				Address: "Moscow",
				Work:    "Sleep",
			},
			mockBehaviour: func(r *mock_service.MockPerson, person server.Person) {
				r.EXPECT().CreatePerson(person).Return(1, nil)
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `""`,
		},
		{
			name:                 "Bad request",
			inputBody:            `{"name": "Test Name", "work": "Sleep"}`,
			inputPerson:          server.Person{},
			mockBehaviour:        func(r *mock_service.MockPerson, person server.Person) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Key: 'Person.Age' Error:Field validation for 'Age' failed on the 'required' tag\nKey: 'Person.Address' Error:Field validation for 'Address' failed on the 'required' tag"}`,
		},
		{
			name:      "Internal Server Error",
			inputBody: `{"name": "Test Name", "age": 22, "address": "Moscow", "work": "Sleep"}`,
			inputPerson: server.Person{
				Name:    "Test Name",
				Age:     22,
				Address: "Moscow",
				Work:    "Sleep",
			},
			mockBehaviour: func(r *mock_service.MockPerson, person server.Person) {
				r.EXPECT().CreatePerson(person).Return(0, errors.New("some internal error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"some internal error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockPerson(c)
			test.mockBehaviour(repo, test.inputPerson)

			services := &service.Service{Person: repo}
			handler := Handler{services}

			r := gin.New()
			r.POST("/api/v1/persons", handler.createPerson)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/persons",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_updatePerson(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	logrus.SetOutput(ioutil.Discard)

	type mockBehaviour func(r *mock_service.MockPerson, id int, inputPerson server.PersonUpdate, outputPerson server.Person)

	tests := []struct {
		name                 string
		id                   int
		inputBody  string
		inputPerson          server.PersonUpdate
		outputPerson         server.Person
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			id:   1,
			inputBody: `{"name":"","age": 3,"work":"","address":""}`,
			inputPerson: server.PersonUpdate{
				Name:    "",
				Age:     3,
				Address: "",
				Work:    "",
			},
			outputPerson: server.Person{
				Name:    "Test Name",
				Age:     3,
				Address: "Moscow",
				Work:    "Sleep",
			},
			mockBehaviour: func(r *mock_service.MockPerson, id int, inputPerson server.PersonUpdate, outputPerson server.Person) {
				r.EXPECT().UpdatePerson(gomock.Eq(id), inputPerson).Return(outputPerson, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponseBody: `{"id":0,"name":"Test Name","age":3,"address":"Moscow","work":"Sleep"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockPerson(c)
			test.mockBehaviour(repo, test.id, test.inputPerson, test.outputPerson)

			services := &service.Service{Person: repo}
			handler := Handler{services}

			r := gin.New()
			r.PATCH("/api/v1/persons/:id", handler.updatePerson)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PATCH", fmt.Sprintf("/api/v1/persons/%d", test.id),
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
