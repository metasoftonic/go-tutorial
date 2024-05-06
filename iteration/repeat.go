package iteration

func Repeat(character string, repeatCount int) string {
	var expected string
	for i := 0; i < repeatCount; i++ {
		expected += character
	}
	return expected
}
