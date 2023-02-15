package main

import (
	"fmt"
	"io"
	"strings"
)

func ceildiv(x, y int) int {
	result := x / y
	if x%y > 0 {
		result++
	}
	return result
}

func spacePad(s string, width int, alignment ColumnAlignment) string {
	var fmtWidth int

	switch alignment {
	case AlignRight:
		fmtWidth = width
	case AlignLeft:
		fmtWidth = 0 - width
	case AlignCenter:
		s = strings.Repeat(" ", ceildiv(width-len(s), 2)) + s
		fmtWidth = 0 - width
	default:
		panic(fmt.Sprintf("spacePad attempt to pad with unknown alignment: %+v", alignment))
	}

	return fmt.Sprintf("%[2]*[1]s", s, fmtWidth)
}

// PrintMarkdownTable prints the given MarkdownTable to STDOUT
func PrintMarkdownTable(out io.Writer, table *MarkdownTable, keepTrailingWhitespace bool) {
	// short-circuit if there is nothing to print
	if len(table.Rows) < 1 {
		return
	}

	columnCount := len(table.Headings) - 1
	var paddedContent string

	// print the header row
	for colNum, header := range table.Headings {
		paddedContent := spacePad(header, table.ColumnWidths[colNum], table.ColumnAlignments[colNum])
		if colNum < columnCount {
			fmt.Fprint(out, paddedContent, " | ")
			continue
		}
		if !keepTrailingWhitespace {
			paddedContent = strings.TrimRight(paddedContent, " ")
		}
		fmt.Fprintln(out, paddedContent)
	}

	// print the delimiter row
	for colNum := range table.Headings {
		switch table.ColumnAlignments[colNum] {
		case AlignRight:
			fmt.Fprint(out, strings.Repeat("-", table.ColumnWidths[colNum]-1), ":")
		case AlignCenter:
			fmt.Fprint(out, ":", strings.Repeat("-", table.ColumnWidths[colNum]-2), ":")
		default:
			fmt.Fprint(out, strings.Repeat("-", table.ColumnWidths[colNum]))
		}
		if colNum < columnCount {
			fmt.Fprint(out, " | ")
		} else {
			fmt.Fprintln(out, "")
		}
	}

	// print the data rows
	for _, row := range table.Rows {
		for colNum, cell := range row {
			paddedContent = spacePad(cell, table.ColumnWidths[colNum], table.ColumnAlignments[colNum])
			if colNum < columnCount {
				fmt.Fprint(out, paddedContent, " | ")
				continue
			}
			if !keepTrailingWhitespace {
				paddedContent = strings.TrimRight(paddedContent, " ")
			}
			fmt.Fprintln(out, paddedContent)
		}
	}
}
