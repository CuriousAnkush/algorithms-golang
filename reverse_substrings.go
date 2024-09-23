package algorithms

func ReverseParentheses(s string) string {
	var stack []rune
	var openingBraces []int

	for idx := range s {
		switch rune(s[idx]) {
		case '(':
			pushToStack(s[idx], &stack)
			openingBraces = append(openingBraces, idx)
		case ')':
			popAndReverse(&stack, &openingBraces)
		default:
			pushToStack(s[idx], &stack)
		}
	}

	return string(stack)
}

func pushToStack(ch byte, stack *[]rune) {
	*stack = append(*stack, rune(ch))
}

func popAndReverse(stack *[]rune, braces *[]int) {
	ob := *braces
	st := *stack
	lastIdx := len(st) - 1
	iniIDx := ob[len(ob)-1]

	for iniIDx < lastIdx {
		st[iniIDx], st[lastIdx] = st[lastIdx], st[iniIDx]
		iniIDx++
		lastIdx--
	}

	st = st[:len(st)-1]
	ob = ob[:len(ob)-1]
	*stack = st
	*braces = ob
}
