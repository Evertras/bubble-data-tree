package datatree

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	defaultIndentSize = 2
)

var (
	styleFieldKey = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#224",
		Dark:  "#b8e",
	}).Bold(true).MarginRight(1)
)

type Model struct {
	data       interface{}
	indentSize int
	showZero   bool

	contents string
}

func New(data interface{}) Model {
	model := Model{
		data:       data,
		indentSize: defaultIndentSize,
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
