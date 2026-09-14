func licenseKeyFormatting(s string, k int) string {
	var builder strings.Builder
	amount := 0

	runes := []rune(s)

	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == '-' {
			continue
		}

		builder.WriteByte(byte(runes[i]))
		amount++

		if amount == k {
			builder.WriteString("-")
			amount = 0
		}
	}

	return strings.ToUpper(strings.Trim(Reverse(builder.String()), "-"))
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}