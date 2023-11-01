# FIPS 140-3 implementation in Go
Covers the following tests:
- Monobit
- Maximum series length
- Poker
- Series length

> Each test takes 20000 bits as an input.

## Work done
- `keytest/monobit.go` - Monobit test implementation.
- `keytest/longestsequence.go` - Maximum Series test implementation.
- `keytest/poker.go` - Poker test implementation.
- `keytest/series.go` - Series test implementation.
- `keytest/keytest.go` - contains a function to run all the tests.
- `keytest/util.go` - contains utility methods.
- `example.go` - a simple example of how to use the tests.

## How to use
The following block of code can be found in `example.go`:
```golang
import (
	"fmt"
	"math/rand"

	"github.com/danielost/fips-140-3/keytest" // This import is needed to run the tests.
)

func foo() {
	// Most likely, you will have a real sequence of bits,
	// but fot the demostration we just generate a random binary value.
	length := 20000
	runeArray := make([]rune, length)
	for i := 0; i < length; i++ {
		runeArray[i] = rune('0' + rand.Intn(2))
	}
	binary := string(runeArray)

	// Run all the tests.
	result, err := keytest.RunAll(binary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
```

If the sequence is considered random, the following output will be produced:

```bash
Starting FIPS 140-3 tests
===================
FIPS 140-3 success:
===================
FIPS 140-3 Monobit: true
FIPS 140-3 Longest Sequence: true
FIPS 140-3 Poker: true
FIPS 140-3 Series Count: true
```
