package keytest

import (
	"fmt"

	bignumbers "github.com/danielost/big-numbers/src"
)

func getBinary(key string) (string, error) {
	bigNumber := bignumbers.BigNumber{}
	err := bigNumber.SetHex(key)
	if err != nil {
		return "", err
	}
	return bigNumber.GetBinary(), nil
}

func validateKeyLength(binary string) error {
	keyLength := len(binary)
	if keyLength == 20000 {
		return nil
	}
	return fmt.Errorf("expected key lenght to be 20000, but got %d", keyLength)
}
