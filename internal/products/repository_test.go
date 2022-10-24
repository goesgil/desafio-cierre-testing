package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBySeller(t *testing.T) {
	//Assign
	sellerID := "FEX112AC"
	newRepo := NewRepository()

	//Act
	prodList, err := newRepo.GetAllBySeller(sellerID)

	// Assert
	assert.Nil(t, err)
	for _, prod := range prodList {
		assert.Equal(t, sellerID, prod.SellerID)
	}
}
