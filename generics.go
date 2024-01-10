package main

import "fmt"

	func getLast[T any](slice []T) T {
		if len(slice) == 0 {
			var myZero T
			return myZero
		}

		return slice[len(slice)-1]
	}

// don't edit below this line

	type email struct {
		message        string
		senderEmail    string
		recipientEmail string
	}

	type payment struct {
		amount         int
		senderEmail    string
		recipientEmail string
	}

	func main() {
		test([]email{}, "email")
		test([]email{
			{
				"Hi Margo",
				"janet@example.com",
				"margo@example.com",
			},
			{
				"Hey Margo I really wanna chat",
				"janet@example.com",
				"margo@example.com",
			},
			{
				"ANSWER ME",
				"janet@example.com",
				"margo@example.com",
			},
		}, "email")
		test([]payment{
			{
				5,
				"jane@example.com",
				"sally@example.com",
			},
			{
				25,
				"jane@example.com",
				"mark@example.com",
			},
			{
				1,
				"jane@example.com",
				"sally@example.com",
			},
			{
				16,
				"jane@example.com",
				"margo@example.com",
			},
		}, "payment")
	}

	func test[T any](s []T, desc string) {
		last := getLast(s)
		fmt.Printf("Getting last %v from slice of length: %v\n", desc, len(s))
		for i, v := range s {
			fmt.Printf("Item #%v: %v\n", i+1, v)
		}
		fmt.Printf("Last item in list: %v\n", last)
		fmt.Println(" --- ")
	}
*/
package main

import (
	"errors"
	"fmt"
	"time"
)

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	if balance < newItem.GetCost() {
		return oldItems, balance, errors.New("insufficient funds")
	}
	oldItems = append(oldItems, newItem)
	balance -= newItem.GetCost()
	return oldItems, balance, nil
}

// don't edit below this line

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}

func main() {
	test(subscription{
		userEmail: "john@example.com",
		startDate: time.Now().UTC(),
		interval:  "yearly",
	},
		[]subscription{},
		1000.00,
	)
	test(subscription{
		userEmail: "jane@example.com",
		startDate: time.Now().UTC(),
		interval:  "monthly",
	},
		[]subscription{
			{
				userEmail: "jane@example.com",
				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7),
				interval:  "monthly",
			},
			{
				userEmail: "jane@example.com",
				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7 * 52 * 2),
				interval:  "yearly",
			},
		},
		686.20,
	)
	test(oneTimeUsagePlan{
		userEmail:        "dillon@example.com",
		numEmailsAllowed: 5000,
	},
		[]oneTimeUsagePlan{},
		756.20,
	)
	test(oneTimeUsagePlan{
		userEmail:        "dalton@example.com",
		numEmailsAllowed: 100000,
	},
		[]oneTimeUsagePlan{
			{
				userEmail:        "dalton@example.com",
				numEmailsAllowed: 34200,
			},
		},
		32.20,
	)
}

func test[T lineItem](newItem T, oldItems []T, balance float64) {
	fmt.Println(" --- ")
	fmt.Printf("Charging customer for a '%s', current balance is %v...\n", newItem.GetName(), balance)
	newItems, newBalance, err := chargeForLineItem(newItem, oldItems, balance)
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
		return
	}
	fmt.Printf("New balance is: %v. Total number of line items is now %v\n", newBalance, len(newItems))
}
