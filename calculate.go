package weightedAverage

import (
	"math"

	"golang.org/x/exp/slices"
)

func Calculate[T Number](
	values []T,
	weightingBase float64,
	canReorderValues bool,
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
	} else {
		if !canReorderValues {
			valuesClone := make([]T, valuesLen)

			copy(valuesClone, values)

			values = valuesClone
		}

		slices.Sort(values)

		valuesMidIndex := float64(valuesLen-1) / 2
		weightingBase = math.Abs(weightingBase)

		for i, value := range values {
			midDistance := math.Abs(valuesMidIndex - float64(i))
			midProximity := valuesMidIndex - midDistance
			weighting := math.Pow(weightingBase, midProximity/valuesMidIndex)

			averageValue += float64(value) * weighting
			averageValueWeightedCount += weighting
		}
	}

	return averageValue / averageValueWeightedCount
}
