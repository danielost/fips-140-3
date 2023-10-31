package keytest

import (
	"fmt"
	"math"
)

func Run(key string) (bool, error) {
	fmt.Printf("Starting FIPS 140-3 tests\n")
	testData, err := calculateTestData(key)
	if err != nil {
		return false, err
	}
	return testData.analyzeData(), nil
}

func calculateTestData(key string) (td TestData, err error) {
	binary, err := getBinary(key)
	if err != nil {
		return td, err
	}
	if err := validateKeyLength(binary); err != nil {
		return td, err
	}
	fmt.Printf("Bits received from input: %d\n", len(binary))

	td.hex = key
	td.binary = binary
	td.zeroSequences = make(map[int]int)
	td.oneSequences = make(map[int]int)

	// Data for the Poker test
	var (
		m = 4
		Y = len(td.binary)
		k = Y / m
		n = make(map[string]int)
	)

	// Data for the monobit and series tests
	var (
		prevBit               = '1'
		currentSequenceLength = 1
	)

	for i, bit := range td.binary {
		// Update the monobit data
		if bit == '0' {
			td.zeroBitCount++
		} else {
			td.oneBitCount++
		}

		// Update the sequence data
		if bit == prevBit {
			currentSequenceLength++
		} else if i > 0 {
			if prevBit == '1' {
				td.oneSequences[currentSequenceLength]++
			} else {
				td.zeroSequences[currentSequenceLength]++
			}
			currentSequenceLength = 1
			prevBit = bit
		}

		// Update the Poker data
		if i%m == 0 {
			block := td.binary[i : i+m]
			n[block]++
		}
	}
	if prevBit == '1' {
		td.oneSequences[currentSequenceLength]++
	} else {
		td.zeroSequences[currentSequenceLength]++
	}

	sum := 0.0

	for _, count := range n {
		for i := 0; i < count; i++ {
			sum += math.Pow(float64(count), 2)
		}
	}

	td.poker = (math.Pow(2, float64(m)) / float64(k) * sum) - float64(k)
	return td, nil
}

func (td *TestData) analyzeData() bool {
	monobitPassed := td.oneBitCount > 9654 && td.oneBitCount < 10346
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
	for i, v := range zeroSequenceActualValues {
		sequenceLengthPassed = sequenceLengthPassed && sequenceAppropriateValues[i+1][0] <= v && sequenceAppropriateValues[i+1][1] >= v
	}
	for i, v := range oneSequenceActualValues {
		sequenceLengthPassed = sequenceLengthPassed && sequenceAppropriateValues[i+1][0] <= v && sequenceAppropriateValues[i+1][1] >= v
	}
	pokerPassed := td.poker < 57.4 && td.poker > 1.03

	passed := monobitPassed && longestSequencePassed && sequenceLengthPassed && pokerPassed
	if passed {
		fmt.Printf("==================\nFIPS 140-3 success\n==================\n")
	} else {
		fmt.Printf("==================\nFIPS 140-3 failure\n==================\n")
	}

	printTestResult(monobitPassed, "Monobit", fmt.Sprintf("%d ones and %d zeros. Both values must be in range (9654, 10346).", td.oneBitCount, td.zeroBitCount))
	printTestResult(longestSequencePassed, "Longest Sequence", fmt.Sprintf("Longest sequence is %d bits long. Must be less that or equal to 36.", longestSequence))
	printTestResult(pokerPassed, "Poker", fmt.Sprintf("Poker value is %0.1f. Must be in range (1.03, 57.4).", td.poker))
	printTestResult(sequenceLengthPassed, "Sequence Count", fmt.Sprintf("Sequences of ones:\t%v\nSequences of zeros:\t%v\nMust be [1:2267-2733, 2:1079-1421, 3:502-748, 4:223-402; 5:90-223, 6+:90-223].", oneSequenceActualValues, zeroSequenceActualValues))
	return passed
}

func printTestResult(passed bool, name, details string) {
	fmt.Printf("FIPS 140-3 %s: %t\n%s\n\n", name, passed, details)
}

func countSequences(sequence map[int]int, longestSequence *int, actualValues []int) {
	for sequenceLength, sequenceCount := range sequence {
		*longestSequence = max(*longestSequence, sequenceLength)
		actualValues[min(sequenceLength-1, 5)] += sequenceCount
	}
}
