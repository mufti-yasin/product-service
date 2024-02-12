package valueobject

import (
	"fmt"
	"strings"
	"time"
)

type SKU struct {
	value string
}

func NewSKU(name string, createdAt time.Time) SKU {
	initials := make([]string, 0)
	words := strings.Fields(name)
	for _, word := range words {
		initials = append(initials, string(word[0]))
	}

	datePart := fmt.Sprintf("%02d%02d%02d", createdAt.Year()%100, int(createdAt.Month()), createdAt.Day())

	skuValue := strings.ToUpper(strings.Join(initials, "")) + "-" + datePart

	return SKU{value: skuValue}
}

func (s SKU) GetValue() string {
	return s.value
}

func (s SKU) Equals(sku string) bool {
	return s.value == sku
}
