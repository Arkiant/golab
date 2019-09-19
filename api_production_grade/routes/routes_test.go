package routes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	assert.HTTPSuccess(t, Health().ServeHTTP, "GET", "/health", nil)
	assert.HTTPBodyContains(t, Health().ServeHTTP, "GET", "/health", nil, "OK")
}
