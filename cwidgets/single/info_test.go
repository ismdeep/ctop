package single

import (
	"testing"

	"github.com/mattn/go-runewidth"
)

func TestMkInfoRowsWrapsLongValues(t *testing.T) {
	rows := mkInfoRows("ENV", "abcdefghijklmnopqrstuvwxyz", 16)

	if len(rows) < 3 {
		t.Fatalf("expected wrapped rows, got %d", len(rows))
	}
	for _, row := range rows {
		if got := tableRowWidth(row); got > 16 {
			t.Fatalf("row width %d exceeds table width: %#v", got, row)
		}
	}
}

func TestWrapTextHandlesWideRunes(t *testing.T) {
	lines := wrapText("你好ab", 4)

	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(lines))
	}
	if lines[0] != "你好" || lines[1] != "ab" {
		t.Fatalf("unexpected wrapped lines: %#v", lines)
	}
}

func TestMkInfoRowsWrapsLongKeysToo(t *testing.T) {
	rows := mkInfoRows("VERY_LONG_ENVIRONMENT_KEY_NAME", "value", 20)

	if len(rows) < 2 {
		t.Fatalf("expected wrapped key rows, got %d", len(rows))
	}
	for _, row := range rows {
		if got := tableRowWidth(row); got > 20 {
			t.Fatalf("row width %d exceeds table width: %#v", got, row)
		}
	}
}

func tableRowWidth(row []string) int {
	const tablePadding = 8

	width := tablePadding
	for _, cell := range row {
		width += runewidth.StringWidth(cell)
	}
	return width
}
