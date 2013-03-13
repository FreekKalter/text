// Copyright 2013 (c) Freek Kalter. All rights reserved.
// Use of this source code is governed by the "Revised BSD License"
// that can be found in the LICENSE file.

// Package columnswriter imlements a write filter that prints evenly distributed
// columns fitted to current terminal window size.
package columnswriter

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// A Writer is a filter that prints evenly distributed columns
type Writer struct {
	output            io.Writer
	nrTerminalColumns int
	minWidth, padding int
	inputSep          rune

	text []byte
}

// Create a writer object with specified values, see Init for explanation of
// paramaters.
func New(output io.Writer, inputSep rune, minWidth, padding int) *Writer {
	return new(Writer).Init(output, inputSep, minWidth, padding)
}

// A Writer must be initialized with a call to Init. The first pararmater (output)
// specifies the filter output. The inputSep is the character by wich each field
// is seperated in the input later on.
//
//  minWidth: Is the minimum widht of a column
//  padding:  The number of spaces between columns
func (w *Writer) Init(output io.Writer, inputSep rune, minWidth, padding int) *Writer {
	w.output = output
	w.padding = padding
	w.minWidth = minWidth
	w.inputSep = inputSep
	nrTerminalColumnsInt32, _ := strconv.ParseInt(os.Getenv("COLUMNS"), 10, 32)
	w.nrTerminalColumns = int(nrTerminalColumnsInt32)
	return w
}

// Implements Writer interface, so you can call it with the whole Fprint family
func (w *Writer) Write(buf []byte) (n int, err error) {
	w.text = append(w.text, buf...)
	n = len(buf)
	return
}

func realLength(text string) (l int) {
	ansiEscBeginRegex := regexp.MustCompile(`^(\x1b\[\d+m)+`)
	ansiEscEndRegex := regexp.MustCompile(`(\x1b\[0m)+$`)
	replaced := ansiEscBeginRegex.ReplaceAllString(ansiEscEndRegex.ReplaceAllString(text, ""), "")
	l = len(replaced)
	return
}

// Does the actual printing, always call this after everthing you want to print has
// been printed.
func (w *Writer) Flush() {

	textString := strings.TrimRight(string(w.text), "\n")
	items := strings.FieldsFunc(textString, func(r rune) bool { return r == w.inputSep || r == '\n' })
	nrItems := len(items)
	var nrColumns, nrRows, totalWidth int = 0, 1, 0

	for _, file := range items {
		if (totalWidth + realLength(file) + 1) > w.nrTerminalColumns {
			break
		}
		totalWidth += realLength(file) + 2
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
				if realLength(items[index]) > maxColumnWidth {
					maxColumnWidth = realLength(items[index])
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
				columnWidth = realLength(items[index])
			}
			columnWidth += w.padding
			if columnWidth < w.minWidth {
				columnWidth = w.minWidth
			}
			// compensate difference in lenght when excluding ansi color escapes
			lenDiff := len(items[index]) - realLength(items[index])
			if x == (nrColumns - 1) { // no spaces at end of line (last column)
				fmt.Fprintf(w.output, "%s", items[index])
			} else {
				fmt.Fprintf(w.output, "%-*s", columnWidth+lenDiff, items[index])
			}
		}
		fmt.Fprint(w.output, "\n")
	}
}
