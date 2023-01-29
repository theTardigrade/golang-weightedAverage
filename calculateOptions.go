package weightedAverage

type CalculateOptions struct {
	WeightingBase    float64
	CanReorderValues bool
}

func CalculateOptionsWithDefaults() *CalculateOptions {
	return &CalculateOptions{
		WeightingBase:    2,
		CanReorderValues: false,
	}
}

var (
	calculateOptionsWithDefaults = CalculateOptionsWithDefaults()
)
