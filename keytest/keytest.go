package keytest

import (
	"fmt"
)

// RunAll runs all the FIPS 140-3 tests for the input binary value.
// It prints the results of each test and returns a boolean indicating overall success.
func RunAll(binary string) (bool, error) {
	// Validate the binary string.
	if err := validateBinary(binary); err != nil {
		return false, err
	}

	fmt.Printf("Starting FIPS 140-3 tests\n")

	// Run individual FIPS 140-3 tests.
	monobit, _ := Monobit(binary)
	longestSequence, _ := LongestSequence(binary)
	poker, _ := Poker(binary)
	series, _ := Series(binary)

	// Check if all tests passed.
	passed := monobit && longestSequence && poker && series

	// Print overall result header.
	if passed {
		fmt.Printf("===================\nFIPS 140-3 success:\n===================\n")
	} else {
		fmt.Printf("===================\nFIPS 140-3 failure:\n===================\n")
	}

	// Print individual test results.
	printTestResult(monobit, "Monobit")
	printTestResult(longestSequence, "Longest Sequence")
	printTestResult(poker, "Poker")
	printTestResult(series, "Series Count")

	return passed, nil
}

func printTestResult(passed bool, name string) {
	fmt.Printf("FIPS 140-3 %s: %t\n", name, passed)
}
