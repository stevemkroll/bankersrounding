package round

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var values = []struct {
	input   float64
	decimal int
	output  float64
}{
	{1.5, 0, 2},
	{2.5, 0, 2},
	{1.535, 2, 1.54},
	{1.525, 2, 1.52},
	{0.5, 0, 0},
	{1.5, 0, 2},
	{0.4, 0, 0},
	{0.6, 0, 1},
	{1.4, 0, 1},
	{1.6, 0, 2},
	{23.5, 0, 24},
	{24.5, 0, 24},
	{-23.5, 0, -24},
	{-24.5, 0, -24},
	{1.534953, 2, 1.53},
	{1.53499999, 2, 1.53},
	{1.5299999, 2, 1.53},
	{1.5350, 2, 1.54},
	{1.53599, 2, 1.54},
	{3.2, 0, 3},
	{3.4, 0, 3},
	{3.5, 0, 4},
	{4.5, 0, 4},
	{5.5, 0, 6},
	{-7.5, 0, -8},
}

// need to check for non decimal numbers

func extractDetails(flt float64, dec int) (string, string, string) {
	str := fmt.Sprintf("%+v", flt)
	pre := ""
	if strings.Contains(str, "-") {
		pre = "-"
		str = strings.Replace(str, "-", "", -1)
	}
	num := strings.SplitN(str, ".", -1)
	return pre, num[0], num[1][:dec+1]
}

func bankersRound(negative string, full string, fractional string) string {
	if len(fractional) > 1 {
		f1 := fractional[:len(fractional)-1]
		f2 := fractional[len(fractional)-1:]
		n1, _ := strconv.Atoi(f1)
		n2, _ := strconv.Atoi(f2)
		if n2 > 4 {
			n1++
			if n2 == 5 && n1%2 != 0 {
				n1--
			}
		}
		fractional = strconv.Itoa(n1)
		return fmt.Sprintf("%s%s.%s", negative, full, fractional)
	}
	n1, _ := strconv.Atoi(full)
	n2, _ := strconv.Atoi(fractional)
	if n2 > 4 {
		n1++
		if n2 == 5 && n1%2 != 0 {
			n1--
		}
	}
	return fmt.Sprintf("%s%+v", negative, n1)
}

func TestBankersRounding(t *testing.T) {
	for _, i := range values {
		negative, full, fractional := extractDetails(i.input, i.decimal)
		result := bankersRound(negative, full, fractional)
		if result != fmt.Sprintf("%+v", i.output) {
			t.Logf("RESULT:%+v...EXPECTED:%+v", result, i.output)
		}
	}
}
