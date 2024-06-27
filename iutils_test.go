package iutils

import (
	"strconv"
	"strings"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	testCases := []struct {
		length int
	}{
		{0},
		{1},
		{5},
		{10},
		{15},
		{20},
		{25},
		{30},
		{50},
		{100},
	}

	for _, tc := range testCases {
		t.Run("Length_"+strconv.Itoa(tc.length), func(t *testing.T) {
			generated := GenerateRandomString(tc.length)
			if len(generated) != tc.length {
				t.Errorf("Expected string of length %d, got %d", tc.length, len(generated))
			}
			for _, c := range generated {
				if !strings.Contains("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", string(c)) {
					t.Errorf("Character %q not in allowed set", c)
				}
			}
		})
	}
}
