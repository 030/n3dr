package count

import (
	"testing"

	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
	"github.com/stretchr/testify/assert"
)

func TestAssets(t *testing.T) {
	asset := models.AssetXO{ContentType: "text/plain"}
	assets := []*models.AssetXO{&asset, &asset, &asset}

	n := Nexus3{}

	repositoriesTotalArtifacts := 0
	repositoriesTotalArtifactsPointer := &repositoriesTotalArtifacts
	n.assets(assets, repositoriesTotalArtifactsPointer)

	assert.Equal(t, 3, *repositoriesTotalArtifactsPointer)
	assert.Equal(t, 3, total)
}

// func TestYourFunction(t *testing.T) {
// 	// Create the Nexus3 mock
// 	mockNexus := new(mocks.Nexus3Interface)

// 	// Set up the expected behavior for the Repos method
// 	mockReposResponse := []*models.AbstractAPIRepository{
// 		// Define your mock response here as needed.
// 	}
// 	mockNexus.On("Repos").Return(mockReposResponse, nil)

// 	_, err := mockNexus.Repos()

// 	// Call your function under test with the mockNexus
// 	// err := YourFunction(mockNexus)

// 	// Assert the results
// 	assert.NoError(t, err)

// 	// Assert that the expected method calls were made
// 	mockNexus.AssertExpectations(t)
// }
