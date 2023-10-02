package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	columnStyle = lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62"))
)

type model struct {
	ldiff     string
	rdiff     string
	lviewport viewport.Model
	rviewport viewport.Model
	ready     bool
}

func main() {
	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() model {
	return model{
		ldiff: "LEFT\nThis was the original\nNewlines\nEverything really\n1\n2\n3\n4\n5\n6\n7\n8\n9\n10",
		rdiff: "RIGHT\nThis is the changed one\nNewlines\nEverything really with changes\n1\n2\n3\n4\n5\n6\n7\n8\n9\n10",
		ready: false,
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		width := msg.Width / 2
		height := 10

		if !m.ready {
			m.lviewport = viewport.New(width, height)
			m.lviewport.YPosition = 10
			m.lviewport.SetContent(fmt.Sprint(msg.Width))

			m.rviewport = viewport.New(width, height)
			m.rviewport.YPosition = 10
			m.rviewport.SetContent(m.rdiff)

			m.ready = true
		} else {
			m.lviewport.SetContent(fmt.Sprint(msg.Width))
			m.lviewport.Width = width
			m.lviewport.Height = height

			m.rviewport.Width = width
			m.rviewport.Height = height
		}
	}

	// Handle keyboard and mouse events in the viewport
	m.lviewport, cmd = m.lviewport.Update(msg)
	cmds = append(cmds, cmd)

	m.rviewport, cmd = m.rviewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	// The header
	s := "Git diff\n\n"

	mainBody := lipgloss.JoinHorizontal(lipgloss.Left, m.lviewport.View(), m.rviewport.View())
	return fmt.Sprintf("%s%s", s, mainBody)
}
