package fizzbuzz

import (
	"fmt"
	"strings"
)

func FizzBuzz(num1, num2, index int) string {
	if isFizzBuzz(num1, num2, index) {
		return "FizzBuzz"
	}
	if isFizz(num1, num2, index) {
		return "Fizz"
	}
	if isBuzz(num1, num2, index) {
		return "Buzz"
	}
	return fmt.Sprint(index)
}

func isFizz(num1, num2, index int) bool {
	if index%num1 == 0 {
		return true
	}
	if strings.Contains(fmt.Sprint(index), fmt.Sprint(num1)) {
		return true
	}
	return false
}
func isBuzz(num1, num2, index int) bool {
	if index%num2 == 0 {
		return true
	}
	if strings.Contains(fmt.Sprint(index), fmt.Sprint(num2)) {
		return true
	}
	return false
}
func isFizzBuzz(num1, num2, index int) bool {
	if index%num1 == 0 && index%num2 == 0 {
		return true
	}
	if strings.Contains(fmt.Sprint(index), fmt.Sprint(num2)) {
		return true
	}
	return false
}
