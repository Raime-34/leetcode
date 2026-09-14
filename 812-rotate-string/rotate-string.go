func rotateString(s string, goal string) bool {
	if len(s) == 1 {
		return s == goal
	}

	if s == goal {
		return true
	}

	var anchorString string
	anchorLength := 0
	for anchorString == "" {
		anchorLength++
		notUnique := make(map[string]bool)
		for i := anchorLength; i <= len(s); i++ {
			subStr := s[i-anchorLength : i]
			if _, ok := notUnique[subStr]; ok {
				notUnique[subStr] = true
			} else {
				notUnique[subStr] = false
			}
		}

		for subStr, isNotUnique := range notUnique {
			if !isNotUnique {
				anchorString = subStr
				break
			}
		}
	}

	anchorS := strings.Index(s, anchorString)
	anchorGoal := strings.Index(goal, anchorString)
	if anchorGoal == -1 {
		return false
	}

	recomposedS := s[anchorS:] + s[:anchorS]
	recomposedGoal := goal[anchorGoal:] + goal[:anchorGoal]

	return recomposedS == recomposedGoal
}
