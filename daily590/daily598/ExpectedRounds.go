package daily598

import "math"

func ExpectedRounds(coins int) int {
	return int(math.Ceil(math.Log(float64(coins))/math.Log(2.0)))
}
