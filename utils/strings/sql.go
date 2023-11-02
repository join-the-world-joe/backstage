package strings

import "fmt"

func WithinParenthesesInt64(l []int64) string {
	c := "("
	length := len(l)
	for k, v := range l {
		if k < length-1 {
			c += fmt.Sprintf("%v, ", v)
			continue
		}
		c += fmt.Sprintf("%v", v)
	}
	c += ")"
	return c
}
