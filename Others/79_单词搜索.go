package Others

/*给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:

board =
[
['A','B','C','E'],
['S','F','C','S'],
['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false

提示：

board 和 word 中只包含大写和小写英文字母。
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3*/

func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])
	for i:=0;i<n;i++{
		for j:=0;j<m;j++{
			if backtrace79(board,word,i,j,0){
				return true
			}
		}
	}
	return false
}

func backtrace79(board [][]byte,word string,i,j,k int) bool {
	// 找到了
	if k == len(word) {
		return true
	}
	// 越界
	if i < 0 || j <0 || i == len(board) || j == len(board[i]){
		return false
	}
	if board[i][j] != word[k] {
		return false
	}
	tmp := board[i][j]
	// 为了防止重复使用，先置换成空值
	board[i][j] = byte(' ')
	// 开始上下左右探测
	//向左
	if backtrace79(board, word, i-1,j,k+1){
		return true
	}
	// right
	if backtrace79(board, word, i+1,j,k+1){
		return true
	}
	// down
	if backtrace79(board, word, i,j+1,k+1){
		return true
	}
	// up
	if backtrace79(board, word, i,j-1,k+1){
		return true
	}
	board[i][j] = tmp
	return false
}