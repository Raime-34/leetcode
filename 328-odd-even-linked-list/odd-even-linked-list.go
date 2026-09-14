func oddEvenList(head *ListNode) *ListNode {
	lastOdd := head

	currentNode := head
	currentIndex := 0
	var prevNode *ListNode
	for {
		if currentNode == nil {
			break
		}

		if currentIndex%2 == 0 && currentIndex != 0 {
			prevNode.Next = currentNode.Next
			currentNode.Next = lastOdd.Next
			lastOdd.Next = currentNode
			lastOdd = currentNode
			currentNode = prevNode
		}

		prevNode = currentNode
		currentNode = currentNode.Next
		currentIndex++
	}

	return head
}