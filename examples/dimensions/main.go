package main

import (
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-data-tree/datatree"
)

type Trainer struct {
	Name        string
	Description string
	Age         int
	Hometown    string
	Pokemon     []Pokemon
}

type Pokemon struct {
	Name        string
	Description string
	Types       []string
}

type Model struct {
	simpleTree datatree.Model
}

func NewModel() Model {
	pikachu := Pokemon{
		Name: "Pikachu",
		Description: `Pikachu is a fictional species in the Pokémon media franchise. Designed by Atsuko Nishida and Ken Sugimori, Pikachu first appeared in the 1996 Japanese video games Pokémon Red and Green created by Game Freak and Nintendo, which were released outside of Japan in 1998 as Pokémon Red and Blue. Pikachu is a yellow, mouse-like creature with electrical abilities. It is a major character in the Pokémon franchise, serving as its mascot and as a major mascot for Nintendo.

Pikachu is widely considered to be the most popular and well-known Pokémon species, largely due to its appearance in the Pokémon anime television series as the companion of protagonist Ash Ketchum. In most vocalized appearances Pikachu is voiced by Ikue Ōtani, though it has been portrayed by other actors, notably Ryan Reynolds in the live-action animated film Pokémon Detective Pikachu. Pikachu has been well received by critics, with particular praise given for its cuteness, and has come to be regarded as an icon of Japanese pop culture.`,
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
