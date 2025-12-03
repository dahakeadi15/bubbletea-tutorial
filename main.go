package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// application State
type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

// define initialstate of the application
func initialModel() model {
	return model{
		choices:  []string{"Buy Apples", "Buy Dates", "Buy IceCream"},
		selected: make(map[int]struct{}),
	}
}

// Init can return an initial `Cmd` that can perform some initial I/O
func (m model) Init() tea.Cmd {
	return nil
}

// Update is called when things happen, it checks what happened and updates the
// model accordingly. Sometimes it returns a `Cmd` to make more things happen.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			m.cursor--

		case "down", "j":
			m.cursor++

		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

// Renders the UI of the application
func (m model) View() string {
	s := "What should we buy at the market?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
