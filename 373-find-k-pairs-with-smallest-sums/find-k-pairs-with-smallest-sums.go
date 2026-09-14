type item struct {
	sum int
	i   int
	j   int
}

type myHeap []item

func composeMyheap(nums1 []int, nums2 []int) myHeap {
	h := make(myHeap, 0, len(nums1))

	for i := range nums1 {
		h = append(h, item{
			i:   i,
			j:   0,
			sum: nums1[i] + nums2[0],
		})
	}

	return h
}

func (h myHeap) Len() int {
	return len(h)
}

func (h myHeap) Less(i, j int) bool {
	if h[i].sum == h[j].sum {
		if h[i].i != h[j].i {
			return h[i].i < h[j].i
		} else {
			return h[i].j < h[j].j
		}
	}
	return h[i].sum < h[j].sum
}

func (h myHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *myHeap) Push(i any) {
	*h = append(*h, i.(item))
}

func (h *myHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n := composeMyheap(nums1, nums2)
	heap.Init(&n)

	result := make([][]int, 0, k)
	for len(result) < k {
		nextMin := heap.Pop(&n).(item)
		result = append(result, []int{nums1[nextMin.i], nums2[nextMin.j]})

		if nextMin.j+1 < len(nums2) {
			i := item{
				i:   nextMin.i,
				j:   nextMin.j + 1,
				sum: nums1[nextMin.i] + nums2[nextMin.j+1],
			}
			heap.Push(&n, i)
		}
	}
	return result
}