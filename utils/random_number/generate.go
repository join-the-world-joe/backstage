package random_number

import (
	"fmt"
	"math/rand"
)

func Generate(min, max, n int) (string, error) {
	var output string
	for i := 1; i <= n; i++ {
		output = output + fmt.Sprintf("%v", rand.Intn(max-min+1)+min)
	}
	return output, nil
}
