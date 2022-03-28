package iteration

const repeateCount = 5

func Repeat(str string) string {
	var repeated string
	for i := 0; i < repeateCount; i++ {
		repeated += str
	}
	return repeated
}
