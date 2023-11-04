package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func listenForActivity(sub chan responseMsg, conn myMpdConnection) tea.Cmd {
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
					lrcString = fmt.Sprintf("No .lrc file in the directory of %s", title)
				}

				parser.Parse(lrcString)

				msg = responseMsg{title: title, lyricType: parser.lrcType}

				line = title
				runner.Reset()
				sub <- msg
			}

			if elapsed != newElapsed {
				elapsed = newElapsed
				sub <- msg
			}

			time.Sleep(time.Millisecond * 250)
			runner.TimeUpdate(elapsed)
		}
	}
}

func waitForActivity(sub chan responseMsg) tea.Cmd {
	return func() tea.Msg {
		return <-sub
	}
}
