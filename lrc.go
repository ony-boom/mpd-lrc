package main

import (
	"strings"
)

type Lyric struct {
	Timestamp float64
	Content   string
}

type Info map[string]string

type Lrc struct {
	Info    Info
	Lyrics  []Lyric
	lrcType LyricType
}

func (lrc *Lrc) Parse(text string) {
	lines := strings.Split(text, "\n")
	lyrics := make([]Lyric, 0)
	info := make(Info)
	lrc.lrcType = LyricSynced

	for _, line := range lines {
		parsedLine := parseLine(line)
		switch parsedLine := parsedLine.(type) {
		case InfoLine:
			info[parsedLine.Key] = parsedLine.Value
		case TimeLine:
			for _, timestamp := range parsedLine.Timestamps {
				lyrics = append(lyrics, Lyric{Timestamp: timestamp, Content: parsedLine.Content})
			}

		case InvalidLine:
			lyrics = append(lyrics, Lyric{Timestamp: 0, Content: line})
			lrc.lrcType = LyricUnSynced
		}
	}

	lrc.Info = info
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

func (lrc *Lrc) Clone() *Lrc {
	clonedLrc := &Lrc{
		Info:    make(Info),
		Lyrics:  make([]Lyric, len(lrc.Lyrics)),
		lrcType: lrc.lrcType,
	}

	for key, value := range lrc.Info {
		clonedLrc.Info[key] = value
	}

	copy(clonedLrc.Lyrics, lrc.Lyrics)

	return clonedLrc
}

func (lrc *Lrc) ToString() string {
	content := ""
	for _, lrcLine := range lrc.Lyrics {
		content += lrcLine.Content + "\n"
	}

	content = strings.TrimSpace(content)
	return content
}
