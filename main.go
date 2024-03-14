package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fhs/gompd/v2/mpd"
)

type LyricType uint

const (
	LyricSynced LyricType = iota
	LyricUnSynced
)

type responseMsg struct {
	title     string
	lyricType LyricType
}

type model struct {
	activeLine    int
	ready         bool
	followLine    bool
	title         string
	content       string
	lyricType     LyricType
	viewport      viewport.Model
	mpdConnection myMpdConnection
	state         chan responseMsg
}

var parser = Lrc{}
var runner = NewRunner(&parser, false)

func main() {
	conf := getConfig()
	mpdConnection := connect(fmt.Sprintf("%s:%d", conf.Host, conf.Port))

	defer func(conn *mpd.Client) {
		_ = conn.Close()
	}(mpdConnection)

	appModel := initModel(mpdConnection)

	p := tea.NewProgram(
		appModel,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Sorry, there's been an error: %v", err)
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
