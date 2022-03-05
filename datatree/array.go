package datatree

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

func (m Model) renderDataNodeArray(data reflect.Value, indentLevel int) string {
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

		marginLeft := indentLevel*m.indentSize - borderChars

		// TODO: Figure out nested arrays being tighter
		innerWidth := marginLeft - borderChars

		style := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			MarginLeft(marginLeft).
			PaddingLeft(padding).
			PaddingRight(padding).
			MaxWidth(innerWidth)

		for i := 0; i < data.Len(); i++ {
			entryStr := m.renderDataNode(data.Index(i), 0)
			entryStr = strings.TrimSpace(entryStr)
			entryStr = wordwrap.String(entryStr, innerWidth)

			result.WriteString(style.Render(entryStr))
			result.WriteString("\n")
		}

	default:
		result.WriteString(fmt.Sprintf("%v", data))
	}

	return result.String()
}
