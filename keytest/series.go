package keytest

// seriesAppropriateValues defines the appropriate values for zero and one sequences
// at different lengths according to the Series test.
var seriesAppropriateValues map[int][]int = map[int][]int{
	1: {2267, 2733},
	2: {1079, 1421},
	3: {502, 748},
	4: {223, 402},
	5: {90, 223},
	6: {90, 223},
}

// Series performs the Series test on a binary string.
// It checks whether the counts of zero and one sequences of different lengths
// fall within the specified bounds defined by FIPS 140-3.
func Series(binary string) (bool, error) {
	// Validate the binary string.
	if err := validateBinary(binary); err != nil {
		return false, err
	}

	// Variables for tracking sequence counts.
	var (
		prevBit               = '1' // Used to compare current bits to previous bits – to determine if a sequence ended.
		currentSequenceLength = 1   // Current sequence length counter.
		zeroSequences         = make([]int, 6)
		oneSequences          = make([]int, 6)

		// Function to increment the count of zero or one sequences based on the current bit.
		incrementCount = func(bit rune) {
			if bit == '1' {
				oneSequences[min(currentSequenceLength-1, 5)]++
			} else {
				zeroSequences[min(currentSequenceLength-1, 5)]++
			}
		}
	)

	// Iterate through the binary string to count sequences.
	for i := 1; i < len(binary); i++ {
		currbit := rune(binary[i])
		if currbit == prevBit {
			currentSequenceLength++
		} else {
			incrementCount(prevBit)
			currentSequenceLength = 1
			prevBit = currbit
		}
	}
	incrementCount(prevBit)

	// Check if the counts of zero and one sequences are within specified bounds.
	for i := 0; i < len(zeroSequences); i++ {
		zeroCount := zeroSequences[i]
		oneCount := oneSequences[i]
		lowerBound := seriesAppropriateValues[i+1][0]
		upperBound := seriesAppropriateValues[i+1][1]
		if zeroCount < lowerBound || zeroCount > upperBound || oneCount < lowerBound || oneCount > upperBound {
			return false, nil
		}
	}
	return true, nil
}
