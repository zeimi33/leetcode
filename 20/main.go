package main

func main() {

}

func isValid(s string) bool {

	stack := []rune{}
	for _, c := range s {
		switch c {
		case '(', '{', '[':
			stack = append(stack, c)
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
