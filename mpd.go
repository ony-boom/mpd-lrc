package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fhs/gompd/v2/mpd"
)

func connect(host string) (conn *mpd.Client) {
	conn, err := mpd.Dial("tcp", host)
	if err != nil {
		log.Fatalf("Can't connect to host: '%s'", host)
	}

	return
}

type myMpdConnection struct {
	conn *mpd.Client
}

func (c myMpdConnection) getTitle() string {
	currentSong, _ := c.conn.CurrentSong()
	title := fmt.Sprintf("%s - %s", currentSong["Artist"], currentSong["Title"])

	return title
}

func (c myMpdConnection) getElapsed() float64 {
	status, _ := c.conn.Status()
	elsapsed, _ := strconv.ParseFloat(status["elapsed"], 32)

	return toFixed(elsapsed, 3)
}

func (c myMpdConnection) getLrcString() (string, error) {
	currenSong, err := c.conn.CurrentSong()
	if err != nil {
		return "", fmt.Errorf("failed to get current song: %w", err)
	}

	audioPath, ok := currenSong["file"]
	if !ok || audioPath == "" {
		return "", fmt.Errorf("no audio file in current song")
	}

	ext := filepath.Ext(audioPath)
	lrcPath := strings.TrimSuffix(audioPath, ext) + ".lrc"

	fullPath := filepath.Join(getConfig().MusicPath, lrcPath)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to read lrc file %q: %w", fullPath, err)
	}

	return strings.TrimSpace(string(data)), nil
}

func (c *myMpdConnection) initConn(conn *mpd.Client) {
	c.conn = conn
}
