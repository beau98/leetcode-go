package top150

var keys = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	dfs(digits, 0, "", &res)
	return res
}

func dfs(str string, index int, s string, res *[]string) {
	if index == len(str) {
		*res = append(*res, s)
		return
	}

	for _, c := range keys[int(str[index])-50] {
		dfs(str, index+1, s+string(c), res)
	}
}
