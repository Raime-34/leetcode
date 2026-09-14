func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == nil {
		return head
	}

	var prevprevNode *ListNode
	prevNode := head
	currentNode := head.Next

	for {
		nextNode := currentNode.Next

		prevNode.Next = prevprevNode
		currentNode.Next = prevNode

		if nextNode == nil {
			break
		}

		prevprevNode = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}

	return currentNode
}
