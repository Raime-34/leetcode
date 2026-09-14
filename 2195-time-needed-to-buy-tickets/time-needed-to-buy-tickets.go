func timeRequiredToBuy(tickets []int, k int) int {
	seconds := 0

	for {
		if len(tickets) == 1 {
			seconds += tickets[0]
			break
		}

		seconds++
		nextTicket := tickets[0] - 1
		tickets = tickets[1:]

		if k == 0 && nextTicket == 0 {
			break
		}

		k--

		if nextTicket > 0 {
			tickets = append(tickets, nextTicket)
		}

		if k < 0 {
			k = len(tickets) - 1
		}
	}

	return seconds
}