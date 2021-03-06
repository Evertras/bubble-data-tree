package datatree

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
	"github.com/muesli/reflow/wordwrap"
)

type renderContext struct {
	keyName          string
	indentLevel      int
	extraMarginWidth int
}

func (m *Model) updateContents() {
	if m.data == nil {
		m.contents = m.strNil

		return
	}

	reflected := reflect.ValueOf(m.data)

	m.contents = strings.TrimSpace(m.renderDataNode(reflected, renderContext{}))
}

func (m *Model) renderDataNode(data reflect.Value, renderCtx renderContext) string {
	for data.Kind() == reflect.Ptr {
		if data.IsNil() {
			return m.strNil
		}

		data = data.Elem()
	}

	var result string

	// Eventually we probably want to be explicit about everything, but for now
	// we'll take the default shortcut
	// nolint: exhaustive
	switch data.Kind() {
	case reflect.Struct:
		result = m.renderDataNodeStruct(data, renderCtx)

	case reflect.Map:
		result = m.renderDataNodeMap(data, renderCtx)

	case reflect.Array, reflect.Slice:
		result = m.renderDataNodeArray(data, renderCtx)

	default:
		result = fmt.Sprintf("%v", data)

		const keySuffixLength = 2

		baseIndentWidth := renderCtx.indentLevel * m.indentSize
		keyWidth := (runewidth.StringWidth(renderCtx.keyName) + keySuffixLength)
		remainingWidth := m.width - baseIndentWidth - keyWidth - renderCtx.extraMarginWidth

		hasNewlines := strings.ContainsAny(result, "\n")
		isTooLong := m.width > 0 && runewidth.StringWidth(result) > remainingWidth

		if hasNewlines || isTooLong {
			nextIndentWith := (renderCtx.indentLevel + 1) * m.indentSize

			marginIndent := lipgloss.NewStyle().MarginLeft(m.indentSize)

			// Add one because this checks for <, not <=
			wrapped := wordwrap.String(result, m.width-nextIndentWith-renderCtx.extraMarginWidth+1)

			result = "\n" + marginIndent.Render(wrapped)
		}
	}

	return trimNewline(result)
}
