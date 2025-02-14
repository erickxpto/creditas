package utils_test

import (
	"creditas/pkg/utils"
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	tests := []struct {
		birthday string
		expected int
	}{
		{"2000-01-01", 25},
		{"1990-06-15", 34},
		{"1980-12-31", 44},
	}

	for _, test := range tests {
		birthday, _ := time.Parse("2006-01-02", test.birthday)
		result := utils.CalculateAge(birthday)
		if result != test.expected {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}
