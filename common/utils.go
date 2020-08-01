package common

import (
	"fmt"
	"strconv"
)

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	fmt.Println("-----", value)
	return value
}
