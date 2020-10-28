package top150

import "math"

/**
Given a 32-bit signed integer, reverse digits of an integer.

Note:
Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

Example 1:
Input: x = 123
Output: 321

Example 2:
Input: x = -123
Output: -321

Example 3:
Input: x = 120
Output: 21

Example 4:
Input: x = 0
Output: 0

Constraints:

-2^31 <= x <= 2^31 - 1
*/

func reverse(x int) int {
	result := 0
	for x != 0 {
		result = result*10 + x%10
		x = x / 10

		if result > math.MaxInt32 || result < -math.MaxInt32-1 {
			return 0
		}
	}
	return result
}
