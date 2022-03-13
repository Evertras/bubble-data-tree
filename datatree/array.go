package datatree

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// This is admittedly getting a bit long, should split this out later...
// nolint: funlen
func (m Model) renderDataNodeArray(data reflect.Value, renderCtx renderContext) string {
	result := strings.Builder{}

	elemType := data.Type().Elem()

	for elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}

	// Eventually we may want to be explicit about everything, but for now just
	// use the default shortcut
	// nolint: exhaustive
	switch elemType.Kind() {
	case reflect.Struct, reflect.Array, reflect.Slice:
		result.WriteString("\n")
		const borderChars = 2
		const padding = 1

		marginLeft := renderCtx.indentLevel*m.indentSize - borderChars

		// TODO: Figure out nested arrays being tighter
		innerWidth := marginLeft - borderChars

		style := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			MarginLeft(marginLeft).
			PaddingLeft(padding).
			PaddingRight(padding).
			MaxWidth(innerWidth)

		const borderWidth = 4

		nestedCtx := renderContext{
			keyName:          renderCtx.keyName,
			indentLevel:      0,
			extraMarginWidth: renderCtx.extraMarginWidth + borderWidth,
		}

		for i := 0; i < data.Len(); i++ {
			entryStr := m.renderDataNode(data.Index(i), nestedCtx)
			entryStr = strings.TrimSpace(entryStr)

			result.WriteString(style.Render(entryStr))
			result.WriteString("\n")
		}

	case reflect.String:
		result.WriteString("\n")
		marginLeft := renderCtx.indentLevel * m.indentSize

		style := lipgloss.NewStyle().MarginLeft(marginLeft)

		for i := 0; i < data.Len(); i++ {
			entryStr := "- " + data.Index(i).String()

			result.WriteString(style.Render(entryStr))
			result.WriteString("\n")
		}

	default:
		result.WriteString(fmt.Sprintf("%v", data))
	}

	return result.String()
}
