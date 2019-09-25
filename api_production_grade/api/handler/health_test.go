package handler

import (
	"testing"

	"github.com/arkiant/golab/api_production_grade/api/service"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {

	s := service.NewService(nil)
	h := NewHandler(s)

	assert.HTTPSuccess(t, h.ServeHTTP, "GET", "/v1/api/status/health", nil)
	assert.HTTPBodyContains(t, h.ServeHTTP, "GET", "/v1/api/status/health", nil, "{\"message\":\"OK\"}")
}
