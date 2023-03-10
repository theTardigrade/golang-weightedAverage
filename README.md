# golang-weightedAverage

## Introduction

This package exposes a function to calculate a weighted average (i.e. a mean that is either skewed towards or away from extreme values).

The ```Calculate``` function takes two arguments: a generic slice of numeric values (either integral or floating-point) and a pointer to an options struct (of the type ```CalculateOptions```) that determines how the average should be calculated.

The main work of function involves iterating through a version of the slice that has been sorted in ascending order (either the original slice, if allowed, or a cloned version, if not), calculating a weighted average of the numbers that it contains.

The result is always given as a 64-bit floating-point number, even if the original slice contained integers.

## Options

The first field in the options struct is ```WeightingBase```, which is a 64-bit floating-point number that defines how weight is distributed.

A weighting base greater than ```1``` gives more weight to the numbers that are towards the middle of the slice (i.e. those numbers closest to the median), while a weighting base less than ```1``` gives more weight to those that are furthest away from the middle, and a weighting base that is exactly ```1``` treats all the numbers equally (producing an ordinary mean value).

If the weighting base is negative, it will be treated the same as if it were positive (i.e. its absolute value will be used).

The second field in the options struct is ```CanReorderValues``` a boolean value that determines whether the algorithm can rearrange the ordering of the original slice or whether a cloned version of the slice should be used internally.

If you call the ```Calculate``` function with the options argument as ```nil```, then a default set of options will be used.

It is also possible to access the default options directly by calling the ```CalculateOptionsWithDefaults``` function, which will return a pointer to a freshly defined struct, where the fields have been set to reasonable defaults: ```WeightingBase``` will be set to ```2``` and ```CanReorderValues``` will be set to ```false```.

## Example

```golang
package main

import (
	"fmt"

	weightedAverage "github.com/theTardigrade/golang-weightedAverage"
)

func main() {
	numbers := []int{9, 2, 8, 6, 7, 100, 4, 3, 5}

	options := weightedAverage.CalculateOptions{
		WeightingBase: 1,
	}
	averageNumber := weightedAverage.Calculate(numbers, &options)

	fmt.Printf(
		"The average number with a weighting base of %.0f is %.2f.\n",
		options.WeightingBase,
		averageNumber,
	) // 16.00

	averageNumber = weightedAverage.Calculate(numbers, nil) // 2 is the default WeightingBase

	fmt.Printf(
		"The average number with a weighting base of %.0f is %.2f.\n",
		weightedAverage.CalculateOptionsWithDefaults().WeightingBase,
		averageNumber,
	) // 13.16

	options.WeightingBase = 32
	averageNumber = weightedAverage.Calculate(numbers, &options)

	fmt.Printf(
		"The average number with a weighting base of %.0f is %.2f.\n",
		options.WeightingBase,
		averageNumber,
	) // 7.17

	options.WeightingBase = 100
	averageNumber = weightedAverage.Calculate(numbers, &options)

	fmt.Printf(
		"The average number with a weighting base of %.0f is %.2f.\n",
		options.WeightingBase,
		averageNumber,
	) // 6.47

	options.WeightingBase = 0.1
	averageNumber = weightedAverage.Calculate(numbers, &options)

	fmt.Printf(
		"The average number with a weighting base of %.1f is %.2f.\n",
		options.WeightingBase,
		averageNumber,
	) // 27.36
}
```