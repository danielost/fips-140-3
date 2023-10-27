package keytest

import bignumbers "github.com/danielost/big-numbers/src"

func getBinary(key string) string {
	bigNumber := bignumbers.BigNumber{}
	bigNumber.SetHex(key)
	return bigNumber.GetBinary()
}
