package fizzbuzz

import "fmt"

func print(n int) {
	fmt.Print(answer(n))
}

func printall(max int) {
	for i := 1; i <= max; i++ {
		print(i)
		fmt.Println("")
	}
}
