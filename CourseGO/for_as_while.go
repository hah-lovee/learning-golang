package coursego

import (
	"fmt"
	"math/rand"
	"time"
)

func while() {

	// i := 1
	// for i <= 5 {

	// 	fmt.Println("Iteration: ", i)
	// 	i++
	// }

	// sum := 0
	// for {
	// 	sum += 10
	// 	fmt.Println("Sum:", sum)
	// 	if sum >= 50 {
	// 		break
	// 	}
	// }

	// num := 1
	// for num <= 10 {
	// 	if num%2 == 0 {
	// 		num++
	// 		continue
	// 	}
	// 	fmt.Println("Odd Number: ", num)
	// 	num++

	// }

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1

	fmt.Println("Welcom to the game!")
	fmt.Println("can you guess what it is?")

	var guess int

	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("congr!")
			break
		} else if guess < target {
			fmt.Println("too low")
		} else if guess > target {
			fmt.Println("too high")
		}

	}
}
