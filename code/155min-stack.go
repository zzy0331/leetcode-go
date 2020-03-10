/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 *
 * https://leetcode-cn.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (51.50%)
 * Likes:    394
 * Dislikes: 0
 * Total Accepted:    76.7K
 * Total Submissions: 148.1K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
 * 
 * 
 * push(x) -- 将元素 x 推入栈中。
 * pop() -- 删除栈顶的元素。
 * top() -- 获取栈顶元素。
 * getMin() -- 检索栈中的最小元素。
 * 
 * 
 * 示例:
 * 
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 * 
 * 
 */

// @lc code=start
/*可以用两个栈分别表示本身的值和最小值;或者使用一个栈,栈内元素包含本身的值和最小值*/

type MinStack struct {
    Stack []Node
}

type Node struct {
    Val int
    Min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(x int)  {
    newNode := Node{Val:x, Min:x}
    if len(this.Stack) != 0 && this.Stack[len(this.Stack)-1].Min < x {
        newNode.Min = this.Stack[len(this.Stack)-1].Min
    }
    
    this.Stack = append(this.Stack, newNode)
}


func (this *MinStack) Pop()  {
    this.Stack = this.Stack[:len(this.Stack)-1]
}


func (this *MinStack) Top() int {
    return this.Stack[len(this.Stack)-1].Val
}


func (this *MinStack) GetMin() int {
    return this.Stack[len(this.Stack)-1].Min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

