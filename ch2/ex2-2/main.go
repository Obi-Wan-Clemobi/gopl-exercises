// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
// Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf that reads
// numbers from its command-line arguemnts or from the standard input if there are no arguments,
// and converts each number into units like temperature in Celsius and Fhrenheit, length in
// feet and meters, weight in poinds and kilograms, and the like
package main

import (
	"bufio"
	"fmt"
	"gopl-exercise/ch2/ex2-2/lengthconv"
	"gopl-exercise/ch2/ex2-2/massconv"
	"os"
	"regexp"
	"strconv"
)

// Parsed arg value with unit
type ParsedValue struct {
	Value float64
	Unit  string
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			exec(arg)
		}
	} else {
		readStdin()
	}
}

func exec(arg string) {
	value, err := parse(arg)
	if err != nil {
		fmt.Println(err)
	} else {
		convert(value)
	}
}

func readStdin() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		exec(scanner.Text())
	}
}

func convert(parsedValue ParsedValue) {
	value := parsedValue.Value
	unit := parsedValue.Unit

	if unit == "m" {
		fmt.Printf("%g%s = %s\n", value, unit, lengthconv.MToFt(lengthconv.Meter(value)).String())
	} else if unit == "ft" {
		fmt.Printf("%g%s = %s\n", value, unit, lengthconv.FtToM(lengthconv.Feet(value)).String())
	} else if unit == "lb" {
		fmt.Printf("%g%s = %s\n", value, unit, massconv.LbsToKgs(massconv.Pound(value)).String())
	} else if unit == "kg" {
		fmt.Printf("%g%s = %s\n", value, unit, massconv.KgToPound(massconv.Kilogram(value)).String())
	} else {
		fmt.Printf("wtf?\n")
	}
}

func parse(arg string) (ParsedValue, error) {
	regex := regexp.MustCompile("^([-+]?\\d*\\.?\\d+)\\s?(\\w+)$")
	result := regex.FindStringSubmatch(arg)
	if len(result) < 2 {
		return ParsedValue{0, ""}, fmt.Errorf("Invalid input: %s", arg)
	}

	value, err := strconv.ParseFloat(result[1], 64)
	unit := result[2]

	if err != nil {
		return ParsedValue{0, ""}, err
	}

	return ParsedValue{value, unit}, nil
}

//!-
