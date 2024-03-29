package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDetails(t *testing.T) {
	carDetailsService := NewCarDetailsService()
	carDetails := carDetailsService.GetDetails()

	assert.NotNil(t, carDetails)
	assert.Equal(t, 1, carDetails.ID)
	assert.Equal(t, "Mitsubishi", carDetails.Brand)
	assert.Equal(t, 2002, carDetails.Year)
}
