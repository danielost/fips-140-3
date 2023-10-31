package keytest

type TestData struct {
	hex           string
	binary        string
	zeroBitCount  int
	oneBitCount   int
	zeroSequences map[int]int
	oneSequences  map[int]int
	poker         float64
}
