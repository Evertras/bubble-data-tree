package datatree

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func (m Model) renderDataNodeMap(data reflect.Value, indentLevel int) string {
	result := strings.Builder{}
	indent := strings.Repeat(" ", indentLevel*m.indentSize)

	iter := data.MapRange()

	keyVals := keyValList{}

	for iter.Next() {
		keyVals = append(keyVals, keyVal{
			key: fmt.Sprintf("%v", iter.Key()),
			val: iter.Value(),
		})
	}

	sort.Sort(keyVals)

	for _, entry := range keyVals {
		result.WriteString("\n")
		keyStr := m.styles.FieldKey.Render(entry.key + ":")

		result.WriteString(indent + keyStr)

		renderedData := m.renderDataNode(entry.val, indentLevel+1)

		if len(renderedData) == 0 || renderedData[0] != '\n' {
			result.WriteString(" ")
		}

		result.WriteString(renderedData)
	}

	return trimNewline(result.String())
}
