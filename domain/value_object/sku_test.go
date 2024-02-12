package valueobject_test

import (
	"strings"
	"testing"
	"time"

	valueobject "item-service/domain/value_object"

	"github.com/stretchr/testify/assert"
)

func TestNewSKU(t *testing.T) {
	createdAt := time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC)
	name := "Example Product Name"
	expectedInitials := []string{"E", "P", "N"}
	expectedDatePart := "220115"
	expectedSKUValue := strings.ToUpper(strings.Join(expectedInitials, "")) + "-" + expectedDatePart

	sku := valueobject.NewSKU(name, createdAt)

	assert.Equal(t, expectedSKUValue, sku.GetValue(), "SKU value should be created correctly")
}

func TestSKUEquals(t *testing.T) {
	createdAt := time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC)
	name := "Example Product Name"
	expectedInitials := []string{"E", "P", "N"}
	expectedDatePart := "220115"
	expectedSKUValue := strings.ToUpper(strings.Join(expectedInitials, "")) + "-" + expectedDatePart

	sku := valueobject.NewSKU(name, createdAt)

	assert.True(t, sku.Equals(expectedSKUValue), "SKU should same")
}
