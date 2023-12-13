package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var languages [10]string = [10]string{"C", "C++", "Java", "Python", "Go", "Rust", "JavaScript", "TypeScript", "Ruby", "PHP"}

type model struct {
	choices  []string
	cursor   int
	selected int
}

// Update implements tea.Model.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.selected = m.cursor
			// Start the spinner.
		}
	}
	return m, nil
}

// View implements tea.Model.
func (m model) View() string {
	s := "What project do you want to start?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		// do not check on initial render
		checked := " "
		if m.selected == i {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"
	return s
}

func initialModel() model {
	return model{
		choices:  languages[:],
		cursor:   0,
		selected: -1,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
