package datatree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViewBasic(t *testing.T) {
	tests := []struct {
		name     string
		data     interface{}
		expected string
	}{
		{
			name:     "String",
			data:     "Hello data tree",
			expected: "Hello data tree",
		},
		{
			name:     "Integer",
			data:     17,
			expected: "17",
		},
		{
			name: "Single Field Struct",
			data: struct {
				Name string
			}{"Ralph"},
			expected: "Name: Ralph",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			model := New(test.data).WithStyleBlank()

			rendered := model.View()

			assert.Equal(t, test.expected, rendered)
		})
	}
}
