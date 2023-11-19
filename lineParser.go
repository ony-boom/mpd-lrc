package main

import (
	"regexp"
	"strconv"
	"strings"
)

var TAGS_REGEXP = regexp.MustCompile(`^(\[.+\])+`)
var INFO_REGEXP = regexp.MustCompile(`^\s*(\w+)\s*:(.*)$`)
var TIME_REGEXP = regexp.MustCompile(`^\s*(\d+)\s*:\s*(\d+(\s*[\.:]\s*\d+)?)\s*$`)

type LineType string

const (
	INVALID LineType = "INVALID"
	INFO    LineType = "INFO"
	TIME    LineType = "TIME"
)

type InvalidLine struct {
	Type LineType
}

type TimeLine struct {
	Type       LineType
	Timestamps []float64
	Content    string
}

type InfoLine struct {
	Type  LineType
	Key   string
	Value string
}

func parseTags(line string) (tags []string, content string) {
	line = strings.TrimSpace(line)
	matches := TAGS_REGEXP.FindStringSubmatch(line)
	if len(matches) == 0 {
		return nil, line
	}
	tag := matches[0]
	content = line[len(tag):]
	tags = strings.Split(tag[1:len(tag)-1], "][")

	return tags, content
}

func parseTime(tags []string, content string) TimeLine {
	timestamps := make([]float64, 0)
	for _, tag := range tags {
		matches := TIME_REGEXP.FindStringSubmatch(tag)
		if len(matches) != 0 {
			minutes, _ := strconv.Atoi(matches[1])
			seconds, _ := strconv.ParseFloat(strings.ReplaceAll(matches[2], " ", ""), 64)
			timestamps = append(timestamps, toFixed(float64(minutes*60)+seconds, 4))
		}
	}
	return TimeLine{
		Type:       TIME,
		Timestamps: timestamps,
		Content:    strings.TrimSpace(content),
	}
}

func parseInfo(tag string) InfoLine {
	matches := INFO_REGEXP.FindStringSubmatch(tag)
	if len(matches) != 0 {
		return InfoLine{
			Type:  INFO,
			Key:   strings.TrimSpace(matches[1]),
			Value: strings.TrimSpace(matches[2]),
		}
	}
	return InfoLine{
		Type: INVALID,
	}
}

func parseLine(line string) interface{} {
	tags, content := parseTags(line)
	if len(tags) > 0 {
		if TIME_REGEXP.MatchString(tags[0]) {
			return parseTime(tags, content)
		} else {
			return parseInfo(tags[0])
		}
	}
	return InvalidLine{
		Type: INVALID,
	}
}
