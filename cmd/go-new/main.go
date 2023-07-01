package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"go-new/pkg/scaffolding"
	"os"
	"strings"
)

func main() {
	model := initialModel()
	model.projectName = startupTextAndGetFolderName()
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func startupTextAndGetFolderName() string {
	fmt.Println("Make sure you are in the folder where you want to create project")
	wd, _ := os.Getwd()
	sep := "/"
	if strings.Contains(wd, "\\") {
		sep = "\\"
	}
	split := strings.Split(wd, sep)
	projectName := split[len(split)-1]
	fmt.Println("Your project will be called: " + projectName)
	return projectName
}

type projectTypes struct {
	projectName string
	choices     []scaffolding.ProjectScaffolding
	cursor      int
	selected    map[int]struct{}
}

func initialModel() projectTypes {
	return projectTypes{
		choices: scaffolding.ProjectList(),

		selected: make(map[int]struct{}),
	}
}

func (m projectTypes) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m projectTypes) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
				go func() {
					m.choices[m.cursor].CreateNew(m.projectName)
				}()
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m projectTypes) View() string {
	if len(m.selected) > 0 {
		return "Creating project..."
	}

	// The header
	s := "What Project Type should be created?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Name())
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
