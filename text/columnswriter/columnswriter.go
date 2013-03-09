package columnswriter

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Writer struct {
	output            io.Writer
	nrTerminalColumns int
	minWidth, padding int
	tabChar           rune

	text []byte
}

func New(output io.Writer, tabChar rune, minWidth, padding int) *Writer {
	return new(Writer).Init(output, tabChar, minWidth, padding)
}

func (w *Writer) Init(output io.Writer, tabChar rune, minWidth, padding int) *Writer {
	w.output = output
	w.padding = padding
	w.minWidth = minWidth
	w.tabChar = tabChar
	nrTerminalColumnsInt32, _ := strconv.ParseInt(os.Getenv("COLUMNS"), 10, 32)
	w.nrTerminalColumns = int(nrTerminalColumnsInt32)
	return w
}

func (w *Writer) Write(buf []byte) (n int, err error) {
	w.text = append(w.text, buf...)
	n = len(buf)
	return
}

func (w *Writer) Flush() {

	textString := strings.TrimRight(string(w.text), "\n")
	items := strings.FieldsFunc(textString, func(r rune) bool { return r == w.tabChar })
	nrItems := len(items)
	var nrColumns, nrRows, totalWidth int = 0, 1, 0

	for _, file := range items {
		if (totalWidth + len(file) + 1) > w.nrTerminalColumns {
			break
		}
		totalWidth += len(file) + 2
		nrColumns++
	}
	calcNrRows := func(items, columns int) int {
		return int(math.Ceil(float64(items) / float64(columns)))
	}
	nrRows = calcNrRows(nrItems, nrColumns)

	totalWidth = totalWidth * 2
	var columnWidths []int
	for totalWidth > w.nrTerminalColumns {
		totalWidth = 0
		columnWidths = []int{}
		for x := 0; x < nrColumns; x++ {
			maxColumnWidth := 0
			for y := 0; y < nrRows; y++ {
				index := y*nrColumns + x
				if index >= nrItems {
					break
				}
				if len(items[index]) > maxColumnWidth {
					maxColumnWidth = len(items[index])
				}
			}
			totalWidth += maxColumnWidth + 2
			columnWidths = append(columnWidths, maxColumnWidth)

		}
		if totalWidth > w.nrTerminalColumns {
			nrColumns--
			nrRows = calcNrRows(nrItems, nrColumns)
		}
	}

	for y := 0; y < nrRows; y++ {
		for x := 0; x < nrColumns; x++ {
			index := y*nrColumns + x
			if index >= nrItems {
				break
			}

			var columnWidth int
			if len(columnWidths) > 0 {
				columnWidth = columnWidths[x]
			} else {
				columnWidth = len(items[index])
			}
			columnWidth += w.padding
			if columnWidth < w.minWidth {
				columnWidth = w.minWidth
			}
			fmt.Fprintf(w.output, "%-*s", columnWidth, items[index])
		}
		fmt.Fprint(w.output, "\n")
	}
}
