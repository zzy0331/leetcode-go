/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (70.18%)
 * Likes:    430
 * Dislikes: 0
 * Total Accepted:    121.4K
 * Total Submissions: 172K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的中序 遍历。
 * 
 * 示例:
 * 
 * 输入: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 * 
 * 输出: [1,3,2]
 * 
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*递归算法*/

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}

func inorderTraversal(root *TreeNode) []int {
	var res []int

	inorder(root, &res)
	return res
}


/*用栈模拟递归*/
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	curr := root

	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, curr.Val)
		curr = curr.Right
	}

	return res
}
// @lc code=end

