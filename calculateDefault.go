package weightedAverage

const (
	calculateDefaultWeightingBase    = 2
	calculateDefaultCanReorderValues = false
)

func CalculateDefault[T Number](
	values []T,
) float64 {
	return Calculate(
		values,
		calculateDefaultWeightingBase,
		calculateDefaultCanReorderValues,
	)
}
