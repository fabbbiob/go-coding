package ds

import (
	"fmt"

	u "github.com/dockerian/go-coding/utils"
)

// Compare returns 0 if a == b; 1 if a > b; or -1 if a < b
func Compare(a, b uint64) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

// EqualSign checks if a and b are both same signed (positive or negative)
func EqualSign(a, b int64) bool {
	return a >= 0 && b >= 0 || a < 0 && b < 0
}

// MultiplyUint64 gets muliplication of two integers
func MultiplyUint64(a, b uint64) (uint64, error) {
	m := a * b
	if m/a != b {
		return 0, fmt.Errorf("overflow: muliplication of %v, %v", a, b)
	}
	return m, nil
}

// SumUint64 gets sum of two integers
func SumUint64(a, b uint64) (uint64, error) {
	s := a + b
	if s-a != b {
		return 0, fmt.Errorf("overflow: sum of %v, %v", a, b)
	}
	return s, nil
}

// ReverseInt64 reverses a decimal integer
func ReverseInt64(number int64) int64 {
	var result int64
	for number != 0 {
		var unit = number % 10
		var test = result*10 + unit
		// u.Debug("result= %v, test= %v, unit= %v\n", result, test, unit)
		// Note: maximium int64 + 1 => minimium int64 (negative)
		//       minimium int64 - 1 => maximium int64 (positive)
		if (test-unit)/10 != result || result != 0 && !EqualSign(test, result) {
			return 0
		}
		// u.Debug("result= %v, test= %v, number= %v\n", result, test, number)
		result = test
		number /= 10
	}

	u.Debug("number= %v, result= %v\n\n", number, result)
	return result
}