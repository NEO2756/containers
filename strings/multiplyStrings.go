package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	multiply("123", "456")
}
func multiply(num1 string, num2 string) string {

	s, l := num1, num2
	if len(s) > len(l) {
		s, l = l, s
	}

	s = reverse(strings.Split(s, ""))
	l = reverse(strings.Split(l, ""))

	final := ""
	for i := 0; i < len(s); i++ {
		res := mul((s[i]), l, i)
		final = add(final, res)
		fmt.Println(final)
	}

	return final
}

func add(s, l string) string {

	if len(s) == 0 {
		return l
	}

	if len(l) == 0 {
		return s
	}

	result := make([]string, 0)
	i := 0
	c := 0
	s = reverse(strings.Split(s, ""))
	l = reverse(strings.Split(l, ""))

	for i = 0; i < len(l); i++ {
		x, y := 0, 0
		if i < len(s) {
			x = int(s[i] - '0')
		}
		if i < len(l) {
			y = int(l[i] - '0')
		}
		res := c + x + y

		if i != len(l)-1 {
			c = res / 10
			res = res % 10
		}
		result = append(result, strconv.Itoa(res))
	}

	return reverse(result)
}
func mul(s byte, num string, decimal int) string {

	c := 0
	res := 0
	result := make([]string, 0)

	for i := 0; i < len(num); i++ {
		res = c + int((s-'0')*(num[i]-'0'))

		if i != len(num)-1 {
			c = res / 10
			res = res % 10
		}

		result = append(result, strconv.Itoa(res))
	}

	str := reverse(result)
	str = str + zeros(decimal)
	return str
}

func zeros(i int) string {
	s := make([]string, 0)
	for ; i > 0; i-- {
		s = append(s, "0")
	}
	return strings.Join(s, "")
}
func reverse(s []string) string {
	for start, end := 0, len(s)-1; start < end; {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}

	return strings.Join(s, "")
}
