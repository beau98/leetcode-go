package top150

import "math"

func myAtoi(s string) int {
	arr := []byte(s)
	result := 0
	flag := 1
	i := 0
	for ; i < len(arr); i++ {
		if arr[i] == 32 {
			continue
		}
		break
	}

	if arr[i] == 43 {
		i++
	}

	if arr[i] == 45 {
		i++
		flag = -1
	}

	for ; i < len(arr) && isDigit(arr[i]); i++ {
		result = result*10 + int(arr[i]) - 48
		if result > math.MaxInt32 {
			return math.MaxInt32
		}
		if result < -math.MaxInt32-1 {
			return math.MinInt32
		}
	}
	return result * flag
}

func isDigit(b byte) bool {
	return (b >= 48) && (b <= 57)
}
