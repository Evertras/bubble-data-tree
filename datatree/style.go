package datatree

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	FieldKey lipgloss.Style
}

var (
	styleDefault = Styles{
		FieldKey: lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
				Light: "#224",
				Dark:  "#b8e",
			}).
			Bold(true),
	}

	styleBlank = Styles{
		FieldKey: lipgloss.NewStyle(),
	}
)
