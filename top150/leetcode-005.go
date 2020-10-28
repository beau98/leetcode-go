package top150

/**
Given a string s, return the longest palindromic substring in s.

Example 1:
Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example 2:
Input: s = "cbbd"
Output: "bb"

Example 3:
Input: s = "a"
Output: "a"

Example 4:
Input: s = "ac"
Output: "a"

Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters (lower-case and/or upper-case),
*/

func longestPalindrome(s string) string {
	handleString := preHandleString(s)
	maxRight := 0
	center := 0
	arr := []rune(handleString)
	help := make([]int, len(arr))
	max := 0
	index := 0
	for i := 0; i < len(arr); i++ {
		if i >= maxRight {
			radius := maxRadius(i, 0, arr)
			help[i] = radius

			if radius+i > maxRight {
				center = i
				maxRight = radius + i
			}
		} else {
			//mirror left
			left := 2*center - i
			if help[left]+i < maxRight {
				help[i] = help[left]

			} else {
				radius := maxRadius(i, maxRight-i, arr)
				if radius+i > maxRight {
					center = i
					maxRight = radius + i
				}
				help[i] = radius
			}
		}
		if help[i] > max {
			index = i
		}
	}
	runes := arr[index-max : index+max]
	var result []rune
	for _, r := range runes {
		if r != '#' {
			result = append(result, r)
		}
	}
	return string(result)
}

// pre handle with '#'
func preHandleString(s string) string {
	res := []rune{'#'}
	for i := 0; i < len(s); i++ {
		res = append(res, rune(s[i]), '#')
	}
	return string(res)
}

// get max radius
func maxRadius(center int, offset int, arr []rune) int {
	res := offset
	l, r := center-offset-1, center+offset+1
	for l >= 0 && r < len(arr) {
		if arr[l] == arr[r] {
			res++
			l--
			r++
		} else {
			break
		}
	}
	return res
}
