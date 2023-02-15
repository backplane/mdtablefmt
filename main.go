package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// ColumnAlignment represents the alignment of the text in a markdown table
type ColumnAlignment int

const (
	// AlignLeft cell text is aligned to the left    ex: "text    "
	AlignLeft ColumnAlignment = iota
	// AlignRight cell text is aligned to the right  ex: "    text"
	AlignRight
	// AlignCenter cell text is centered in the cell ex: "  text  "
	AlignCenter
)

// MarkdownTable stores a markdown-style table of strings and the corresponding metadata including headings, column widths, and column alignments
type MarkdownTable struct {
	Headings         []string
	ColumnAlignments []ColumnAlignment
	ColumnWidths     map[int]int
	Rows             [][]string
}

var noStrip bool

const helpText = `usage: %s

utility for formatting Markdown tables; reads valid markdown table data on the
standard input and writes a properly spaced and padded version on the standard
output

For more information, visit the source code repository:
https://github.com/backplane/mdtablefmt

`

// NewMarkdownTable is the constructor for MarkdownTable
func NewMarkdownTable() *MarkdownTable {
	headings := make([]string, 0)
	columnAlignments := make([]ColumnAlignment, 0)
	columnWidths := make(map[int]int)
	rows := make([][]string, 0)
	return &MarkdownTable{
		Headings:         headings,
		ColumnAlignments: columnAlignments,
		ColumnWidths:     columnWidths,
		Rows:             rows,
	}
}

func init() {
	flag.BoolVar(&noStrip, "no-strip", false, "fully pad the rightmost output column instead of stripping trailing whitespace")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, helpText, os.Args[0])
		flag.PrintDefaults()
	}

}

func main() {
	flag.Parse()

	// Load a table from STDIN
	table, err := ParseMarkdownTable(os.Stdin)
	if err != nil {
		log.Fatalf("failed to parse markdown table from STDIN, error: %s", err)
	}

	// Print the table back out (reformatted)
	PrintMarkdownTable(os.Stdout, table, noStrip)
}
