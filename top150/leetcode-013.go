package top150

func romanToInt(s string) int {

	ints := make([]int, len(s))
	//to int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			ints[i] = 1
		case 'V':
			ints[i] = 5
		case 'X':
			ints[i] = 10
		case 'L':
			ints[i] = 50
		case 'C':
			ints[i] = 100
		case 'D':
			ints[i] = 500
		case 'M':
			ints[i] = 1000
		}
	}

	res := 0
	for i := 0; i < len(ints)-1; i++ {
		if ints[i] < ints[i+1] {
			res -= ints[i]
		} else {
			res += ints[i]
		}
	}
	res += ints[len(ints)-1]
	return res
}
