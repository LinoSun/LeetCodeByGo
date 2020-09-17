package Others

/*数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

 

示例：

输入：n = 3
输出：[
"((()))",
"(()())",
"(())()",
"()(())",
"()()()"
]*/

func generateParenthesis(n int) []string {
	res := new([]string)
	backtrack22(n,n,"",res)

	return *res
}

func backtrack22(left,right int,tmp string,res *[]string){
	// 回溯条件
	//右括号用完了就意味着左括号用完了
	if right == 0{
		*res = append(*res, tmp)
		return
	}
	// 生成左括号
	if left > 0{
		backtrack22(left-1,right,tmp+"(",res)
	}
	// 括号成对出现
	if right > left{
		backtrack22(left,right-1,tmp+")",res)
	}
}