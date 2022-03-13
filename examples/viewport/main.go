package main

import (
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-data-tree/datatree"
)

type Trainer struct {
	Name        string
	Age         int
	Hometown    string
	Description string
	Pokemon     []Pokemon
}

type Pokemon struct {
	Age         int
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
		Name: "Pikachu",
		Description: `Pikachu is a fictional species in the Pokémon media franchise. Designed by Atsuko Nishida and Ken Sugimori, Pikachu first appeared in the 1996 Japanese video games Pokémon Red and Green created by Game Freak and Nintendo, which were released outside of Japan in 1998 as Pokémon Red and Blue. Pikachu is a yellow, mouse-like creature with electrical abilities. It is a major character in the Pokémon franchise, serving as its mascot and as a major mascot for Nintendo.

Pikachu is widely considered to be the most popular and well-known Pokémon species, largely due to its appearance in the Pokémon anime television series as the companion of protagonist Ash Ketchum. In most vocalized appearances Pikachu is voiced by Ikue Ōtani, though it has been portrayed by other actors, notably Ryan Reynolds in the live-action animated film Pokémon Detective Pikachu. Pikachu has been well received by critics, with particular praise given for its cuteness, and has come to be regarded as an icon of Japanese pop culture.`,
		Types: []string{"Electric"},
	}

	pidgey := Pokemon{
		Name:        "Pidgey",
		Description: "Very docile.\nIf attacked, it will often kick up sand to protect itself rather than fight back.",
		Types:       []string{"Normal", "Flying"},
	}

	torterra := Pokemon{
		Name:        "Torterra",
		Description: "ちいさな　ポケモンたちが　あつまり　うごかない　ドダイトスの　せなかで　すづくりを　はじめることがある。　（『ポケットモンスター ブリリアントダイヤモンド』より）",
		Types:       []string{"Grass", "Ground"},
	}

	dragonite := Pokemon{
		Age:         3,
		Name:        "Dragonite",
		Description: "Dragonite is a draconic, bipedal Pokémon with light orange skin.",
		Types:       []string{"Dragon", "Flying"},
	}

	ash := Trainer{
		Name:     "サトシ",
		Age:      14,
		Hometown: "Pallet Town",
		Description: `Ash Ketchum (Japanese: サトシ Satoshi) is the main character of the Pokémon anime. He is also the main character of various manga based on the anime, including The Electric Tale of Pikachu, Ash & Pikachu, and Pocket Monsters Diamond & Pearl.

He is a Pokémon Trainer from Pallet Town whose goal is to become a Pokémon Master. His starter Pokémon was a Pikachu that he received from Professor Oak after arriving late at his laboratory. In Pokémon the Series: Sun & Moon, he becomes the first Champion of the Alola region's Pokémon League.`,
		Pokemon: []Pokemon{pikachu, pidgey, torterra, dragonite},
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
		m.simpleTree = m.simpleTree.WithWidth(msg.Width)
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
