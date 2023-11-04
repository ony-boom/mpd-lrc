package main

import (
	"errors"
	"sort"
	"strconv"
)

type Runner struct {
	lrc          *Lrc
	currentIndex int
	offset       bool
}

func NewRunner(lrc *Lrc, offset bool) *Runner {
	return &Runner{
		lrc:          lrc,
		currentIndex: -1,
		offset:       offset,
	}
}

func (r *Runner) SetLrc(lrc *Lrc) {
	r.lrc = lrc.Clone()
	r.lrcUpdate()
}

func (r *Runner) lrcUpdate() {
	if r.offset {
		r.offsetAlign()
	}
	r.sortLyrics()
}

func (r *Runner) offsetAlign() {
	if offset, exists := r.lrc.Info["offset"]; exists {
		offsetValue, err := strconv.ParseFloat(offset, 64)
		if err == nil {
			r.lrc.Offset(offsetValue)
			delete(r.lrc.Info, "offset")
		}
	}
}

func (r *Runner) sortLyrics() {
	sort.SliceStable(r.lrc.Lyrics, func(i, j int) bool {
		return r.lrc.Lyrics[i].Timestamp < r.lrc.Lyrics[j].Timestamp
	})
}

func (r *Runner) TimeUpdate(timestamp float64) {
	r.currentIndex = r.findIndex(timestamp, r.currentIndex)
}

func (r *Runner) findIndex(timestamp float64, startIndex int) int {
	for i, lyric := range r.lrc.Lyrics {
		if timestamp >= lyric.Timestamp {
			startIndex = i
		} else {
			break
		}
	}
	return startIndex
}

func (r *Runner) GetInfo() Info {
	return r.lrc.Info
}

func (r *Runner) GetLyrics() []Lyric {
	return r.lrc.Lyrics
}

func (r *Runner) GetLyric(index int) (Lyric, error) {
	if index >= 0 && index < len(r.lrc.Lyrics) {
		return r.lrc.Lyrics[index], nil
	}
	return Lyric{}, errors.New("index not exist")
}

func (r *Runner) CurIndex() int {
	return r.currentIndex
}

func (r *Runner) CurLyric() (Lyric, error) {
	return r.GetLyric(r.currentIndex)
}

func (r *Runner) Reset() {
	r.currentIndex = -1
}
