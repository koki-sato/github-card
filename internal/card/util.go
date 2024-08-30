package card

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func wrapDescription(description string) []string {
	splitter := " "
	words := strings.Split(description, splitter)
	line := ""
	lines := []string{}
	for _, word := range words {
		line += word + splitter
		if len(line) > 55 {
			lines = append(lines, line)
			line = ""
		}
	}
	if line != "" {
		lines = append(lines, line)
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
