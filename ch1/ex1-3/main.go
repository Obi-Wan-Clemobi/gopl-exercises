// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Exercise 1.3: Experiment to measure the difference in runing time between our potentially inefficient versions and the ones that uses
// strings.Join.
package main

import (
	"fmt"
	"strings"
	"time"
)

const size = 10000

func main() {

	fakeValues := make([]string, size)
	for i := 0; i < size; i++ {
		fakeValues[i] = string(i)
	}

	time1 := inefficientVersion(fakeValues)
	time2 := funcUsingJoin(fakeValues)

	fmt.Printf("v1 took: %v\n", time1)
	fmt.Printf("v2 took: %v\n", time2)

	fmt.Printf("inefficientVersion(v1) to funcUsingJoin(v2) delta: %v", (time1 - time2))
}

func inefficientVersion(fakeValues []string) int64 {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range fakeValues[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	return time.Since(start).Nanoseconds()
}

func funcUsingJoin(fakeValues []string) int64 {
	start := time.Now()
	fmt.Println(strings.Join(fakeValues[1:], " "))
	return time.Since(start).Nanoseconds()
}

//!-
