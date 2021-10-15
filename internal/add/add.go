package add

import "fmt"

func Run(x, y float64) {
	fmt.Printf("adding %f to %f\n", x, y)

	res := add(x, y)
	fmt.Printf("result: %f\n", res)
}

func add(x, y float64) float64 {
	return x + y
}
