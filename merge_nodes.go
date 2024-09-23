package algorithms

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes2(head *ListNode) *ListNode {
	var firstZeroPrev *ListNode
	currentNode := head.Next
	var secondZero *ListNode
	var sumSoFar int
	for currentNode != nil {
		if currentNode.Val == 0 {
			secondZero = currentNode
			if firstZeroPrev == nil {
				head = &ListNode{
					Val: sumSoFar,
				}
				firstZeroPrev = head
				//  If second zero is the tail of the linklist
				if secondZero.Next == nil {
					head.Next = nil
				} else {
					head.Next = secondZero

					secondZero = nil
				}
			} else {
				firstZeroPrev.Next = &ListNode{
					Val: sumSoFar,
				}
				firstZeroPrev = firstZeroPrev.Next
				if secondZero.Next == nil {
					firstZeroPrev.Next = nil
				} else {

					secondZero = nil
				}
			}
			sumSoFar = 0
		}
		sumSoFar += currentNode.Val
		currentNode = currentNode.Next
	}

	return head
}

func mergeNodes(head *ListNode) *ListNode {
	currentZero := head
	var lastZero *ListNode

	currentNode := head.Next

	for currentNode != nil {
		if currentNode.Val != 0 {
			currentZero.Val += currentNode.Val
		} else {
			currentZero.Next = currentNode
			lastZero = currentZero
			currentZero = currentNode
		}
		currentNode = currentNode.Next
	}
	lastZero.Next = nil

	return head
}

//func NodesMerge(firstZeroPrev, firstZero, SecondZero, head *ListNode, sumSoFar int) *ListNode {
//	//  We are merging between head and the next zero
//	if firstZeroPrev == nil {
//		head = &ListNode{
//			Val: sumSoFar,
//		}
//		firstZeroPrev = head
//		//  If second zero is the tail of the linklist
//		if SecondZero.Next == nil {
//			head.Next = nil
//		} else {
//			head.Next = SecondZero
//			firstZero = SecondZero
//			SecondZero = nil
//		}
//	} else {
//		firstZeroPrev.Next = &ListNode{
//			Val: sumSoFar,
//		}
//		firstZeroPrev = firstZeroPrev.Next
//		if SecondZero.Next == nil {
//			firstZeroPrev.Next = nil
//		} else {
//			firstZero = SecondZero
//			SecondZero = nil
//		}
//	}
//	return head
//}
