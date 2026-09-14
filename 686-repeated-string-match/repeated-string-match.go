func repeatedStringMatch(a string, b string) int {
	repeatTime := len(b) / len(a)
	var builder strings.Builder

	for range repeatTime {
		builder.WriteString(a)
	}

	if strings.Contains(builder.String(), b) {
		return repeatTime
	} else {
		for range 2 {
			repeatTime++
			builder.WriteString(a)
			if strings.Contains(builder.String(), b) {
				return repeatTime
			}
		}
		return -1
	}
}