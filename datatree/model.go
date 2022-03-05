package datatree

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	defaultIndentSize = 2
)

type Model struct {
	data       interface{}
	indentSize int
	showZero   bool
	styles     Styles

	contents string

	strNil string
}

func New(data interface{}) Model {
	model := Model{
		data:       data,
		indentSize: defaultIndentSize,
		styles:     styleDefault,
		strNil:     "<nil>",
	}

	model.updateContents()

	return model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	body := strings.Builder{}

	body.WriteString(m.contents)

	return body.String()
}
