package single

import (
	"strings"

	ui "github.com/gizak/termui"
)

var displayInfo = []string{"id", "name", "image", "ports", "IPs", "state", "created", "uptime", "health"}

type Info struct {
	*ui.Table
	data map[string]string
}

func NewInfo() *Info {
	p := ui.NewTable()
	p.Height = 4
	p.Width = colWidth[0]
	p.FgColor = ui.ThemeAttr("par.text.fg")
	p.Separator = false
	i := &Info{p, make(map[string]string)}
	return i
}

func (w *Info) Set(k, v string) {
	w.data[k] = v

	// rebuild rows
	w.Rows = [][]string{}
	for _, k := range displayInfo {
		if v, ok := w.data[k]; ok {
			w.Rows = append(w.Rows, mkInfoRows(k, v, w.Width)...)
		}
	}

	w.Height = len(w.Rows) + 2
}

// Build row(s) from a key and value string
func mkInfoRows(k, v string, tableWidth int) (rows [][]string) {
	lines := strings.Split(v, "\n")
	keyWidth, valueWidth := infoColumnWidths(tableWidth)
	keyLines := wrapText(k, keyWidth)
	valueLines := wrapText(lines[0], valueWidth)
	rows = append(rows, zipRows(keyLines, valueLines)...)

	// append any additional lines in separate row
	if len(lines) > 1 {
		for _, line := range lines[1:] {
			if line != "" {
				rows = append(rows, zipRows([]string{""}, wrapText(line, valueWidth))...)
			}
		}
	}

	return rows
}

func infoColumnWidths(tableWidth int) (int, int) {
	const tablePadding = 8
	const minKeyWidth = 8
	const maxKeyWidth = 24

	contentWidth := tableWidth - tablePadding
	if contentWidth < 2 {
		return 1, 1
	}

	keyWidth := contentWidth / 3
	if contentWidth >= minKeyWidth+4 && keyWidth < minKeyWidth {
		keyWidth = minKeyWidth
	}
	if keyWidth > maxKeyWidth {
		keyWidth = maxKeyWidth
	}
	minValueWidth := contentWidth / 2
	if minValueWidth < 1 {
		minValueWidth = 1
	}
	if keyWidth > contentWidth-minValueWidth {
		keyWidth = contentWidth - minValueWidth
	}
	if keyWidth < 1 {
		keyWidth = 1
	}

	valueWidth := contentWidth - keyWidth
	if valueWidth < 1 {
		valueWidth = 1
	}
	return keyWidth, valueWidth
}

func zipRows(left, right []string) (rows [][]string) {
	rowCount := len(left)
	if len(right) > rowCount {
		rowCount = len(right)
	}
	if rowCount == 0 {
		return nil
	}

	for i := 0; i < rowCount; i++ {
		l := ""
		r := ""
		if i < len(left) {
			l = left[i]
		}
		if i < len(right) {
			r = right[i]
		}
		rows = append(rows, []string{l, r})
	}
	return rows
}
