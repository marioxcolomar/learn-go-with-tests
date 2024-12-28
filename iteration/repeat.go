package iteration

const repeatCount = 5

func Repeat(input string, times int) string {
	count := times
	if count <= 0 {
		count = repeatCount
	}
	var repeated string
	for i := 0; i < count; i++ {
		repeated += input
	}
	return repeated
}
