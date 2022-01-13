package sorting

import (
	"fmt"
	"strconv"
)

const (
	Format1 = "This is the number %.1f"
	Format2 = "This is a box containing the number %.1f"
	Format3 = "This is a fancy box containing the number %.1f"
)

func describeNumber(s string, f float64) string {
	return fmt.Sprintf(s, f)
}

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return describeNumber(Format1, f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return describeNumber(Format2, float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if _, ok :=fnb.(FancyNumber); ok {
		i, err := strconv.Atoi(fnb.Value())
		if err == nil {
			return i
		}
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return describeNumber(Format3, float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch i.(type) {
	case int:
		return DescribeNumber(float64(i.(int)))
	case float64:
		return DescribeNumber(i.(float64))
	case NumberBox:
		return DescribeNumberBox(i.(NumberBox))
	case FancyNumberBox:
		return DescribeFancyNumberBox(i.(FancyNumberBox))
	default:
		return "Return to sender"
	}
  
}
