var r, _ = regexp.Compile("(^[A-Z]?)([A-Z]*|[a-z]*)$")

func detectCapitalUse(word string) bool {
	return r.MatchString(word)
}