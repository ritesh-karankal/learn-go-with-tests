package iteration

func Repeat(char string, iterate int) string {
	var repeated string
	for i := 0; i < iterate; i++ {
		repeated += char
	}
	return repeated
}
