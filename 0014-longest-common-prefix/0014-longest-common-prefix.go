func longestCommonPrefix(strs []string) string {
    str := ""
	minLength := len(strs[0])
	for _, s := range strs{
		minLength = min(minLength, len(s))
	}
	Outerloop:
		for i := 0; i < minLength; i++{ //номер буквы
			for j := 0; j < len(strs) -1; j++{ //номер слова
				if strs[j][i] != strs[j+1][i]{
					break Outerloop
				}
			}
			str += string(strs[0][i])
		}
		return str
}