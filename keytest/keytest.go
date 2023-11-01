package keytest

import (
	"fmt"
	"math"
)

// Runs all the FIPS tests for the input key (hex).
func Run(key string) (bool, error) {
	binary, err := getBinary(key)
	if err != nil {
		return false, err
	}
	if len(binary) != 20000 {
		return false, fmt.Errorf("expected key lenght to be 20000, but got %d", len(binary))
	}
	fmt.Printf("Starting FIPS 140-3 tests\n")
	testData := calculateTestData(binary)
	return testData.analyzeData(), nil
}

// Based on the input binary value,
// returns a TestData object filled
// with the tests results.
func calculateTestData(binary string) TestData {
	fmt.Printf("Bits received from input: %d\n", len(binary))
	td := InitTestData(binary)

	// Data for the Poker test.
	var (
		m = 4                    // Poker block size.
		k = len(binary) / m      // Number of Poker blocks.
		n = make(map[string]int) // Counts the number of times each block appears.
	)

	// Data for the series test.
	var (
		prevBit               = '1' // Used to compare current bits to previos bits – to determine if a sequence ended.
		currentSequenceLength = 1   // Current sequence length counter.
	)

	for i, bit := range td.binary {
		// Update the monobit data.
		td.updateMonobit(bit)

		// Update the series data.
		if bit == prevBit {
			currentSequenceLength++
		} else if i > 0 {
			td.updateSeries(prevBit, currentSequenceLength)
			currentSequenceLength = 1
			prevBit = bit
		}

		// Update the Poker data.
		if i%m == 0 {
			td.updatePokerBlockCount(i, m, n)
		}
	}

	td.updateSeries(prevBit, currentSequenceLength)
	td.calculatePoker(m, k, n)
	return td
}

func (td *TestData) updateMonobit(bit rune) {
	if bit == '0' {
		td.zeroBitCount++
	} else {
		td.oneBitCount++
	}
}

func (td *TestData) updateSeries(bit rune, sequenceLength int) {
	if bit == '1' {
		td.oneSequences[sequenceLength]++
	} else {
		td.zeroSequences[sequenceLength]++
	}
}

func (td *TestData) updatePokerBlockCount(i, m int, n map[string]int) {
	block := td.binary[i : i+m]
	n[block]++
}

func (td *TestData) calculatePoker(m, k int, n map[string]int) {
	// Calculate the sum for the Poker value.
	sum := 0.0
	for _, count := range n {
		sum += math.Pow(float64(count), 2)
	}

	// Calculate the Poker value.
	td.poker = math.Pow(2, float64(m))/float64(k)*sum - float64(k)
}

// Analyzes the data inside TestData.
func (td *TestData) analyzeData() bool {
	// Monobit test.
	monobitPassed := td.oneBitCount > 9654 && td.oneBitCount < 10346

	// Series test.
	sequenceAppropriateValues := map[int][]int{
		1: {2267, 2733},
		2: {1079, 1421},
		3: {502, 748},
		4: {223, 402},
		5: {90, 223},
		6: {90, 223},
	}
	zeroSequenceActualValues := make([]int, 6)
	oneSequenceActualValues := make([]int, 6)
	longestSequence := 0

	countSequences(td.oneSequences, &longestSequence, oneSequenceActualValues)
	countSequences(td.zeroSequences, &longestSequence, zeroSequenceActualValues)

	longestSequencePassed := longestSequence <= 36
	sequenceLengthPassed := true
	for i := 0; sequenceLengthPassed && i < len(zeroSequenceActualValues); i++ {
		zeroValue := zeroSequenceActualValues[i]
		oneValue := oneSequenceActualValues[i]
		zeroPassed := sequenceAppropriateValues[i+1][0] <= zeroValue && sequenceAppropriateValues[i+1][1] >= zeroValue
		onePassed := sequenceAppropriateValues[i+1][0] <= oneValue && sequenceAppropriateValues[i+1][1] >= oneValue
		sequenceLengthPassed = sequenceLengthPassed && zeroPassed && onePassed
	}
	pokerPassed := td.poker < 57.4 && td.poker > 1.03

	passed := monobitPassed && longestSequencePassed && sequenceLengthPassed && pokerPassed
	if passed {
		fmt.Printf("===================\nFIPS 140-3 success:\n===================\n")
	} else {
		fmt.Printf("===================\nFIPS 140-3 failure:\n===================\n")
	}

	printTestResult(monobitPassed, "Monobit", fmt.Sprintf("%d ones and %d zeros.", td.oneBitCount, td.zeroBitCount))
	printTestResult(longestSequencePassed, "Longest Sequence", fmt.Sprintf("Longest sequence is %d bits long.", longestSequence))
	printTestResult(pokerPassed, "Poker", fmt.Sprintf("Poker value is %0.1f.", td.poker))
	printTestResult(sequenceLengthPassed, "Sequence Count", fmt.Sprintf("Sequences of ones:\t%v\nSequences of zeros:\t%v", oneSequenceActualValues, zeroSequenceActualValues))
	return passed
}

func printTestResult(passed bool, name, details string) {
	fmt.Printf("FIPS 140-3 %s: %t\n%s\n\n", name, passed, details)
}

// Counts the number each suquence occurs.
// Also updates the longest sequence value.
func countSequences(sequence map[int]int, longestSequence *int, actualValues []int) {
	for sequenceLength, sequenceCount := range sequence {
		*longestSequence = max(*longestSequence, sequenceLength)
		actualValues[min(sequenceLength-1, 5)] += sequenceCount
	}
}
