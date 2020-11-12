package interview

import (
	"fmt"
	"sort"
	"testing"
)
 /*

 Codewriting

 You have N types of securities available to buy and for each
 security i, each has a starting price Pi.
 Your friend predicts that the stock price will be Si for each
 associated security.
 But based on volatility of the market, you only want to buy
 up to Ai shares of security i.

 Given M dollars to spend, calculate the maximum return you
 could potentially gain based on the predicted prices Si (and
 including any cash you have remaining). You can think about
 this as the final portfolio value according to the predicted
 prices.

 Assume you are able to buy fractional amounts of shares.

 N = Number of securities
 Pi = Current Price
 Si = Expected Future Price
 Ai = Maximum units you are willing to purchase
 M = Dollars available to invest
 --- Sample Input ---
 current_prices: [15, 20]
 predicted_prices: [30, 45]
 max_number_of_units: [3, 3]
 dollars_to_spend: 30

 AAPL's current price is $15, predicted price is $30, and max units to buy is 3.
 TSLA current price is $20, predicted price is $45, and max units to buy is 3.
 We have $30 to spend.

 ---Sample Output---
 Max portfolio value is $67.5.

 [execution time limit] 4 seconds (go)

 [input] array.float current_prices

 [input] array.float predicted_prices

 [input] array.integer max_number_of_units

 [input] integer dollars_to_spend

 [output] float
  */

func portfolio_value_optimization(current_prices []float64, predicted_prices []float64, max_number_of_units []int, dollars_to_spend int) float64 {
	n := len(current_prices)
	gain := make([]float64, n)
	best := make([]int, n)
	for i, x := range current_prices {
		gain[i] = (predicted_prices[i] - x) / x
		best[i] = i
	}
	sort.Slice(best, func(i, j int) bool {
		return gain[best[i]] > gain[best[j]]
	})
	r := 0.0
	d := float64(dollars_to_spend)
	i := 0
	fmt.Println(best)
	fmt.Println(gain)
	h := map[int]float64{}
	for d > 0 && i < n {
		fmt.Println("exit if <0", gain[best[i]])
		if gain[best[i]] <= 0.0 {
			fmt.Println("exit", gain[best[i]])
			break
		}
		x := d / current_prices[best[i]]
		fmt.Println("buy", x, "max", max_number_of_units[best[i]])
		if x > float64(max_number_of_units[best[i]]) {
			x = float64(max_number_of_units[best[i]])
		}
		r += predicted_prices[best[i]] * x
		d -= current_prices[best[i]] *x
		h[best[i]] = r
		i++
	}
	return r + d
}

func TestPortfolio1(t *testing.T) {
	AssertPortfolio(t,
		67.5,
		[]float64{15, 20},
		[]float64{30, 45},
		[]int{3, 3},
		30)
}

func TestPortfolio2(t *testing.T) {
	AssertPortfolio(t,
		60,
		[]float64{15, 30},
		[]float64{30, 50},
		[]int{3, 3},
		30)
}

func TestPortfolio3(t *testing.T) {
	AssertPortfolio(t,
		60,
		[]float64{15, 15},
		[]float64{45, 10},
		[]int{1, 1},
		30)
}

func TestPortfolio4(t *testing.T) {
	AssertPortfolio(t,
		265,
		[]float64{15, 25, 40, 30},
		[]float64{45, 35, 50, 25},
		[]int{3, 3, 3, 4},
		140)
}

func TestPortfolio45(t *testing.T) {
	AssertPortfolio(t,
		221.83177570093457,
		[]float64{31.3, 32.1, 25.4, 31.7, 45.1},
		[]float64{56.8, 82.8, 7.4, 12.5, 59.7},
		[]int{60, 20, 29, 75, 5},
		86)
}

func TestPortfolio6(t *testing.T) {
	AssertPortfolio(t,
		167,
		[]float64{5, 4, 30, 6, 7},
		[]float64{8, 7.5, 27, 9, 7},
		[]int{10, 8, 40, 9, 17	},
		100)
}

func AssertPortfolio(
	t *testing.T,
	expected float64,
	current_prices []float64,
	predicted_prices []float64,
	max_number_of_units []int,
	dollars_to_spend int) {
	actual := portfolio_value_optimization(current_prices, predicted_prices, max_number_of_units, dollars_to_spend)
	if actual != expected {
		t.Fatal("Expected", expected, "found", actual)
	}
}

