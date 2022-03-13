package datatree

import (
	"reflect"
	"sort"
	"strings"
)

func (m *Model) renderDataNodeStruct(data reflect.Value, renderCtx renderContext) string {
	result := strings.Builder{}
	indent := strings.Repeat(" ", renderCtx.indentLevel*m.indentSize)

	fieldNames := []string{}

	result.WriteString("\n")

	for i := 0; i < data.Type().NumField(); i++ {
		field := data.Type().Field(i)

		if !field.IsExported() {
			continue
		}

		fieldNames = append(fieldNames, field.Name)
	}

	sort.Strings(fieldNames)

	for _, fieldName := range fieldNames {
		field := data.FieldByName(fieldName)

		if !m.showZero && field.IsZero() {
			continue
		}

		for field.Kind() == reflect.Ptr && !field.IsNil() {
			field = field.Elem()
		}

		result.WriteString(indent)
		result.WriteString(m.styles.FieldKey.Render(fieldName + ":"))

		nextCtx := renderContext{
			keyName:          fieldName,
			indentLevel:      renderCtx.indentLevel + 1,
			extraMarginWidth: renderCtx.extraMarginWidth,
		}
		renderedData := m.renderDataNode(field, nextCtx)

		if len(renderedData) != 0 && renderedData[0] != '\n' {
			result.WriteString(" ")
		}

		result.WriteString(renderedData)
		result.WriteString("\n")
	}

	return result.String()
}
