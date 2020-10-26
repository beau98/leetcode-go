package top150

/**
Given a string s, find the length of the longest substring without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Example 4:
Input: s = ""
Output: 0

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

func LengthOfLongestSubstring(s string) int {
	smap := make(map[byte]int)
	start, res := 0, 0
	for i := 0; i < len(s); i++ {
		index, ok := smap[s[i]]
		if ok && index >= start {
			start = index + 1
		}
		smap[s[i]] = i
		res = max(res, i-start+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
