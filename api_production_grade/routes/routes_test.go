package routes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	assert.HTTPSuccess(t, HealthRoute().ServeHTTP, "GET", "/health", nil)
	assert.HTTPBodyContains(t, HealthRoute().ServeHTTP, "GET", "/health", nil, "{\"message\":\"OK\"}")
}
