package main

import (
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-data-tree/datatree"
)

type Trainer struct {
	Name     string
	Age      int
	Hometown string
	Pokemon  []Pokemon
}

type Pokemon struct {
	Name        string
	Description string
	Types       []string
}

type Model struct {
	simpleTree datatree.Model
	viewport   viewport.Model

	ready bool
}

func NewModel() Model {
	pikachu := Pokemon{
		Name:        "Pikachu",
		Description: "Pikachu that can generate powerful electricity have cheek sacs that are extra soft and super stretchy.",
		Types:       []string{"Electric"},
	}

	pidgey := Pokemon{
		Name:        "Pidgey",
		Description: "Very docile. If attacked, it will often kick up sand to protect itself rather than fight back.",
		Types:       []string{"Normal", "Flying"},
	}

	torterra := Pokemon{
		Name:        "Torterra",
		Description: "ちいさな　ポケモンたちが　あつまり　うごかない　ドダイトスの　せなかで　すづくりを　はじめることがある。　（『ポケットモンスター ブリリアントダイヤモンド』より）",
		Types:       []string{"Grass", "Ground"},
	}

	ash := Trainer{
		Name:     "サトシ",
		Age:      14,
		Hometown: "Pallet Town",
		Pokemon:  []Pokemon{pikachu, pidgey, torterra},
	}

	return Model{
		simpleTree: datatree.New(ash),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.simpleTree, cmd = m.simpleTree.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			cmds = append(cmds, tea.Quit)
		}

	case tea.WindowSizeMsg:
		// The help message at top
		const headerHeight = 3
		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-headerHeight)
			m.viewport.HighPerformanceRendering = true
			m.viewport.SetContent(m.simpleTree.View())
			m.viewport.YPosition = headerHeight
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - headerHeight
		}

		cmds = append(cmds, viewport.Sync(m.viewport))
	}

	m.viewport, cmd = m.viewport.Update(msg)

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	body := strings.Builder{}

	body.WriteString("Data tree using a viewport for lots of data\nPress up/down to scroll, q or ctrl+c to quit\n\n")

	body.WriteString(m.viewport.View())

	return body.String()
}

func main() {
	p := tea.NewProgram(
		NewModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}