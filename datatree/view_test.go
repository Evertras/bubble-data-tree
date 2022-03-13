package datatree

import (
	"strings"
	"testing"

	"github.com/mattn/go-runewidth"
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
		width    int
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
			name:     "Array Of Integers",
			data:     []int{1, 3, 2},
			expected: "[1 3 2]",
		},
		{
			name:     "Array Of Strings",
			data:     []string{"Hello", "Data"},
			expected: "- Hello\n- Data",
		},
		{
			name: "Struct With Map And Array Of Strings",
			data: struct {
				Name    string
				Meta    map[string]string
				Regions []string
			}{
				Name: "",
				Meta: map[string]string{
					"owner": "evertras",
				},
				Regions: []string{"Asia", "North America"},
			},
			expected: `Meta:
  owner: evertras
Regions:
  - Asia
  - North America`,
		},
		{
			name: "Struct string with newlines is nested",
			data: struct {
				Description string
			}{"First\nSecond"},
			expected: `Description:
  First 
  Second`,
		},
		{
			name:  "Long string in struct wraps properly in root",
			width: 80,
			data: struct {
				Name        string
				Description string
			}{
				Name: "Pikachu",
				// This is a very long line, that is the point
				// nolint: lll
				Description: "Pikachu is a fictional species in the Pokémon media franchise. Designed by Atsuko Nishida and Ken Sugimori, Pikachu first appeared in the 1996 Japanese video games Pokémon Red and Green created by Game Freak and Nintendo, which were released outside of Japan in 1998 as Pokémon Red and Blue. Pikachu is a yellow, mouse-like creature with electrical abilities. It is a major character in the Pokémon franchise, serving as its mascot and as a major mascot for Nintendo.",
			},
			expected: `Description:
  Pikachu is a fictional species in the Pokémon media franchise. Designed by    
  Atsuko Nishida and Ken Sugimori, Pikachu first appeared in the 1996 Japanese  
  video games Pokémon Red and Green created by Game Freak and Nintendo, which   
  were released outside of Japan in 1998 as Pokémon Red and Blue. Pikachu is a  
  yellow, mouse-like creature with electrical abilities. It is a major character
  in the Pokémon franchise, serving as its mascot and as a major mascot for     
  Nintendo.                                                                     
Name: Pikachu`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			model := New(test.data).WithStyleBlank().WithWidth(test.width)

			rendered := model.View()

			assert.Len(t, rendered, len(test.expected))

			assert.Equal(t, test.expected, rendered)

			if test.width > 0 {
				lines := strings.Split(rendered, "\n")

				for _, line := range lines {
					assert.LessOrEqual(t, runewidth.StringWidth(line), test.width)
				}
			}
		})
	}
}
