package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"app/middleware"
	"app/models"
)

type MockAuthenticatedUser struct {
	mock.Mock
}

func (m *MockAuthenticatedUser) GetUserBySessionID() *models.UserError {
	args := m.Called()
	return args.Get(0).(*models.UserError)
}

func TestRequireLogin(t *testing.T) {
	mockUser := &MockAuthenticatedUser{}

	req, err := http.NewRequest("GET", "/check", nil)
	assert.NoError(t, err)
	cookie := &http.Cookie{Name: "session", Value: "5d047922-4ea7-4d4e-a596-2c4ed482514b"}
	req.AddCookie(cookie)

	rr := httptest.NewRecorder()
	mockUser.On("GetUserBySessionID").Return(&models.UserError{})

	next := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	handler := middleware.RequireLogin(next)
	handler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
