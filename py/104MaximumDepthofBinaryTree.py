from typing import Optional, List

import pytest


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        # 如果節點為空，返回深度為0
        if not root:
            return 0

        # 使用遞迴的方式計算左子樹的深度
        left_depth = self.maxDepth(root.left)
        # 使用遞迴的方式計算右子樹的深度
        right_depth = self.maxDepth(root.right)

        # 返回左、右子樹的最大深度 + 1
        # +1 代表當前層的深度
        return max(left_depth, right_depth) + 1


# Helper function to create a binary tree from a list.
def create_binary_tree(data: List[Optional[int]], index: int = 0) -> Optional[TreeNode]:
    if index < len(data) and data[index] is not None:
        node = TreeNode(data[index])
        node.left = create_binary_tree(data, 2 * index + 1)
        node.right = create_binary_tree(data, 2 * index + 2)
        return node
    return None


@pytest.mark.parametrize(
    "tree_data, expected_depth",
    [
        ([3, 9, 20, None, None, 15, 7], 3),  # Standard test case
        ([1, None, 2], 2),  # Only right child for root
        ([1], 1),  # Only root node
        ([], 0),  # Empty tree
        ([1, 2, 3, 4, 5, 6, 7, 8, 9], 4)  # Balanced tree with depth of 4
    ]
)
def test_max_depth(tree_data, expected_depth):
    tree = create_binary_tree(tree_data)
    solution = Solution()
    assert solution.maxDepth(tree) == expected_depth


# This line is required to run the tests when you run the script.
if __name__ == "__main__":
    pytest.main([__file__])
