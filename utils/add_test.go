package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{1, -1, 0},
		{1, 1, 2},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		assert.Equal(t, test.expected, result)
	}
}