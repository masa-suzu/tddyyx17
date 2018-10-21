package fizzbuzz

import (
	"fmt"
	"strconv"
)

func answer(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}

func print(n int) {
	fmt.Print(answer(n))
}

func printall(max int) {
	for i := 1; i <= max; i++ {
		print(i)
		fmt.Println("")
	}
}
