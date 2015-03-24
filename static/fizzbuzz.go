package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	for i := 0; i < 50; i++ {
		var output string
		if m := math.Mod(float64(i), 3.0); m == 0 {
			output += "Fizz"
		}

		if math.Mod(float64(i), 5.0) == 0 {
			output += "Buzz"
		}

		if len(output) == 0 {
			fmt.Print(strconv.Itoa(i))
			// fmt.Print(i)
		}
		fmt.Println(output)
	}
}
