package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fhs/gompd/v2/mpd"
)

type responseMsg struct {
	title string
	lines []Lyric
}

type model struct {
	activeLine    int
	ready         bool
	followLine    bool
	title         string
	content       string
	viewport      viewport.Model
	mpdConnection myMpdConnection
	state         chan responseMsg
}

func main() {
	conf := getConfig()
	mpdConnection := connect(fmt.Sprintf("%s:%d", conf.Host, conf.Port))

	defer func(conn *mpd.Client) {
		_ = conn.Close()
	}(mpdConnection)

	p := tea.NewProgram(
		initModel(mpdConnection),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initModel(mpdConn *mpd.Client) model {
	myMpdConnection := myMpdConnection{}
	myMpdConnection.initConn(mpdConn)

	title := myMpdConnection.getTitle()

	return model{
		content:       "",
		followLine:    true,
		title:         title,
		mpdConnection: myMpdConnection,
		state:         make(chan responseMsg),
	}
}
