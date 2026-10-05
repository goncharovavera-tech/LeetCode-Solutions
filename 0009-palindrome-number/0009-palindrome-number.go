func isPalindrome(x int) bool {
	chars := []rune(fmt.Sprint(x))
	length := len(chars)
	chars_r := make([]rune, length)
	for index, element := range chars{
		chars_r[length - index - 1] = element
	}
	switch{
	case slices.Equal(chars, chars_r):
		return true
	default: return false
	}
}