package keytest

import (
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
