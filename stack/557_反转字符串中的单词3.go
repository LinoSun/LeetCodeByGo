package stack

/*给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

 

示例：

输入："Let's take LeetCode contest"
输出："s'teL ekat edoCteeL tsetnoc"
 

提示：

在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。*/

func reverseWords(s string) string {
	sArr := []byte(s)
	stack,res := make([]byte,0),make([]byte,0)
	for _,v := range sArr{
		if v == ' '{
			n := len(stack)
			for i:=n-1;i>=0;i--{
				res = append(res, stack[i])
			}
			res = append(res, ' ')
			stack = make([]byte,0)
			continue
		}
		stack = append(stack, v)
	}
	n := len(stack)
	for i:=n-1;i>=0;i--{
		res = append(res, stack[i])
	}
	return string(res)
}
