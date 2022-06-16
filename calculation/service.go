package calculation

import (
	"github.com/tzuhsitseng/demo-checkout/configs"
	"github.com/tzuhsitseng/demo-checkout/domain"
)

type Service struct {
	strategy domain.CalcStrategy
}

func NewService() *Service {
	var strategy domain.CalcStrategy

	if configs.NewStrategy {
		strategy = NewStrategyB()
	} else {
		strategy = NewStrategyA()
	}

	return &Service{strategy: strategy}
}

func (s *Service) Calculate(userLevel, userPoints, price int) (int, int, error) {
	coin, points := s.strategy.Calculate(userLevel, userPoints, price)
	return coin, points, nil
}
