/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode-cn.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (29.61%)
 * Likes:    485
 * Dislikes: 0
 * Total Accepted:    86.9K
 * Total Submissions: 293.2K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
 * 
 * 假设一个二叉搜索树具有如下特征：
 * 
 * 
 * 节点的左子树只包含小于当前节点的数。
 * 节点的右子树只包含大于当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 * 
 * 
 * 示例 1:
 * 
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 * 输出: false
 * 解释: 输入为: [5,1,4,null,null,3,6]。
 * 根节点的值为 5 ，但是其右子节点值为 4 。
 * 
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

 const INT_MAX = int(^uint(0) >> 1)
 const INT_MIN = ^INT_MAX
 //递归传递子树的最大值和最小值,并实时更新
 func validBst(root *TreeNode, upper int, lower int) bool {
	 if root == nil {
		 return true
	 }
 
	 num := root.Val
	 if num <= lower || num >= upper {
		 return false
	 }
 
	 return validBst(root.Left, num, lower) && validBst(root.Right, upper, num)
 }
 
 func isValidBST(root *TreeNode) bool {
	 return validBst(root, INT_MAX, INT_MIN)
 }
 
 
 /*中序遍历是有序的，遍历时直接判断是否有序*/
 func inorder(root *TreeNode, tmp *int) bool {
	 if root != nil {
		 if false == inorder(root.Left, tmp) {
			 return false
		 } 
 
		 if *tmp >= root.Val {
			 return false
		 }
 
		 *tmp = root.Val
		 return inorder(root.Right, tmp)
	 }
 
	 return true
 }
 
 func isValidBST(root *TreeNode) bool {
	 tmp := INT_MIN
	 return inorder(root, &tmp)
 }
 
 // @lc code=end
 
 