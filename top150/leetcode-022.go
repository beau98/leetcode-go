package top150

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	generateDFS("", n, 2*n, &res)
	return res
}

/**
max: 传入的max,表示最终括号串的长度
strs: 答案数组
index: 当前剩余左括号的数量
*/
func generateDFS(s string, index int, max int, strs *[]string) {
	// base case
	if max == len(s) {
		*strs = append(*strs, s)
		return
	}
	//can filled with left parenthesis
	if index > 0 {
		generateDFS(s+"(", index-1, max, strs)
	}
	//can filled with right parenthesis
	if max/2-index > len(s)/2 {
		generateDFS(s+")", index, max, strs)
	}
}
