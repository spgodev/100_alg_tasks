package recursion

// ReverseString рекурсивно переворачивает строку.
// ReverseString("go") = "og"
// ReverseString("hello") = "olleh"
func ReverseString(s string) string {
	// TODO
	if len(s) <= 1 {
		return s
	}
	return s[len(s)-1:] + ReverseString(s[:len(s)-1])
}
