package random_number

import "testing"

func TestGenerate(t *testing.T) {
	n := 10
	digit := 4
	for i := 1; i <= n; i++ {
		code, err := Generate(0, 9, digit)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("code: ", code)
	}
}
