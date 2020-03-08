// 暴力；基本情况怎么样
// 找最近重复子问题
// 1: 1
// 2: 2
// 3: f(1) + f(2)
// 4: f(2) + f(3)
// n: f(n-2) + f(n-1) ->斐波拉契数列

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	f1 := 1
	f2 := 2
	var f3 int
	for i := 3; i <= n; i++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
	return f3
}
