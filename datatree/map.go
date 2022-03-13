package datatree

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func (m *Model) renderDataNodeMap(data reflect.Value, renderCtx renderContext) string {
	result := strings.Builder{}
	indent := strings.Repeat(" ", renderCtx.indentLevel*m.indentSize)

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

		nextCtx := renderContext{
			indentLevel:      renderCtx.indentLevel + 1,
			keyName:          entry.key,
			extraMarginWidth: renderCtx.extraMarginWidth,
		}

		renderedData := m.renderDataNode(entry.val, nextCtx)

		if len(renderedData) == 0 {
			continue
		}

		if renderedData[0] != '\n' {
			result.WriteString(" ")
		}

		result.WriteString(renderedData)
	}

	return trimNewline(result.String())
}
