package top150

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	longestCommonPrefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(strs[i]) && j < len(longestCommonPrefix); j++ {
			if longestCommonPrefix[j] != strs[i][j] {
				longestCommonPrefix = longestCommonPrefix[0:j]
				break
			}
		}
		if len(longestCommonPrefix) > len(strs[i]) {
			longestCommonPrefix = longestCommonPrefix[0:len(strs[i])]
		}
	}
	return longestCommonPrefix
}
