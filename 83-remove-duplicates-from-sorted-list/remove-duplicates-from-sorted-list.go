func deleteDuplicates(head *ListNode) *ListNode {
	nodes := make(map[int]struct{}, 0)

	if head == nil {
		return head
	}

	for {
		nodes[head.Val] = struct{}{}

		if head.Next == nil {
			break
		}

		head = head.Next
	}

	nodesAsSlice := []int{}
	for val := range nodes {
		nodesAsSlice = append(nodesAsSlice, val)
	}
	slices.Sort(nodesAsSlice)

	var tail *ListNode
	for i := len(nodesAsSlice) - 1; i >= 0; i-- {
		newHead := &ListNode{
			Val:  nodesAsSlice[i],
			Next: tail,
		}

		tail = newHead
	}

	return tail
}