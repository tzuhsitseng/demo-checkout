package configs

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	VIPDiscounts map[int]float64
	PointsOffset float64
	NewStrategy  bool
)

func Init() error {
	_ = godotenv.Load()

	if vipDiscounts := os.Getenv("vip_discounts"); vipDiscounts == "" {
		return errors.New("cannot get env of vip discounts")
	} else {
		discounts := strings.Split(vipDiscounts, ",")
		VIPDiscounts = make(map[int]float64, len(discounts))

		for _, v := range discounts {
			levelWithRatio := strings.Split(v, ":")
			if len(levelWithRatio) == 2 {
				level, err := strconv.ParseInt(levelWithRatio[0], 10, 32)
				if err != nil {
					return errors.New("something wrong on level")
				}

				ratio, err := strconv.ParseFloat(levelWithRatio[1], 64)
				if err != nil {
					return errors.New("something wrong on ratio")
				}

				VIPDiscounts[int(level)] = ratio
			}
		}
		if len(VIPDiscounts) == 0 {
			return errors.New("cannot get env of vip discounts")
		}
	}

	if pointsOffset := os.Getenv("points_offset"); pointsOffset == "" {
		return errors.New("cannot get env of points offset")
	} else {
		var err error
		PointsOffset, err = strconv.ParseFloat(pointsOffset, 64)
		if err != nil {
			return errors.New("something wrong on offset")
		}
	}

	if newStrategy := os.Getenv("new_strategy"); newStrategy == "" {
		return errors.New("cannot get env of new strategy")
	} else {
		var err error
		NewStrategy, err = strconv.ParseBool(newStrategy)
		if err != nil {
			return errors.New("something wrong on new strategy")
		}
	}

	return nil
}
