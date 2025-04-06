package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var (
		newList *ListNode

		listArr []int = []int{}
	)

	// looping through the linked lists and adding the values to the listArr
	for list1 != nil {
		listArr = append(listArr, list1.Val)
		list1 = list1.Next
	}

	for list2 != nil {
		listArr = append(listArr, list2.Val)
		list2 = list2.Next
	}

	// edge cases
	if len(listArr) == 0 {
		return newList
	}

	// sorting the listArr
	for i := 0; i < len(listArr); i++ {
		for j := i + 1; j < len(listArr); j++ {
			if listArr[i] > listArr[j] {
				listArr[i], listArr[j] = listArr[j], listArr[i]
			}
		}
	}

	// creating the new linked list
	var (
		head *ListNode
	)

	for _, val := range listArr {
		if newList == nil {
			newList = &ListNode{Val: val}
			head = newList
		} else {
			newList.Next = &ListNode{Val: val}
			newList = newList.Next
		}
	}
	// returning the new linked list
	if newList != nil {
		return head
	}
	return nil
}

func TestMergeTwoLists() {
	var (
		list1 = &ListNode{Val: 1}
		list2 = &ListNode{Val: 2}
	)

	list1.Next = &ListNode{Val: 4}
	list2.Next = &ListNode{Val: 3}

	mergedList := MergeTwoLists(list1, list2)

	for mergedList != nil {
		println(mergedList.Val)
		mergedList = mergedList.Next
	}
}
