package keytest

import "fmt"

func validateBinary(binary string) error {
	length := len(binary)
	if length != 20000 {
		return fmt.Errorf("expected the input to be 20000 bits long, but got %d", length)
	}
	for _, bit := range binary {
		if bit != '0' && bit != '1' {
			return fmt.Errorf("binary string contains prohibited symbols")
		}
	}
	return nil
}
