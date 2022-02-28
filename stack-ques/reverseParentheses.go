package main

import (
	"strings"
)

var stack []string

func push(v string) {
	stack = append(stack, v)
}

func pop() string {
	e := stack[len(stack)-1]

	stack = stack[:len(stack)-1]
	return e
}
func empty() bool {
	if len(stack) == 0 {
		return true
	}
	return false
}

func reverseParentheses(ss string) string {

	s := strings.Split(ss, "")
	for i := 0; i < len(s); i++ {

		if s[i] == ")" {
			tmp := make([]string, 0)

			for c := pop(); c != "("; {
				tmp = append(tmp, c)
				c = pop()
			}

			for _, v := range tmp {
				push(v)
			}
		} else {
			push(s[i])
		}
	}

	return strings.Join(stack, "")
}
