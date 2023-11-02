package strings

import "testing"

func TestInInt64(t *testing.T) {
	t.Log(WithinParenthesesInt64([]int64{1, 2, 3}))
}
