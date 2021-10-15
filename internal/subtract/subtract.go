package subtract

import "fmt"

func Run(x, y float64) {
	fmt.Printf("subtracting %f from %f\n", x, y)

	res := subtract(x, y)
	fmt.Printf("result: %f\n", res)
}

func subtract(x, y float64) float64 {
	return x - y
}
