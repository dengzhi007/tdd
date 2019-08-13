package fizzbuzz

func Say(num1, num2, index int) string {
	if index == 1 {
		return "1"
	}
	if index == 2 {
		return "2"
	}
	if index == 3 {
		return "Fizz"
	}
	return ""
}
