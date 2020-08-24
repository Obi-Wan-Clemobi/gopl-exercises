// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Exercise 1.2: Modify th4e echo program to also print the index and value of each of its arguments, one per line
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args {
		fmt.Println(index, ":", value)
	}
}

//!-
