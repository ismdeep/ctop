package single

import "github.com/mattn/go-runewidth"

func wrapText(text string, width int) []string {
	if text == "" {
		return []string{""}
	}
	if width < 1 {
		width = 1
	}

	var (
		lines []string
		line  []rune
		w     int
	)

	for _, r := range text {
		rw := runewidth.RuneWidth(r)
		if rw == 0 {
			rw = 1
		}
		if w+rw > width && len(line) > 0 {
			lines = append(lines, string(line))
			line = line[:0]
			w = 0
		}
		line = append(line, r)
		w += rw
	}

	if len(line) > 0 {
		lines = append(lines, string(line))
	}

	return lines
}
