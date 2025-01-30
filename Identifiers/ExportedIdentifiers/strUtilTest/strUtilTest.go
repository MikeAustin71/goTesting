package strUtilTest


func ReverseAString(s string) string {

	return invisibleReverseAString(s)
}

/*
	Taken from Todd McLeod's reverseTwo.go
	https://github.com/GoesToEleven/GolangTraining/tree/master/02_package/stringutil
*/
func invisibleReverseAString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i + 1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
