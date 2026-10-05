func lengthOfLastWord(s string) int {
	new_s := strings.Fields(s)
	return len(new_s[len(new_s)-1])
}