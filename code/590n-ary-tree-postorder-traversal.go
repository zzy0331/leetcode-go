/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/description/
 *
 * algorithms
 * Easy (72.78%)
 * Likes:    57
 * Dislikes: 0
 * Total Accepted:    19K
 * Total Submissions: 26.1K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的后序遍历。
 * 
 * 例如，给定一个 3叉树 :
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 返回其后序遍历: [5,6,3,2,4,1].
 * 
 * 
 * 
 * 说明: 递归法很简单，你可以使用迭代法完成此题吗?
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

/*递归*/
func order(root *Node, res *[]int) {
	if root != nil {
		for _, v := range root.Children {
			order(v, res)
		}
		*res = append(*res, root.Val)
	}
}

func postorder(root *Node) []int {
	var res []int

	order(root, &res)

	return res
}
// @lc code=end

