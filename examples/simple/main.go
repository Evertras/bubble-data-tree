package main

import (
	"log"
	"strings"

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
	Name  string
	Types []string
}

type Model struct {
	simpleTree datatree.Model
}

func NewModel() Model {
	pikachu := Pokemon{
		Name:  "Pikachu",
		Types: []string{"Electric"},
	}

	pidgey := Pokemon{
		Name:  "Pidgey",
		Types: []string{"Flying"},
	}

	torterra := Pokemon{
		Name:  "Torterra",
		Types: []string{"Grass", "Ground"},
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
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	body := strings.Builder{}

	body.WriteString("A very simple default data tree (non-interactive)\nPress q or ctrl+c to quit\n\n")

	body.WriteString(m.simpleTree.View())
	body.WriteString("\n")

	return body.String()
}

func main() {
	p := tea.NewProgram(NewModel())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
