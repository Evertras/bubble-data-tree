package datatree

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInline(t *testing.T) {
	tests := []struct {
		name     string
		data     interface{}
		expected bool
	}{
		{
			name:     "String",
			data:     "Hello",
			expected: true,
		},
		{
			name:     "Int",
			data:     3,
			expected: true,
		},
		{
			name:     "Struct",
			data:     struct{ size int }{3},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := isInline(reflect.ValueOf(test.data))

			assert.Equal(t, test.expected, actual)
		})
	}
}
