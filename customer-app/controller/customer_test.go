package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"customerapp/domain"
	"customerapp/mocks"
)

var (
	r                 *mux.Router
	handler           CustomerController
	mockCustomerStore *mocks.MockCustomerStore
	mockCtrl          *gomock.Controller
	w                 *httptest.ResponseRecorder
)

const (
	URL string = "/api/customers/"
)

func setup(t *testing.T) {
	r = mux.NewRouter()
	mockCtrl = gomock.NewController(t)
	mockCustomerStore = mocks.NewMockCustomerStore(mockCtrl)
	handler = CustomerController{Store: mockCustomerStore}
	w = httptest.NewRecorder()
}

func TestCustomerController_Create(t *testing.T) {
	setup(t)
	c := getMockCustomer("cust101")
	mockCustomerStore.EXPECT().Create(c).Return(nil).Times(1)
	r.HandleFunc(URL, handler.Create).Methods(http.MethodPost)
	cJson, err := json.Marshal(c)
	if err != nil {
		t.Error(err)
	}
	req, err := http.NewRequest(http.MethodPost, URL, strings.NewReader(string(cJson)))
	if err != nil {
		t.Error(err)
	}
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code, fmt.Sprintf("Http status expected: %d, got: %d", http.StatusCreated, w.Code))
}

func TestCustomerController_Create_Duplicate(t *testing.T) {
	setup(t)
	c := getMockCustomer("cust101")
	mockCustomerStore.EXPECT().Create(c).Return(domain.ErrCustomerIdAlreadyExist).Times(2)
	mockCustomerStore.Create(c)
	r.HandleFunc(URL, handler.Create).Methods(http.MethodPost)
	cJson, err := json.Marshal(c)
	if err != nil {
		t.Error(err)
	}
	req, err := http.NewRequest(http.MethodPost, URL, strings.NewReader(string(cJson)))
	if err != nil {
		t.Error(err)
	}
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code, fmt.Sprintf("Http status expected: %d, got: %d", http.StatusBadRequest, w.Code))
}

func TestCustomerController_Change_CustomerNotFound(t *testing.T) {
	setup(t)
	id := "cust103"
	c := getMockCustomer(id)
	mockCustomerStore.EXPECT().Update(id, c).Return(domain.ErrCustomerNotFound).Times(1)
	r.HandleFunc(URL+"{id}", handler.Change).Methods(http.MethodPatch)
	cJson, err := json.Marshal(c)
	if err != nil {
		t.Error(err)
	}
	req, err := http.NewRequest(http.MethodPatch, URL+id, strings.NewReader(string(cJson)))
	if err != nil {
		t.Error(err)
	}
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code, fmt.Sprintf("Http status expected: %d, got: %d", http.StatusNotFound, w.Code))
}

func TestCustomerController_Change(t *testing.T) {
	setup(t)
	id := "cust101"
	c := getMockCustomer(id)
	mockCustomerStore.EXPECT().Create(c).Return(nil).Times(1)
	mockCustomerStore.EXPECT().Update(id, c).Return(nil).Times(1)
	r.HandleFunc(URL, handler.Create).Methods(http.MethodPost)
	r.HandleFunc(URL+"{id}", handler.Change).Methods(http.MethodPatch)
	cJson, err := json.Marshal(c)
	if err != nil {
		t.Error(err)
	}
	req, err := http.NewRequest(http.MethodPost, URL, strings.NewReader(string(cJson)))
	if err != nil {
		t.Error(err)
	}
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code, fmt.Sprintf("Http status expected: %d, got: %d", http.StatusCreated, w.Code))

	//checking update
	req2, err := http.NewRequest(http.MethodPatch, URL+id, strings.NewReader(string(cJson)))
	if err != nil {
		t.Error(err)
	}
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusNoContent, w2.Code, fmt.Sprintf("Http status expected: %d, got: %d", http.StatusNoContent, w.Code))
}

func getMockCustomer(id string) domain.Customer {
	return domain.Customer{ID: id, Name: "John", Email: "john.doe@fake.com"}
}
