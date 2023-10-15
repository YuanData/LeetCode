from typing import Optional

import pytest


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # 初始化前一个节点为None，因为反转后的最后一个节点将指向None
        prev = None
        # 初始化当前节点为链表的头节点
        current = head

        # 迭代链表，直到当前节点为None，即遍历完整个链表
        while current:
            # 保存下一个节点的引用，以便在反转后重新连接
            next_node = current.next
            # 将当前节点的指针反向指向前一个节点
            current.next = prev
            # 更新prev为当前节点，以便下一轮迭代使用
            prev = current
            # 更新current为下一个节点，继续迭代
            current = next_node

        # 当循环结束时，prev将成为新的头节点，即反转后的链表的头部
        return prev

    def reverseListRecursively(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # 递归的终止条件：当链表为空或只有一个节点时，无需反转，直接返回
        if not head or not head.next:
            return head

        # 递归地反转子链表
        new_head = self.reverseListRecursively(head.next)

        # 将当前节点的下一个节点的指针指向当前节点，实现反转
        head.next.next = head
        head.next = None

        return new_head


@pytest.fixture
def solution():
    return Solution()


def test_reverse_list_case(solution):
    head2 = ListNode(1, ListNode(2))
    reversed_head2 = solution.reverseListRecursively(head2)
    assert reversed_head2.val == 2
    assert reversed_head2.next.val == 1
    assert reversed_head2.next.next is None


test_data = [
    (ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(5))))),
     [5, 4, 3, 2, 1]),
    (ListNode(1, ListNode(2)), [2, 1]),
    (None, [])
]


def clone_linked_list(head):
    """实用功能，用于克隆链表。"""
    if not head:
        return None
    return ListNode(head.val, clone_linked_list(head.next))


@pytest.mark.parametrize("method_name", ["reverseList", "reverseListRecursively"])
@pytest.mark.parametrize("input_head, expected_values", test_data)
def test_reverse_list(solution, method_name, input_head, expected_values):
    cloned_head = clone_linked_list(input_head)
    method = getattr(solution, method_name)
    reversed_head = method(cloned_head)

    current = reversed_head
    for expected_val in expected_values:
        assert current.val == expected_val
        current = current.next
    assert current is None


if __name__ == '__main__':
    pytest.main()
