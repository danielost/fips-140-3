package keytest

import (
	"fmt"
)

func Run(key string) (bool, error) {
	binary := getBinary(key)

	// var (
	// 	zeroBitCount    = 0
	// 	oneBitCount     = 0
	// 	longestSequence = 0
	// )

	// sequenceCount := make(map[uint]uint)

	fmt.Println(binary)

	return true, nil
}
