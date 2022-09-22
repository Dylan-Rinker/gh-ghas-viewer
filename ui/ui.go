package ui

import (
	"log"
	"strings"

	"github.com/Dylan-Rinker/gh-ghas-viewer/config"
	"github.com/Dylan-Rinker/gh-ghas-viewer/ui/components/help"
	"github.com/Dylan-Rinker/gh-ghas-viewer/ui/context"
	"github.com/Dylan-Rinker/gh-ghas-viewer/ui/keys"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	keys keys.KeyMap
	// sidebar       sidebar.Model
	// prSidebar     prsidebar.Model
	// issueSidebar  issuesidebar.Model
	// currSectionId int
	help help.Model
	// alerts        []section.Section
	// ready         bool
	// isSidebarOpen bool
	// tabs tabs.Model
	ctx context.ProgramContext
	// // taskSpinner   spinner.Model
	// tasks map[string]context.Task
}

func NewModel() Model {
	// tabsModel := tabs.NewModel()
	m := Model{
		keys: keys.Keys,
		help: help.NewModel(),
		// currSectionId: 1,
		// tabs:          tabsModel,
		// sidebar:       sidebar.NewModel(),
		// prSidebar:     prsidebar.NewModel(),
		// issueSidebar:  issuesidebar.NewModel(),
		// taskSpinner:   spinner.Model{Spinner: spinner.Dot},
		// tasks:         map[string]context.Task{},
	}

	// m.ctx = context.ProgramContext{StartTask: func(task context.Task) tea.Cmd {
	// 	task.StartTime = time.Now()
	// 	m.tasks[task.Id] = task
	// 	return m.taskSpinner.Tick
	// }}
	return m
}

func initScreen() tea.Msg {
	var err error

	config, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	return initMsg{Config: config}
}

func (m Model) Init() tea.Cmd {
	// return tea.Batch(initScreen, tea.EnterAltScreen)
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd tea.Cmd
		// sidebarCmd      tea.Cmd
		// prSidebarCmd    tea.Cmd
		// issueSidebarCmd tea.Cmd
		helpCmd tea.Cmd
		cmds    []tea.Cmd
		// currSection     = m.getCurrSection()
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.ctx.Error = nil

		switch {
		case key.Matches(msg, m.keys.Quit):
			// cmd = tea.Quit
			return m, tea.Quit
		}
	}

	cmds = append(
		cmds,
		cmd,
		// sidebarCmd,
		helpCmd,
		// sectionCmd,
		// prSidebarCmd,
		// issueSidebarCmd,
	)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := strings.Builder{}
	// s.WriteString(m.tabs.View(m.ctx))
	s.WriteString("GHAS Viewer App\n")
	// mainContent := ""

	// mainContent = lipgloss.JoinHorizontal(
	// 	lipgloss.Top,
	// m.getCurrSection().View(),
	// 	m.sidebar.View(),
	// )

	// s.WriteString(mainContent)
	// s.WriteString("\n")

	// if m.ctx.Error != nil {
	// 	s.WriteString(
	// 		styles.ErrorStyle.
	// 			Width(m.ctx.ScreenWidth).
	// 			Render(fmt.Sprintf("%s %s",
	// 				constants.FailureGlyph,
	// 				lipgloss.NewStyle().
	// 					Foreground(styles.DefaultTheme.WarningText).
	// 					Render(m.ctx.Error.Error()),
	// 			)),
	// 	)
	// } else if len(m.tasks) > 0 {
	// 	s.WriteString(m.renderRunningTask())
	// } else {
	s.WriteString(m.help.View(m.ctx))
	// }

	return s.String()
}

type initMsg struct {
	Config config.Config
}
