package daily680

import (
	"fmt"
	"testing"
)

type Unit struct {
	Name string
	Factor float64
}

var units map[string]Unit

func AddBaseUnit(name string) {
	units = map[string]Unit{}
	units[name] = Unit{name, 1}
}

func AddDerivedUnit(name string, factor float64, basedon string) {
	base, ok := units[basedon]
	if !ok {
		panic(fmt.Sprint("Unable to find base unit", basedon, "for new unit", name))
	}
	units[name] = Unit{name, factor * base.Factor}
}

func ConvertUnit(amount float64, fromName, toName string) float64 {
	from, ok := units[fromName]
	if !ok {
		panic(fmt.Sprint("Unable to find source unit", fromName))
	}
	to, ok := units[toName]
	if !ok {
		panic(fmt.Sprint("Unable to find source unit", toName))
	}
	return amount * from.Factor / to.Factor
}

func TestDaily689(t *testing.T) {
	AddBaseUnit("inch")
	AddDerivedUnit("thou", 1/12_000, "inch")
	AddDerivedUnit("foot", 12, "inch")
	AssertEqualFloat64(t, 12, ConvertUnit(1, "foot", "inch"), "1 foot to inch")
}

func AssertEqualFloat64(t * testing.T, expected, actual float64, message string) {
	if expected - actual > expected / 1000 {
		t.Fatal("expected", expected, "actual", actual, message)
	}
}
