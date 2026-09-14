var (
	hidden      = "*****"
	hiddenPhone = "***-***-"
	re          = regexp.MustCompile(`[-()+\s]`)
)

func maskPII(s string) string {
	var res string
	var builder strings.Builder

	if i := strings.Index(s, "@"); i > 0 {
		builder.WriteByte(s[0])
		builder.WriteString(hidden)
		builder.WriteString(s[i-1:])
		res = strings.ToLower(builder.String())
	} else {
		s = string(re.ReplaceAll([]byte(s), []byte("")))
		if i := len(s) - 10; i > 0 {
			if i > 3 {
				i = 3
			}

			builder.WriteByte('+')
			for range i {
				builder.WriteByte('*')
			}
			builder.WriteByte('-')
		}
		builder.WriteString(hiddenPhone)
		builder.WriteString(s[len(s)-4:])
		res = builder.String()
	}

	return res
}
