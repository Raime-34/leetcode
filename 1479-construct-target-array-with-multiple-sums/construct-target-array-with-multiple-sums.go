type hp []int

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *hp) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *hp) Push(i any) {
	*h = append(*h, i.(int))
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) onlyOnes() bool {
	return !slices.ContainsFunc(*h, func(n int) bool {
		return n > 1
	})
}

func (h *hp) sum() int {
	sum := 0
	for i := 0; i < h.Len(); i++ {
		sum += (*h)[i]
	}
	return sum
}

func isPossible(target []int) bool {
	if len(target) <= 1 {
		return slices.Contains(target, 1)
	}

	h := hp(target)

	heap.Init(&h)
	for !h.onlyOnes() {
		next := heap.Pop(&h).(int)
		sum := h.sum()
		diff := next - sum

		if diff > 100_000 {
			mod := next % sum
			if sum == 1 {
				heap.Push(&h, 1)
				continue
			}
			if mod == 0 && sum > 1 {
				return false
			}
			if mod != 0 && sum > 1 {
				heap.Push(&h, mod)
				continue
			}
		}

		if diff <= 0 {
			return false
		}

		if diff > 0 {
			heap.Push(&h, diff)
		}
	}

	return true
}
