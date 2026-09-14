func countStudents(students []int, sandwiches []int) int {
	for i := 0; ; i++ {
		if len(students) == 0 || len(sandwiches) == 0 {
			break
		}

		nextStudent := students[0]
		nextSandwitch := sandwiches[0]

		if !slices.Contains(students, nextSandwitch) {
			break
		}

		students = students[1:]
		if nextSandwitch == nextStudent {
			sandwiches = sandwiches[1:]
		} else {
			students = append(students, nextStudent)
		}
	}

	return len(students)
}