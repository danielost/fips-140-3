package keytest

// Constants defining the monobit test bounds.
const (
	monobitLowerBound = 9654
	monobitUpperBound = 10346
)

// Monobit performs the monobit test on a binary string.
// It checks whether the number of zeros and ones in the binary string fall within the specified bounds.
func Monobit(binary string) (bool, error) {
	// Validate the binary string.
	if err := validateBinary(binary); err != nil {
		return false, err
	}

	// Count the number of zeros in the binary string.
	zeros := 0
	for _, bit := range binary {
		if bit == '0' {
			zeros++
		}
	}

	// Check if the number of zeros is within the specified bounds.
	return zeros > monobitLowerBound && zeros < monobitUpperBound, nil
}
