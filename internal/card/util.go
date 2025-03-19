package card

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	maxDescriptionLineLength = 60
)

func wrapDescription(description string) []string {
	splitter := " "
	words := strings.Split(description, splitter)
	line := ""
	lines := []string{}
	for _, word := range words {
		if len(line+word) > maxDescriptionLineLength {
			line = strings.TrimSpace(line)
			if line != "" {
				lines = append(lines, line)
				line = ""
			}
		}
		line += word + splitter
	}
	if line != "" {
		lines = append(lines, strings.TrimSpace(line))
	}
	return lines
}

func formatCount(count int) string {
	if count < 950 {
		return strconv.Itoa(count)
	}
	k := float64(count) / 1000
	if k < 99.9 {
		return fmt.Sprintf("%.1fk", k)
	}
	return fmt.Sprintf("%dk", int(math.Round(k)))
}
