package main

import (
	"sort"
	"strings"
)

type Lyric struct {
	Timestamp float64
	Content   string
	Active    bool
}

type CombineLyric struct {
	Timestamps []float64
	Content    string
}

type Info map[string]string

type LineFormat string

const (
	CRLF LineFormat = "\r\n"
	CR   LineFormat = "\r"
	LF   LineFormat = "\n"
)

type ToStringOptions struct {
	Combine    bool
	Sort       bool
	LineFormat LineFormat
}

type Lrc struct {
	Info   Info
	Lyrics []Lyric
}

func (lrc *Lrc) Parse(text string) {
	lines := strings.Split(text, "\n")
	lyrics := make([]Lyric, 0)
	info := make(Info)

	for _, line := range lines {
		parsedLine := parseLine(line)
		switch parsedLine := parsedLine.(type) {
		case InfoLine:
			info[parsedLine.Key] = parsedLine.Value
		case TimeLine:
			for _, timestamp := range parsedLine.Timestamps {
				lyrics = append(lyrics, Lyric{Timestamp: timestamp, Content: parsedLine.Content})
			}
		}
	}

	lrc.Info = info

	sort.SliceStable(lyrics, func(i, j int) bool {
		return lyrics[i].Timestamp < lyrics[j].Timestamp
	})

	lrc.Lyrics = lyrics
}

func (lrc *Lrc) Offset(offsetTime float64) {
	for i := range lrc.Lyrics {
		lrc.Lyrics[i].Timestamp += offsetTime
		if lrc.Lyrics[i].Timestamp < 0 {
			lrc.Lyrics[i].Timestamp = 0
		}
	}
}
