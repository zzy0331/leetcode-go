/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/
 *
 * algorithms
 * Easy (72.72%)
 * Likes:    70
 * Dislikes: 0
 * Total Accepted:    21.9K
 * Total Submissions: 30.1K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的前序遍历。
 * 
 * 例如，给定一个 3叉树 :
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 返回其前序遍历: [1,3,5,6,2,4]。
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
		*res = append(*res, root.Val)
		for _, v := range root.Children {
			order(v, res)
		}
	}
}

func preorder(root *Node) []int {
	var res []int

	order(root, &res)

	return res
}
// @lc code=end

