// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.

// Exercise 1.4: Modify dup2 to print the names of all files in which each duplciated line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for file, dups := range counts {
		for line, n := range dups {
			if n > 1 {
				fmt.Printf("%s: %d\t%s\n", file, n, line)
			}
		}
	}
}

func countLines(f *os.File, filesLineCount map[string]map[string]int) {
	input := bufio.NewScanner(f)

	var counts map[string]int = filesLineCount[f.Name()]
	if filesLineCount[f.Name()] == nil {
		counts = make(map[string]int)
		filesLineCount[f.Name()] = counts
	}

	for input.Scan() {
		counts[input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err()
}

//!-
