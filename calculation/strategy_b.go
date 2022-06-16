package calculation

import (
	"github.com/tzuhsitseng/demo-checkout/configs"
	"github.com/tzuhsitseng/demo-checkout/utils"
)

type StrategyB struct {
}

func NewStrategyB() *StrategyB {
	return &StrategyB{}
}

func (s *StrategyB) Calculate(level, points, price int) (int, int) {
	if level > 0 {
		price = int(float64(price) * configs.VIPDiscounts[level])
	}
	if points >= 100 {
		price = price * 9 / 10
	}
	if points > 0 {
		canUsedPoints := utils.Min(int(float64(points)/configs.PointsOffset), points)
		price = utils.Max(price-canUsedPoints, 0)
		points = canUsedPoints
	}
	return price, points
}
