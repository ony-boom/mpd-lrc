package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const useHighPerformanceRenderer = false

var (
	paddingBlock, paddingInline = 1, 1
	basePadding                 = 1

	baseStyle = lipgloss.NewStyle().Padding(basePadding)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(getConfig().TitleColor)).
			Border(lipgloss.NormalBorder(), true, false, true, false).
			Padding(paddingBlock, paddingInline)

	followActiveStyle = lipgloss.NewStyle().Padding(paddingBlock*2, 0)

	activeLineStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(getConfig().ActiveLineColor)).
			Bold(true)
)

func (m model) Init() tea.Cmd {
	return tea.Batch(listenForActivity(m.state, m.mpdConnection), waitForActivity(m.state))
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}

		if msg.String() == "ctrl+f" {
			m.followLine = !m.followLine && m.isLyricSynced()
		}

	case responseMsg:
		m.title = msg.title
		m.lyricType = msg.lyricType

		if m.isLyricSynced() {
			m.content, m.activeLine = getSyncedContent()
		} else {
			m.followLine = false
			m.content = parser.ToString()
		}

		cmds = append(cmds, waitForActivity(m.state))

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-headerHeight)
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - headerHeight
		}
	}

	// Just in case I decided to add style to this
	content := m.viewport.Style.Render(m.content)
	m.viewport.SetContent(content)

	if m.lyricType == LyricSynced {
		isVisible := isLineVisible(m.viewport.YOffset, m.viewport.Height, m.activeLine, m.viewport.TotalLineCount())

		if !isVisible && m.followLine {
			m.viewport.SetYOffset(m.activeLine)
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return baseStyle.Render(fmt.Sprintf("%s\n%s", m.headerView(), m.viewport.View()))
}

func getSyncedContent() (string, int) {
	content := ""

	for index, lrcLine := range parser.Lyrics {
		newLine := lipgloss.NewStyle().UnsetForeground().Render(lrcLine.Content)

		if index == runner.CurIndex() {
			newLine = activeLineStyle.Render(newLine)
		}

		if lrcLine.Timestamp != parser.Lyrics[len(parser.Lyrics)-1].Timestamp {
			content += newLine + "\n"
		}
	}

	return content, runner.CurIndex()
}

func isLineVisible(YOffset, height, lineIndex, totalLines int) bool {
	if lineIndex < 0 || lineIndex >= totalLines {
		return false
	}

	top := max(0, YOffset)
	bottom := clamp(YOffset+height, top, totalLines)

	return lineIndex >= top && lineIndex < bottom
}

func (m model) headerView() string {
	syncIcon := "󰯓"
	followActiveStyle.Foreground(lipgloss.Color("2"))

	if !m.followLine {
		syncIcon = "󱔶"
		followActiveStyle.Foreground(lipgloss.Color("9"))
	}

	syncIcon = followActiveStyle.Render(syncIcon)
	title := titleStyle.Width((m.viewport.Width - paddingInline - basePadding) - lipgloss.Width(syncIcon)).Render(m.title)

	return lipgloss.JoinHorizontal(lipgloss.Top, title, syncIcon)
}

func (m model) isLyricSynced() bool {
	return m.lyricType == LyricSynced
}
