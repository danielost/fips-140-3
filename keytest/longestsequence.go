package keytest

// threshold defines the maximum allowed length for the longest sequence in the LongestSequence test.
const threshold int = 36

// LongestSequence performs the LongestSequence test on a binary string.
// It checks whether the length of the longest sequence of consecutive bits is less than the specified threshold.
func LongestSequence(binary string) (bool, error) {
	// Validate the binary string.
	if err := validateBinary(binary); err != nil {
		return false, err
	}

	// Variables for tracking sequence lengths.
	var (
		prevBit               = '1' // Used to compare current bits to previous bits – to determine if a sequence ended.
		currentSequenceLength = 1   // Current sequence length counter.
		longestSequenceLength = 1
	)

	// Iterate through the binary string to find the longest sequence.
	for i := 1; i < len(binary); i++ {
		currbit := rune(binary[i])
		if currbit == prevBit {
			currentSequenceLength++
		} else {
			// Update the longest sequence length if the current sequence is longer.
			longestSequenceLength = max(longestSequenceLength, currentSequenceLength)
			currentSequenceLength = 1
			prevBit = currbit
		}
	}
	// Update the longest sequence length for the last sequence.
	longestSequenceLength = max(longestSequenceLength, currentSequenceLength)

	// Check if the length of the longest sequence is less than the specified threshold.
	return longestSequenceLength < threshold, nil
}
