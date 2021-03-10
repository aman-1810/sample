package strings

// Reverse a string
/*
	Since strings in Go are immutable, we first convert the string to a mutable array of runes ([]rune),
	perform the reverse operation on that, and then re-cast to a string.
*/
func Reverse(s string) string {
	runes := []rune(s)
	reversedRunes := reverseRunes(runes)
	return string(reversedRunes)
}

// Reverses an array of runes
// This function is not exported (It is only visible inside the `strings` package)
func reverseRunes(r []rune) []rune {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}
