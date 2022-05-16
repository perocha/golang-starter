package internalapp1

// Version of the internal package
var Version = "1.0"

func Test() string {
	return "Test::internalapp1"
}

// Sum two integer numbers
func Sum(number1, number2 int) int {
	return number1 + number2
}
