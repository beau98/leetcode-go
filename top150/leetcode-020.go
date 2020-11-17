package top150

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, r := range s {
		if isPush(r) {
			stack = append(stack, r)
		} else if !pop(r, &stack) {
			return false
		}
	}
	return len(stack) == 0
}

func pop(r rune, stack *[]rune) bool {
	if len(*stack) == 0 {
		return false
	}
	if r == ']' && (*stack)[len(*stack)-1] != '[' {
		return false
	}
	if r == '}' && (*stack)[len(*stack)-1] != '{' {
		return false
	}
	if r == ')' && (*stack)[len(*stack)-1] != '(' {
		return false
	}
	*stack = (*stack)[:len(*stack)-1]
	return true
}

func isPush(r rune) bool {
	if r == '(' || r == '[' || r == '{' {
		return true
	}
	return false
}
