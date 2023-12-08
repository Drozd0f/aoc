package day8

// gcd Function to find the greatest common divisor (GCD) of two numbers
func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm Function to find the least common multiple (LCM) of two numbers
func lcm(a, b int64) int64 {
	return a * b / gcd(a, b)
}

// findLCM Function to find the LCM for an array of numbers
func findLCM(numbers []int64) int64 {
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]

	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}

	return result
}
