/*
package main

import "fmt"

	type Message struct {
		Recipient string
		Text      string
	}

// Don't touch above this line

	func sendMessage(m Message) {
		fmt.Printf("To: %v\n", m.Recipient)
		fmt.Printf("Message: %v\n", m.Text)
	}

// Don't touch below this line

	func test(recipient string, text string) {
		m := Message{Recipient: recipient, Text: text}
		sendMessage(m)
		fmt.Println("=====================================")
	}

	func main() {
		test("Lane", "Textio is getting better everyday!")
		test("Allan", "This pointer stuff is weird...")
		test("Tiffany", "What time will you be home for dinner?")
	}

package main

import (

	"fmt"
	"strings"

)

	func removeProfanity(message *string) {
		filtered := *message
		filtered = strings.ReplaceAll(filtered, "dang", "****")
		filtered = strings.ReplaceAll(filtered, "shoot", "*****")
		filtered = strings.ReplaceAll(filtered, "heck", "****")
		*message = filtered
	}

// don't touch below this line

	func test(messages []string) {
		for _, message := range messages {
			removeProfanity(&message)
			fmt.Println(message)
		}
	}

	func main() {
		messages1 := []string{
			"well shoot, this is awful",
			"dang robots",
			"dang them to heck",
		}

		messages2 := []string{
			"well shoot",
			"Allan is going straight to heck",
			"dang... that's a tough break",
		}

		test(messages1)
		test(messages2)
	}

package main

import (

	"fmt"
	"strings"

)

	func removeProfanity(message *string) {
		if message == nil {
			return
		}
		messageVal := *message
		messageVal = strings.ReplaceAll(messageVal, "dang", "****")
		messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
		messageVal = strings.ReplaceAll(messageVal, "heck", "****")
		*message = messageVal
	}

// don't touch below this line

	func test(messages []string) {
		for _, message := range messages {
			if message == "" {
				removeProfanity(nil)
				fmt.Println("nil message detected")
			} else {
				removeProfanity(&message)
				fmt.Println(message)
			}
		}
	}

	func main() {
		messages := []string{
			"well shoot, this is awful",
			"",
			"dang robots",
			"dang them to heck",
			"",
		}

		messages2 := []string{
			"well shoot",
			"",
			"Allan is going straight to heck",
			"dang... that's a tough break",
			"",
		}

		test(messages)
		test(messages2)

}
*/
package main

import (
	"fmt"
)

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

// don't edit below this line

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func test(e *email, newMessage string) {
	fmt.Println("-- before --")
	e.print()
	fmt.Println("-- end before --")
	e.setMessage(newMessage)
	fmt.Println("-- after --")
	e.print()
	fmt.Println("-- end after --")
	fmt.Println("==========================")
}

func (e email) print() {
	fmt.Println("message:", e.message)
	fmt.Println("fromAddress:", e.fromAddress)
	fmt.Println("toAddress:", e.toAddress)
}

func main() {
	test(&email{
		message:     "this is my first draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my second draft")

	test(&email{
		message:     "this is my third draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my fourth draft")

}
