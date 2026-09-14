func copyRandomList(head *Node) *Node {
	var (
		newHead     = &Node{}
		currentNode = newHead
	)

	indisies := make(map[*Node]*Node)
	later := make(map[*Node][]*Node)

	if head == nil {
		return nil
	}

	for i := 0; ; i++ {
		currentNode.Val = head.Val
		indisies[head] = currentNode
		if ptr, ok := indisies[head.Random]; ok {
			currentNode.Random = ptr
		} else {
			if later[head.Random] == nil {
				later[head.Random] = []*Node{}
			}
			later[head.Random] = append(later[head.Random], currentNode)
		}
		if ptrs, ok := later[head]; ok {
			for _, ptr := range ptrs {
				ptr.Random = currentNode
			}
		}

		if head.Next == nil {
			break
		}

		head = head.Next
		nextNode := &Node{}
		currentNode.Next = nextNode
		currentNode = nextNode
	}

	return newHead
}
