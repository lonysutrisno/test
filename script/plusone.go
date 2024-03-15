package script

import (
	"fmt"
	"math"
	"strconv"
)

func PlusOne() []int {
	digits := []int{5, 2, 2, 6, 5, 7, 1, 9, 0, 3, 8, 6, 8, 6, 5, 2, 1, 8, 7, 9, 8, 3, 8, 4, 7, 2, 5, 8, 9}
	var res []int
	if digits[len(digits)-1] < 9 {
		digits[len(digits)-1] = digits[len(digits)-1] + 1

		return digits
	}
	lendigit := len(digits) - 1
	var tempdigit int64

	for _, v := range digits {

		tempdigit = tempdigit + (int64(v) * int64(math.Pow(10.0, float64(lendigit))))
		lendigit = lendigit - 1
	}
	tempdigit = tempdigit + 1
	fmt.Println(tempdigit)

	numberString := strconv.FormatInt(tempdigit, 10)
	for _, char := range numberString {
		digit, _ := strconv.Atoi(string(char))
		res = append(res, digit)
	}

	return res
}
