package keytest

import (
	"math"
)

// Constants defining the poker test bounds.
const (
	pokerLowerBound = 1.03
	pokerUpperBound = 57.4
)

// Poker performs the poker test on a binary string.
// It checks whether the sequences of length m each appear approximately the same number of times
// as expected in a random sequence.
func Poker(binary string) (bool, error) {
	// Validate the binary string.
	if err := validateBinary(binary); err != nil {
		return false, err
	}

	// Poker test parameters.
	var (
		m = 4                    // Poker block size.
		k = len(binary) / m      // Number of Poker blocks.
		n = make(map[string]int) // Counts the number of times each block appears.
	)

	// Count occurrences of each Poker block.
	for i := 0; i < k; i++ {
		block := binary[i*m : (i+1)*m]
		n[block]++
	}

	// Calculate the test statistic X3.
	sum := 0.0
	for _, count := range n {
		sum += float64(count * count)
	}
	X3 := math.Pow(2, float64(m))/float64(k)*sum - float64(k)

	// Check if X3 is within the specified bounds.
	return X3 > pokerLowerBound && X3 < pokerUpperBound, nil
}
