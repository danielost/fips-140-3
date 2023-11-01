package keytest

type TestData struct {
	binary        string
	zeroBitCount  int
	oneBitCount   int
	zeroSequences map[int]int
	oneSequences  map[int]int
	poker         float64
}

func InitTestData(binary string) TestData {
	return TestData{
		binary:        binary,
		zeroSequences: make(map[int]int),
		oneSequences:  make(map[int]int),
	}
}
