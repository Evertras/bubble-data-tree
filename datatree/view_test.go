package datatree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// This is a long function but only because there's a lot of repetitive test
// cases, not because the logic is complicated
// nolint: funlen
func TestViewDefaultBlank(t *testing.T) {
	var emptyPtr *int
	intVal := 3

	tests := []struct {
		name     string
		data     interface{}
		expected string
	}{
		{
			name:     "Nil",
			data:     nil,
			expected: "<nil>",
		},
		{
			name:     "NilIntPointer",
			data:     emptyPtr,
			expected: "<nil>",
		},
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
			name:     "Integer Pointer",
			data:     &intVal,
			expected: "3",
		},
		{
			name: "Single Field Struct",
			data: struct {
				Name string
			}{"Ralph"},
			expected: "Name: Ralph",
		},
		{
			name: "Multiple Field Flat Struct",
			data: struct {
				Part  string
				Count int
			}{"Button", 3},
			expected: "Count: 3\nPart: Button",
		},
		{
			name: "Multiple Field Flat Struct With Unexported",
			data: struct {
				Part   string
				Count  int
				sneaky int
			}{"Button", 3, 3},
			expected: "Count: 3\nPart: Button",
		},
		{
			name: "Struct Zero Value Hidden",
			data: struct {
				Part  string
				Count int
			}{"", 3},
			expected: "Count: 3",
		},
		{
			name: "Struct With Pointer Field",
			data: struct {
				Count *int
			}{&intVal},
			expected: "Count: 3",
		},
		{
			name: "Simple String Map",
			data: map[string]string{
				"abc": "def",
				"123": "ok",
			},
			expected: "123: ok\nabc: def",
		},
		{
			name: "Array Of Integers",
			data: []int{1, 3, 2},
			expected: "[1 3 2]",
		},
		{
			name: "Array Of Strings",
			data: []string{"Hello", "Data"},
			expected: "- Hello\n- Data",
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
