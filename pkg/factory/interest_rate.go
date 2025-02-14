package factory

import "creditas/pkg/services"

type InterestRateFactory struct{}

func (f *InterestRateFactory) CreateStrategy(rateType string) services.InterestRateStrategy {
	switch rateType {
	case "variable":
		return &services.VariableInterestRateStrategy{}
	default:
		return &services.DefaultInterestRateStrategy{}
	}
}
