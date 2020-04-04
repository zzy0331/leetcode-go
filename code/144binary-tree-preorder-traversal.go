/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Medium (64.09%)
 * Likes:    232
 * Dislikes: 0
 * Total Accepted:    86.7K
 * Total Submissions: 133.5K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 前序 遍历。
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
 * 输出: [1,2,3]
 * 
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

/*递归*/
func preorder(root *TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		preorder(root.Left, res)
		preorder(root.Right, res)
	}
}

func preorderTraversal(root *TreeNode) []int {
	var res []int

	preorder(root, &res)

	return res
}

/*用栈模拟递归*/
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	curr := root
	
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			res = append(res, curr.Val)
			stack = append(stack, curr)
			curr = curr.Left
		}
	
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		curr = curr.Right
	}
	
	return res
}

// @lc code=end

