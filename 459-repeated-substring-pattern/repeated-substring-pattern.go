func repeatedSubstringPattern(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		subStr := s[:i]
		l := len(subStr)
		for j := i; j < len(s); j += l {
			if j+l > len(s) {
				break
			}
			if subStr != s[j:j+l] {
				break
			}
			if j == len(s)-l {
				return true
			}
		}
	}

	return false
}
