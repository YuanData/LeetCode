package lc

import (
	"reflect"
	"testing"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head

	// 遍历链表
	for current != nil {
		// 先保存当前节点的下一个节点
		next := current.Next
		// 将当前节点的 Next 指向 prev
		current.Next = prev
		// 更新 prev 为当前节点
		prev = current
		// 移动到下一个节点
		current = next
	}

	return prev
}

func reverseListRecursively(head *ListNode) *ListNode {
	// 如果链表为空，或者链表只有一个节点，那么直接返回该节点。
	if head == nil || head.Next == nil {
		return head
	}

	// 递归调用，处理链表中的下一个节点。
	// newHead 将会是反转后链表的头节点。
	newHead := reverseListRecursively(head.Next)

	// 当前节点的下一个节点（即head.Next）已经被递归地反转了。
	// 我们将下一个节点的 Next 指向当前节点，实现反转。
	head.Next.Next = head

	// 由于当前节点已经被反转，它不再指向原来的下一个节点，所以将 Next 设置为 nil。
	head.Next = nil

	// 返回新的头节点。
	return newHead
}

// 辅助函数：将切片转化为链表
func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, v := range nums[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

// 辅助函数：将链表转化为切片
func listToSlice(node *ListNode) []int {
	var res []int
	current := node
	for current != nil {
		res = append(res, current.Val)
		current = current.Next
	}
	return res
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5},
			output: []int{5, 4, 3, 2, 1},
		},
		{
			input:  []int{1, 2},
			output: []int{2, 1},
		},
		{
			input:  []int{},
			output: []int{},
		},
	}

	for _, test := range tests {
		head := sliceToList(test.input)
		got := listToSlice(reverseList(head))
		if len(test.input) == 0 && len(got) == 0 {
			continue
		}
		if !reflect.DeepEqual(got, test.output) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.output, got)
		}
	}
}

func TestReverseListRecursively(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5},
			output: []int{5, 4, 3, 2, 1},
		},
		{
			input:  []int{1, 2},
			output: []int{2, 1},
		},
		{
			input:  []int{},
			output: []int{},
		},
	}

	for _, test := range tests {
		head := sliceToList(test.input)
		got := listToSlice(reverseListRecursively(head))
		if len(test.input) == 0 && len(got) == 0 {
			continue
		}
		if !reflect.DeepEqual(got, test.output) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.output, got)
		}
	}
}
