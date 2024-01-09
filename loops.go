/*
package main

import (

	"fmt"

)

	func bulkSend(numMessages int) float64 {
		totalCost := 0.0
		for i := 0; i < numMessages; i++ {
			totalCost += 1.0 + (0.01 * float64(i))
		}
		return totalCost
	}

// don't edit below this line

	func test(numMessages int) {
		fmt.Printf("Sending %v messages\n", numMessages)
		cost := bulkSend(numMessages)
		fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
		fmt.Println("===============================================================")
	}

	func main() {
		test(10)
		test(20)
		test(30)
		test(40)
		test(50)
	}

package main

import (

	"fmt"

)

	func maxMessages(thresh float64) int {
		totalCost := 0.0
		for i := 0; ; i++ {
			totalCost += 1.0 + (0.01 * float64(i))
			if totalCost > thresh {
				return i
			}
		}
	}

// don't edit below this line

	func test(thresh float64) {
		fmt.Printf("Threshold: %.2f\n", thresh)
		max := maxMessages(thresh)
		fmt.Printf("Maximum messages that can be sent: = %v\n", max)
		fmt.Println("===============================================================")
	}

	func main() {
		test(10.00)
		test(20.00)
		test(30.00)
		test(5.1)
		test(40.00)
		test(50.00)
	}

package main

import (
	"fmt"
)

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	for actualCostInPennies <= float64(maxCostInPennies) {
		actualCostInPennies *= costMultiplier
		maxMessagesToSend++
	}
	if actualCostInPennies > float64(maxCostInPennies) {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}

// don't touch below this line

func test(costMultiplier float64, maxCostInPennies int) {
	maxMessagesToSend := getMaxMessagesToSend(costMultiplier, maxCostInPennies)
	fmt.Printf("Multiplier: %v\n",
		costMultiplier,
	)
	fmt.Printf("Max cost: %v\n",
		maxCostInPennies,
	)
	fmt.Printf("Max messages you can send: %v\n",
		maxMessagesToSend,
	)
	fmt.Println("====================================")
}

func main() {
	test(1.1, 5)
	test(1.3, 10)
	test(1.35, 25)
}

package main

import (
	"fmt"
)

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		word := ""
		if i%3 == 0 && i%5 == 0 {
			word = "fizzbuz"
		} else if i%3 == 0 {
			word = "fizz"
		} else if i%5 == 0 {
			word = "buzz"
		} else {
			word = fmt.Sprintf("%v", i)
		}

		fmt.Println(word)
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}

package main

import (
	"fmt"
)

func printPrimes(max int) {
	for i := 2; i < max+1; i++ {
		if i == 2 {
			fmt.Println(i)
		}

		if i%2 == 0 {
			continue
		}

		isPrime := true
		for j := 3; j*j < i+1; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if !isPrime {
			continue
		}

		fmt.Println(i)
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
*/