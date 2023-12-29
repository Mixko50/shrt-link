package utility_test

import (
	"github.com/stretchr/testify/assert"
	"shrt-server/utility"
	"shrt-server/utility/text"
	"testing"
)

func TestUtility_GenerateSlug(t *testing.T) {
	// This example shows how to use GenerateSlug function.
	// Create a new utility.
	utility := utility.NewUtility()

	// Generate a new slug.
	slug := utility.GenerateSlug()

	// Print the slug.
	assert.Equal(t, 6, len(slug))
	assert.True(t, text.IsAlphanumeric(slug))
}
