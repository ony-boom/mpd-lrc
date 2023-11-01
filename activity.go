package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func listenForActivity(sub chan responseMsg, conn myMpdConnection) tea.Cmd {
	parser := Lrc{}
	return func() tea.Msg {
		line := ""
		var elapsed float64 = 0

		var msg responseMsg

		for {
			title := conn.getTitle()
			newElapsed := conn.getElapsed()

			if line != title {
				lrcString, err := conn.getLrcString()

				if err != nil {
					lrcString = fmt.Sprintf("[00:00]No .lrc file in the directory of %s", title)
				}
				parser.Parse(lrcString)

				msg = responseMsg{title: title, lines: parser.Lyrics}

				line = title
				sub <- msg
			}

			if elapsed != newElapsed {

				setActiveLine(newElapsed, &msg.lines)

				elapsed = newElapsed

				sub <- msg
			}

			time.Sleep(time.Millisecond * 250)
		}
	}
}

func waitForActivity(sub chan responseMsg) tea.Cmd {
	return func() tea.Msg {
		return <-sub
	}
}
