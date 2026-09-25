package challenges

import "fmt"

type PascalTriangle struct {
}

func (challenge *PascalTriangle) DemonstrateChallenge() {
	var input int
	fmt.Print("Enter an integer: ")
	if _, err := fmt.Scanln(&input); err != nil {
		fmt.Println("Invalid integer")
		return
	}

	for i := 0; i < input; i++ {
		fmt.Printf("%*s", (input-i-1)*2, "")
		for j := 0; j <= i; j++ {
			value := factorial(i) / (factorial(j) * factorial(i-j))
			fmt.Printf("%4d", value)
		}
		fmt.Println()
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
