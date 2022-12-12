package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

}

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("database error"))

	assert.NotNil(t, err)

	status := err.Status()
	assert.EqualValues(t, http.StatusInternalServerError, status)

	message := err.Message()
	assert.EqualValues(t, "this is the message", message)

	errVal := err.Error()
	assert.EqualValues(t, "message: this is the message - status: 500 - error: internal_server_error - causes: [database error]", errVal)

	causes := err.Causes()
	assert.NotNil(t, causes)
	assert.EqualValues(t, 1, len(causes))
	assert.EqualValues(t, "database error", causes[0])
}

func TestNewBadRequestError(t *testing.T) {
}

func TestNewNotFoundError(t *testing.T) {
}

func TestNewError(t *testing.T) {

}
