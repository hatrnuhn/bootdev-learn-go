/*
package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

// don't touch below this line

func main() {
	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}
*/

/*
package main

import "fmt"

func main() {
	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}
*/

/*
package main

import "fmt"

func main() {
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio,", firstName)
}

// don't edit below this line

func getNames() (string, string) {
	return "John", "Doe"
}
*/

/*
package main

import (
	"fmt"
)

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}

// don't edit below this line

func test(age int) {
	fmt.Println("Age:", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)
	fmt.Println("You are an adult in", yearsUntilAdult, "years")
	fmt.Println("You can drink in", yearsUntilDrinking, "years")
	fmt.Println("You can rent a car in", yearsUntilCarRental, "years")
	fmt.Println("====================================")
}

func main() {
	test(4)
	test(10)
	test(22)
	test(35)
}
*/

/*
package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

// don't edit below this line

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("====================================")
}

func main() {
	test(messageToSend{
		phoneNumber: 148255510981,
		message:     "Thanks for signing up",
	})
	test(messageToSend{
		phoneNumber: 148255510982,
		message:     "Love to have you aboard!",
	})
	test(messageToSend{
		phoneNumber: 148255510983,
		message:     "We're so excited to have you",
	})
}
*/

/*
package main

import (

	"fmt"

)

	type messageToSend struct {
		message   string
		sender    user
		recipient user
	}

	type user struct {
		name   string
		number int
	}

	func canSendMessage(mToSend messageToSend) bool {
		if mToSend.sender.name == "" {
			return false
		}
		if mToSend.recipient.name == "" {
			return false
		}
		if mToSend.sender.number == 0 {
			return false
		}
		if mToSend.recipient.number == 0 {
			return false
		}
		return true
	}

// don't touch below this line

	func test(mToSend messageToSend) {
		fmt.Printf(`sending "%s" from %s (%v) to %s (%v)...`,
			mToSend.message,
			mToSend.sender.name,
			mToSend.sender.number,
			mToSend.recipient.name,
			mToSend.recipient.number,
		)
		fmt.Println()
		if canSendMessage(mToSend) {
			fmt.Println("...sent!")
		} else {
			fmt.Println("...can't send message")
		}
		fmt.Println("====================================")
	}

	func main() {
		test(messageToSend{
			message: "you have an appointment tommorow",
			sender: user{
				name:   "Brenda Halafax",
				number: 16545550987,
			},
			recipient: user{
				name:   "Sally Sue",
				number: 19035558973,
			},
		})
		test(messageToSend{
			message: "you have an event tommorow",
			sender: user{
				number: 16545550987,
			},
			recipient: user{
				name:   "Suzie Sall",
				number: 19035558973,
			},
		})
		test(messageToSend{
			message: "you have an party tommorow",
			sender: user{
				name: "Njorn Halafax",
			},
			recipient: user{
				name:   "Becky Sue",
				number: 19035558973,
			},
		})
		test(messageToSend{
			message: "you have a birthday tommorow",
			sender: user{
				name:   "Eli Halafax",
				number: 16545550987,
			},
			recipient: user{
				number: 19035558973,
			},
		})
		test(messageToSend{
			message: "you have a birthday tommorow",
			sender: user{
				name:   "Jason Bjorn",
				number: 16545550987,
			},
			recipient: user{
				name: "Jim Bond",
			},
		})
	}

package main

import "fmt"

	type sender struct {
		rateLimit int
		user
	}

	type user struct {
		name   string
		number int
	}

// don't edit below this line

	func test(s sender) {
		fmt.Println("Sender name:", s.name)
		fmt.Println("Sender number:", s.number)
		fmt.Println("Sender rateLimit:", s.rateLimit)
		fmt.Println("====================================")
	}

	func main() {
		test(sender{
			rateLimit: 10000,
			user: user{
				name:   "Deborah",
				number: 18055558790,
			},
		})
		test(sender{
			rateLimit: 5000,
			user: user{
				name:   "Sarah",
				number: 19055558790,
			},
		})
		test(sender{
			rateLimit: 1000,
			user: user{
				name:   "Sally",
				number: 19055558790,
			},
		})
	}

package main

import "fmt"

	type authenticationInfo struct {
		username string
		password string
	}

	func (authInfo authenticationInfo) getBasicAuth() string {
		return fmt.Sprintf(
			"Authorization: Basic %s:%s",
			authInfo.username,
			authInfo.password,
		)
	}

// don't touch below this line

	func test(authInfo authenticationInfo) {
		fmt.Println(authInfo.getBasicAuth())
		fmt.Println("====================================")
	}

	func main() {
		test(authenticationInfo{
			username: "Google",
			password: "12345",
		})
		test(authenticationInfo{
			username: "Bing",
			password: "98765",
		})
		test(authenticationInfo{
			username: "DDG",
			password: "76921",
		})
	}

package main

import (

	"fmt"
	"time"

)

	func sendMessage(msg message) {
		fmt.Println(msg.getMessage())
	}

	type message interface {
		getMessage() string
	}

// don't edit below this line

	type birthdayMessage struct {
		birthdayTime  time.Time
		recipientName string
	}

	func (bm birthdayMessage) getMessage() string {
		return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
	}

	type sendingReport struct {
		reportName    string
		numberOfSends int
	}

	func (sr sendingReport) getMessage() string {
		return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
	}

	func test(m message) {
		sendMessage(m)
		fmt.Println("====================================\n")
	}

	func main() {
		test(sendingReport{
			reportName:    "First Report",
			numberOfSends: 10,
		})
		test(birthdayMessage{
			recipientName: "John Doe",
			birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
		})
		test(sendingReport{
			reportName:    "First Report",
			numberOfSends: 10,
		})
		test(birthdayMessage{
			recipientName: "Bill Deer",
			birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
		})
	}

package main

import (

	"fmt"

)

	type employee interface {
		getName() string
		getSalary() int
	}

	type contractor struct {
		name         string
		hourlyPay    int
		hoursPerYear int
	}

	func (c contractor) getName() string {
		return c.name
	}

	func (c contractor) getSalary() int {
		return c.hourlyPay * c.hoursPerYear
	}

// don't touch below this line

	type fullTime struct {
		name   string
		salary int
	}

	func (ft fullTime) getSalary() int {
		return ft.salary
	}

	func (ft fullTime) getName() string {
		return ft.name
	}

	func test(e employee) {
		fmt.Println(e.getName(), e.getSalary())
		fmt.Println("====================================")
	}

	func main() {
		test(fullTime{
			name:   "Jack",
			salary: 50000,
		})
		test(contractor{
			name:         "Bob",
			hourlyPay:    100,
			hoursPerYear: 73,
		})
		test(contractor{
			name:         "Jill",
			hourlyPay:    872,
			hoursPerYear: 982,
		})
	}

package main

import (

	"fmt"

)

	func (e email) cost() float64 {
		if e.isSubscribed {
			return 0.05 * float64(len(e.body))
		}

		return 0.01 * float64(len(e.body))
	}

	func (e email) print() {
		fmt.Println(e.body)
	}

// don't touch below this line

	type expense interface {
		cost() float64
	}

	type printer interface {
		print()
	}

	type email struct {
		isSubscribed bool
		body         string
	}

	func print(p printer) {
		p.print()
	}

	func test(e expense, p printer) {
		fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
		p.print()
		fmt.Println("====================================\n")
	}

	func main() {
		e := email{
			isSubscribed: true,
			body:         "hello there",
		}
		test(e, e)
		e = email{
			isSubscribed: false,
			body:         "I want my money back",
		}
		test(e, e)
		e = email{
			isSubscribed: true,
			body:         "Are you free for a chat?",
		}
		test(e, e)
		e = email{
			isSubscribed: false,
			body:         "This meeting could have been an email",
		}
		test(e, e)
	}

package main

import (

	"fmt"

)

	func getExpenseReport(e expense) (string, float64) {
		em, ok := e.(email)
		if ok {
			return em.toAddress, em.cost()
		}

		s, ok := e.(sms)
		if ok {
			return s.toPhoneNumber, s.cost()
		}

		return "", 0.0
	}

// don't touch below this line

	func (e email) cost() float64 {
		if !e.isSubscribed {
			return float64(len(e.body)) * .05
		}
		return float64(len(e.body)) * .01
	}

	func (s sms) cost() float64 {
		if !s.isSubscribed {
			return float64(len(s.body)) * .1
		}
		return float64(len(s.body)) * .03
	}

	func (i invalid) cost() float64 {
		return 0.0
	}

	type expense interface {
		cost() float64
	}

	type email struct {
		isSubscribed bool
		body         string
		toAddress    string
	}

	type sms struct {
		isSubscribed  bool
		body          string
		toPhoneNumber string
	}

type invalid struct{}

	func estimateYearlyCost(e expense, averageMessagesPerYear int) float64 {
		return e.cost() * float64(averageMessagesPerYear)
	}

	func test(e expense) {
		address, cost := getExpenseReport(e)
		switch e.(type) {
		case email:
			fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
			fmt.Println("====================================")
		case sms:
			fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
			fmt.Println("====================================")
		default:
			fmt.Println("Report: Invalid expense")
			fmt.Println("====================================")
		}
	}

	func main() {
		test(email{
			isSubscribed: true,
			body:         "hello there",
			toAddress:    "john@does.com",
		})
		test(email{
			isSubscribed: false,
			body:         "This meeting could have been an email",
			toAddress:    "jane@doe.com",
		})
		test(email{
			isSubscribed: false,
			body:         "This meeting could have been an email",
			toAddress:    "elon@doe.com",
		})
		test(sms{
			isSubscribed:  false,
			body:          "This meeting could have been an email",
			toPhoneNumber: "+155555509832",
		})
		test(sms{
			isSubscribed:  false,
			body:          "This meeting could have been an email",
			toPhoneNumber: "+155555504444",
		})
		test(invalid{})
	}

package main

import (

	"fmt"

)

	func getExpenseReport(e expense) (string, float64) {
		switch ex := e.(type) {
		case sms:
			return ex.toPhoneNumber, ex.cost()
		case email:
			return ex.toAddress, ex.cost()
		default:
			return "", 0.0
		}
	}

// don't touch below this line

	func (e email) cost() float64 {
		if !e.isSubscribed {
			return float64(len(e.body)) * .05
		}
		return float64(len(e.body)) * .01
	}

	func (s sms) cost() float64 {
		if !s.isSubscribed {
			return float64(len(s.body)) * .1
		}
		return float64(len(s.body)) * .03
	}

	func (i invalid) cost() float64 {
		return 0.0
	}

	type expense interface {
		cost() float64
	}

	type email struct {
		isSubscribed bool
		body         string
		toAddress    string
	}

	type sms struct {
		isSubscribed  bool
		body          string
		toPhoneNumber string
	}

type invalid struct{}

	func estimateYearlyCost(e expense, averageMessagesPerYear int) float64 {
		return e.cost() * float64(averageMessagesPerYear)
	}

	func test(e expense) {
		address, cost := getExpenseReport(e)
		switch e.(type) {
		case email:
			fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
			fmt.Println("====================================")
		case sms:
			fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
			fmt.Println("====================================")
		default:
			fmt.Println("Report: Invalid expense")
			fmt.Println("====================================")
		}
	}

	func main() {
		test(email{
			isSubscribed: true,
			body:         "hello there",
			toAddress:    "john@does.com",
		})
		test(email{
			isSubscribed: false,
			body:         "This meeting could have been an email",
			toAddress:    "jane@doe.com",
		})
		test(email{
			isSubscribed: false,
			body:         "Wanna catch up later?",
			toAddress:    "elon@doe.com",
		})
		test(sms{
			isSubscribed:  false,
			body:          "I'm a Nigerian prince, please send me your bank info so I can deposit $1000 dollars",
			toPhoneNumber: "+155555509832",
		})
		test(sms{
			isSubscribed:  false,
			body:          "I don't need this",
			toPhoneNumber: "+155555504444",
		})
		test(invalid{})
	}

package main

import (

	"fmt"

)

	func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
		costForCustomer, err := sendSMS(msgToCustomer)
		if err != nil {
			return 0.0, err
		}

		costForSpouse, err := sendSMS(msgToSpouse)
		if err != nil {
			return 0.0, err
		}

		return costForCustomer + costForSpouse, nil
	}

// don't edit below this line

	func sendSMS(message string) (float64, error) {
		const maxTextLen = 25
		const costPerChar = .0002
		if len(message) > maxTextLen {
			return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
		}
		return costPerChar * float64(len(message)), nil
	}

	func test(msgToCustomer, msgToSpouse string) {
		defer fmt.Println("========")
		fmt.Println("Message for customer:", msgToCustomer)
		fmt.Println("Message for spouse:", msgToSpouse)
		totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Total cost: $%.4f\n", totalCost)
	}

	func main() {
		test(
			"Thanks for coming in to our flower shop today!",
			"We hope you enjoyed your gift.",
		)
		test(
			"Thanks for joining us!",
			"Have a good day.",
		)
		test(
			"Thank you.",
			"Enjoy!",
		)
		test(
			"We loved having you in!",
			"We hope the rest of your evening is absolutely fantastic.",
		)
	}

package main

import (

	"fmt"

)

	func getSMSErrorString(cost float64, recipient string) string {
		return fmt.Sprintf(
			"SMS that costs %.2f to be sent to %s can not be sent",
			cost,
			recipient,
		)
	}

// don't edit below this line

	func test(cost float64, recipient string) {
		s := getSMSErrorString(cost, recipient)
		fmt.Println(s)
		fmt.Println("====================================")
	}

	func main() {
		test(1.4, "+1 (435) 555 0923")
		test(2.1, "+2 (702) 555 3452")
		test(32.1, "+1 (801) 555 7456")
		test(14.4, "+1 (234) 555 6545")
	}

package main

import (

	"fmt"

)

	type divideError struct {
		dividend float64
	}

	func (divErr divideError) Error() string {
		return fmt.Sprintf(
			"can not divide %.2f by zero",
			divErr.dividend,
		)
	}

// don't edit below this line

	func divide(dividend, divisor float64) (float64, error) {
		if divisor == 0 {
			// We convert the `divideError` struct to an `error` type by returning it
			// as an error. As an error type, when it's printed its default value
			// will be the result of the Error() method
			return 0, divideError{dividend: dividend}
		}
		return dividend / divisor, nil
	}

	func test(dividend, divisor float64) {
		defer fmt.Println("====================================")
		fmt.Printf("Dividing %.2f by %.2f ...\n", dividend, divisor)
		quotient, err := divide(dividend, divisor)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Quotient: %.2f\n", quotient)
	}

	func main() {
		test(10, 0)
		test(10, 2)
		test(15, 30)
		test(6, 3)
	}

package main

import (

	"errors"
	"fmt"

)

	func divide(x, y float64) (float64, error) {
		if y == 0 {
			return 0.0, errors.New("no dividing by 0")
		}
		return x / y, nil
	}

// don't edit below this line

	func test(x, y float64) {
		defer fmt.Println("====================================")
		fmt.Printf("Dividing %.2f by %.2f ...\n", x, y)
		quotient, err := divide(x, y)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Quotient: %.2f\n", quotient)
	}

	func main() {
		test(10, 0)
		test(10, 2)
		test(15, 30)
		test(6, 3)
	}

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
*/