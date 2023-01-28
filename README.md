# golang-weightedAverage

## How It Works

This package exposes a function to calculate a weighted average (i.e. a mean that is either skewed towards or away from extreme values).

The ```Calculate``` function takes three arguments: a generic slice of numeric values (either integral or floating-point) and a weighting base (as a 64-bit floating-point number), followed by a boolean value that determines whether the algorithm can reorder the original slice or whether a cloned version of the slice should be used internally.

The main work of function involves iterating through a version of the slice that has been sorted in ascending order (either the original slice, if allowed, or a cloned version, if not), calculating a weighted average of the numbers that it contains.

A weighting base greater than ```1``` gives more weight to the numbers that are towards the middle of the slice (i.e. those numbers closest to the median), while a weighting base less than ```1``` gives more weight to those that are furthest away from the middle, and a weighting base that is exactly ```1``` treats all the numbers equally (producing an ordinary mean value).

If the weighting base is negative, it will be treated the same as if it were positive (i.e. its absolute value will be used).

The result is always given as a 64-bit floating-point number, even if the original slice contained integers.

There is also a ```CalculateDefault``` function, which takes a generic slice of numeric values as its single argument: it uses a default weighting base of ```2``` and does not allow the original slice to be reordered.

## Example

```golang
package main

import (
	"fmt"

	weightedAverage "github.com/theTardigrade/golang-weightedAverage"
)

func main() {
	numbers := []int{9, 2, 8, 6, 7, 100, 4, 3, 5}
	averageNumber := weightedAverage.Calculate(numbers, 1, true)

	fmt.Printf("The average number with a weighting base of 1 is %.2f.\n", averageNumber) // 16.00

	averageNumber = weightedAverage.Calculate(numbers, 2, true)

	fmt.Printf("The average number with a weighting base of 2 is %.2f.\n", averageNumber) // 13.16

	averageNumber = weightedAverage.Calculate(numbers, 32, true)

	fmt.Printf("The average number with a weighting base of 32 is %.2f.\n", averageNumber) // 7.17

	averageNumber = weightedAverage.Calculate(numbers, 100, true)

	fmt.Printf("The average number with a weighting base of 100 is %.2f.\n", averageNumber) // 6.47

	averageNumber = weightedAverage.Calculate(numbers, 0.1, true)

	fmt.Printf("The average number with a weighting base of 0.1 is %.2f.\n", averageNumber) // 27.36

	fmt.Printf(
		"These two function calls are equivalent: %t.\n",
		weightedAverage.Calculate(numbers, 2, false) == weightedAverage.CalculateDefault(numbers),
	) // true
}
```