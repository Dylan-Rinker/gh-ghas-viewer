package ui

import (
	"fmt"
	"strings"

	"github.com/Dylan-Rinker/gh-ghas-viewer/ui/styles"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	keys          keys.KeyMap
	sidebar       sidebar.Model
	prSidebar     prsidebar.Model
	issueSidebar  issuesidebar.Model
	currSectionId int
	help          help.Model
	alerts        []section.Section
	ready         bool
	isSidebarOpen bool
	// tabs          tabs.Model
	ctx context.ProgramContext
	// taskSpinner   spinner.Model
	tasks map[string]context.Task
}

func NewModel() Model {

	m := Model{}

	return m
}

func initScreen() tea.Msg {
	// var err error

	// config, err := config.ParseConfig()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// return initMsg{Config: config}
	return nil
}

func (m Model) Init() tea.Cmd {
	// return tea.Batch(initScreen, tea.EnterAltScreen)
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// switch msg := msg.(type) {
	// case tea.KeyMsg:
	// 	switch msg.String() {
	// 	case "ctrl+c", "q":
	// 		return m, tea.Quit
	// 	case "up", "k":
	// 		if m.cursor > 0 {
	// 			m.cursor--
	// 		}
	// 	case "down", "j":
	// 		if m.cursor < len(m.choices)-1 {
	// 			m.cursor++
	// 		}
	// 	case "enter", " ":
	// 		_, ok := m.selected[m.cursor]
	// 		if ok {
	// 			delete(m.selected, m.cursor)
	// 		} else {
	// 			m.selected[m.cursor] = struct{}{}
	// 		}
	// 	}
	// }

	return m, nil
}

func (m Model) View() string {
	s := strings.Builder{}
	// s.WriteString(m.tabs.View(m.ctx))
	s.WriteString("\n")
	mainContent := ""

	mainContent = lipgloss.JoinHorizontal(
		lipgloss.Top,
		// m.getCurrSection().View(),
		m.sidebar.View(),
	)

	s.WriteString(mainContent)
	s.WriteString("\n")

	if m.ctx.Error != nil {
		s.WriteString(
			styles.ErrorStyle.
				Width(m.ctx.ScreenWidth).
				Render(fmt.Sprintf("%s %s",
					constants.FailureGlyph,
					lipgloss.NewStyle().
						Foreground(styles.DefaultTheme.WarningText).
						Render(m.ctx.Error.Error()),
				)),
		)
	} else if len(m.tasks) > 0 {
		s.WriteString(m.renderRunningTask())
	} else {
		s.WriteString(m.help.View(m.ctx))
	}

	return s.String()
}
