/*
package main

import (

	"fmt"
	"time"

)

	func sendEmail(message string) {
		go func() {
			time.Sleep(time.Millisecond * 250)
			fmt.Printf("Email received: '%s'\n", message)
		}()
		fmt.Printf("Email sent: '%s'\n", message)
	}

// Don't touch below this line

	func test(message string) {
		sendEmail(message)
		time.Sleep(time.Millisecond * 500)
		fmt.Println("========================")
	}

	func main() {
		test("Hello there Stacy!")
		test("Hi there John!")
		test("Hey there Jane!")
	}

package main

import (

	"fmt"
	"time"

)

	func filterOldEmails(emails []email) {
		isOldChan := make(chan bool)

		go sendIsOld(isOldChan, emails)

		isOld := <-isOldChan
		fmt.Println("email 1 is old:", isOld)
		isOld = <-isOldChan
		fmt.Println("email 2 is old:", isOld)
		isOld = <-isOldChan
		fmt.Println("email 3 is old:", isOld)
	}

// TEST SUITE -- Don't touch below this line

	func sendIsOld(isOldChan chan<- bool, emails []email) {
		for _, e := range emails {
			if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
				isOldChan <- true
				continue
			}
			isOldChan <- false
		}
	}

	type email struct {
		body string
		date time.Time
	}

	func test(emails []email) {
		filterOldEmails(emails)
		fmt.Println("==========================================")
	}

	func main() {
		test([]email{
			{
				body: "Are you going to make it?",
				date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "I need a break",
				date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "What were you thinking?",
				date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
			},
		})
		test([]email{
			{
				body: "Yo are you okay?",
				date: time.Date(2018, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "Have you heard of that website Boot.dev?",
				date: time.Date(2017, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "It's awesome honestly.",
				date: time.Date(2016, 0, 0, 0, 0, 0, 0, time.UTC),
			},
		})
		test([]email{
			{
				body: "Today is the day!",
				date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "What do you want for lunch?",
				date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "Why are you the way that you are?",
				date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
			},
		})
		test([]email{
			{
				body: "Did we do it?",
				date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "Letsa Go!",
				date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "Okay...?",
				date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
			},
		})
	}

package main

import (

	"fmt"
	"time"

)

	func waitForDbs(numDBs int, dbChan chan struct{}) {
		for i := 0; i < numDBs; i++ {
			<-dbChan
		}
	}

// don't touch below this line

	func test(numDBs int) {
		dbChan := getDatabasesChannel(numDBs)
		fmt.Printf("Waiting for %v databases...\n", numDBs)
		waitForDbs(numDBs, dbChan)
		time.Sleep(time.Millisecond * 10) // ensure the last print statement happens
		fmt.Println("All databases are online!")
		fmt.Println("=====================================")
	}

	func main() {
		test(3)
		test(4)
		test(5)
	}

	func getDatabasesChannel(numDBs int) chan struct{} {
		ch := make(chan struct{})
		go func() {
			for i := 0; i < numDBs; i++ {
				ch <- struct{}{}
				fmt.Printf("Database %v is online\n", i+1)
			}
		}()
		return ch
	}

package main

import "fmt"

	func addEmailsToQueue(emails []string) chan string {
		emailsToSend := make(chan string, len(emails))
		for _, email := range emails {
			emailsToSend <- email
		}
		return emailsToSend
	}

// TEST SUITE - Don't Touch Below This Line

	func sendEmails(batchSize int, ch chan string) {
		for i := 0; i < batchSize; i++ {
			email := <-ch
			fmt.Println("Sending email:", email)
		}
	}

	func test(emails ...string) {
		fmt.Printf("Adding %v emails to queue...\n", len(emails))
		ch := addEmailsToQueue(emails)
		fmt.Println("Sending emails...")
		sendEmails(len(emails), ch)
		fmt.Println("==========================================")
	}

	func main() {
		test("Hello John, tell Kathy I said hi", "Whazzup bruther")
		test("I find that hard to believe.", "When? I don't know if I can", "What time are you thinking?")
		test("She says hi!", "Yeah its tomorrow. So we're good.", "Cool see you then!", "Bye!")
	}

package main

import (

	"fmt"
	"time"

)

	func countReports(numSentCh chan int) int {
		total := 0
		for {
			numSent, ok := <-numSentCh
			if !ok {
				break
			}
			total += numSent
		}

		return total
	}

// don't touch below this line

	func test(numBatches int) {
		numSentCh := make(chan int)
		go sendReports(numBatches, numSentCh)

		fmt.Println("Start counting...")
		numReports := countReports(numSentCh)
		fmt.Printf("%v reports sent!\n", numReports)
		fmt.Println("========================")
	}

	func main() {
		test(3)
		test(4)
		test(5)
		test(6)
	}

	func sendReports(numBatches int, ch chan int) {
		for i := 0; i < numBatches; i++ {
			numReports := i*23 + 32%17
			ch <- numReports
			fmt.Printf("Sent batch of %v reports\n", numReports)
			time.Sleep(time.Millisecond * 100)
		}
		close(ch)
	}

package main

import (

	"fmt"
	"time"

)

	func concurrentFib(n int) {
		numCh := make(chan int)
		go func() {
			fibonacci(n, numCh)
		}()

		for v := range numCh {
			fmt.Println(v)
		}
	}

// don't touch below this line

	func test(n int) {
		fmt.Printf("Printing %v numbers...\n", n)
		concurrentFib(n)
		fmt.Println("==============================")
	}

	func main() {
		test(10)
		test(5)
		test(20)
		test(13)
	}

	func fibonacci(n int, ch chan int) {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			ch <- x
			x, y = y, x+y
			time.Sleep(time.Millisecond * 10)
		}
		close(ch)
	}

package main

import (

	"fmt"
	"math/rand"
	"time"

)

	func logMessages(chEmails, chSms chan string) {
		for {
			select {
			case email, ok := <-chEmails:
				if !ok {
					return
				}
				logSms(email)
			case sms, ok := <-chSms:
				if !ok {
					return
				}
				logEmail(sms)
			}
		}
	}

// don't touch below this line

	func logSms(sms string) {
		fmt.Println("SMS:", sms)
	}

	func logEmail(email string) {
		fmt.Println("Email:", email)
	}

	func test(sms []string, emails []string) {
		fmt.Println("Starting...")

		chSms, chEmails := sendToLogger(sms, emails)

		logMessages(chEmails, chSms)
		fmt.Println("===============================")
	}

	func main() {
		rand.Seed(0)
		test(
			[]string{
				"hi friend",
				"What's going on?",
				"Welcome to the business",
				"I'll pay you to be my friend",
			},
			[]string{
				"Will you make your appointment?",
				"Let's be friends",
				"What are you doing?",
				"I can't believe you've done this.",
			},
		)
		test(
			[]string{
				"this song slaps hard",
				"yooo hoooo",
				"i'm a big fan",
			},
			[]string{
				"What do you think of this song?",
				"I hate this band",
				"Can you believe this song?",
			},
		)
	}

	func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
		chSms = make(chan string)
		chEmails = make(chan string)
		go func() {
			for i := 0; i < len(sms) && i < len(emails); i++ {
				done := make(chan struct{})
				s := sms[i]
				e := emails[i]
				t1 := time.Millisecond * time.Duration(rand.Intn(1000))
				t2 := time.Millisecond * time.Duration(rand.Intn(1000))
				go func() {
					time.Sleep(t1)
					chSms <- s
					done <- struct{}{}
				}()
				go func() {
					time.Sleep(t2)
					chEmails <- e
					done <- struct{}{}
				}()
				<-done
				<-done
				time.Sleep(10 * time.Millisecond)
			}
			close(chSms)
			close(chEmails)
		}()
		return chSms, chEmails
	}

package main

import (

	"fmt"
	"time"

)

	func saveBackups(snapshotTicker, saveAfter <-chan time.Time) {
		for {
			select {
			case <-snapshotTicker:
				takeSnapshot()
			case <-saveAfter:
				saveSnapshot()
				return
			default:
				waitForData()
				time.Sleep(time.Millisecond * 500)
			}
		}
	}

// don't touch below this line

	func takeSnapshot() {
		fmt.Println("Taking a backup snapshot...")
	}

	func saveSnapshot() {
		fmt.Println("All backups saved!")
	}

	func waitForData() {
		fmt.Println("Nothing to do, waiting...")
	}

	func test() {
		snapshotTicker := time.Tick(800 * time.Millisecond)
		saveAfter := time.After(2800 * time.Millisecond)
		saveBackups(snapshotTicker, saveAfter)
		fmt.Println("===========================")
	}

	func main() {
		test()
	}

package main

import (
	"fmt"
	"time"
)

func pingPong(numPings int) {
	pings := make(chan struct{})
	pongs := make(chan struct{})
	go ponger(pings, pongs)
	go pinger(pings, pongs, numPings)

	time.Sleep(time.Duration(numPings) * 100 * time.Millisecond)
}

// don't touch below this line

func pinger(pings, pongs chan struct{}, numPings int) {
	go func() {
		sleepTime := 50 * time.Millisecond
		for i := 0; i < numPings; i++ {
			fmt.Printf("sending ping %v\n", i)
			pings <- struct{}{}
			time.Sleep(sleepTime)
			sleepTime *= 2
		}
		close(pings)
	}()
	i := 0
	for range pongs {
		fmt.Println("got pong", i)
		i++
	}
	fmt.Println("pongs done")
}

func ponger(pings, pongs chan struct{}) {
	i := 0
	for range pings {
		fmt.Printf("got ping %v, sending pong %v\n", i, i)
		pongs <- struct{}{}
		i++
	}
	fmt.Println("pings done")
	close(pongs)
}

func test(numPings int) {
	fmt.Println("Starting game...")
	pingPong(numPings)
	fmt.Println("===== Game over =====")
}

func main() {
	test(4)
	test(3)
	test(2)
	test(100)
}
*/