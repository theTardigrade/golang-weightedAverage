package weightedAverage

import (
	"math"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func Calculate[T constraints.Integer | constraints.Float](
	values []T,
	weightingBase float64,
) float64 {
	valuesLen := len(values)

	switch valuesLen {
	case 0:
		return 0
	case 1:
		return float64(values[0])
	}

	var averageValue float64
	var averageValueWeightedCount float64

	if weightingBase == 1 {
		for _, value := range values {
			averageValue += float64(value)
			averageValueWeightedCount++
		}

		return averageValue / averageValueWeightedCount
	}

	valuesClone := make([]T, valuesLen)

	copy(valuesClone, values)

	slices.Sort(valuesClone)

	valuesMidIndex := float64(valuesLen-1) / 2
	weightingBase = math.Abs(weightingBase)

	for i, value := range valuesClone {
		midDistance := math.Abs(valuesMidIndex - float64(i))
		midProximity := valuesMidIndex - midDistance
		weighting := math.Pow(weightingBase, midProximity/valuesMidIndex)

		averageValue += float64(value) * weighting
		averageValueWeightedCount += weighting
	}

	return averageValue / averageValueWeightedCount
}
