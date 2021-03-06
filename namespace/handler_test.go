package namespace

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestCreateNamespaceHandlerShouldReturnErrorWhenBodyCannotBeParsed(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	payload := `
		{
			adasdas
		}
	`
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/org/test-service/namespaces", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/namespaces", CreateNamespaceHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateNamespaceHandlerShouldReturnErrorWhenProcessorReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	payload := `
		{
			"namespace": "default"
		}
	`
	mock.EXPECT().createNewNamespaceProcessor("org", "test-service", &namespaceView{
		Namespace: "default",
	}).Return(errors.New("Test Error"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/org/test-service/namespaces", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/namespaces", CreateNamespaceHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestCreateNamespaceHandlerShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	payload := `
		{
			"namespace": "default"
		}
	`
	mock.EXPECT().createNewNamespaceProcessor("org", "test-service", &namespaceView{
		Namespace: "default",
	}).Return(nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/org/test-service/namespaces", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/namespaces", CreateNamespaceHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestRetrieveAllNamespaceHandlerShouldReturn500WhenError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	mock.EXPECT().retrieveAllNamespaceProcessor("org", "test-service").Return([]byte(""), errors.New("TestError"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/org/test-service/namespaces", nil)
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/namespaces", RetrieveAllNamespaceHandler).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
func TestRetrieveAllNamespaceHandlerShouldReturn200(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	mock.EXPECT().retrieveAllNamespaceProcessor("org", "test-service").Return([]byte("[]"), nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/org/test-service/namespaces", nil)
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/namespaces", RetrieveAllNamespaceHandler).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
