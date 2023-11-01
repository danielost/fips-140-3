package main

import (
	"fmt"
	"math/rand"

	"github.com/danielost/fips-140-3/keytest"
)

func main() {
	length := 20000
	runeArray := make([]rune, length)

	for i := 0; i < length; i++ {
		runeArray[i] = rune('0' + rand.Intn(2))
	}

	binary := string(runeArray)

	fmt.Println(keytest.RunAll(binary))
}
