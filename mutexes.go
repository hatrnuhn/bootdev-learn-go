/*
package main

import (

	"fmt"
	"sort"
	"sync"
	"time"

)

	type safeCounter struct {
		counts map[string]int
		mu     *sync.Mutex
	}

// don't touch above this line

	func (sc safeCounter) inc(key string) {
		sc.mu.Lock()
		defer sc.mu.Unlock()
		sc.slowIncrement(key)
	}

	func (sc safeCounter) val(key string) int {
		sc.mu.Lock()
		defer sc.mu.Unlock()
		return sc.slowVal(key)
	}

// don't touch below this line

	func (sc safeCounter) slowIncrement(key string) {
		tempCounter := sc.counts[key]
		time.Sleep(time.Microsecond)
		tempCounter++
		sc.counts[key] = tempCounter
	}

	func (sc safeCounter) slowVal(key string) int {
		time.Sleep(time.Microsecond)
		return sc.counts[key]
	}

	type emailTest struct {
		email string
		count int
	}

	func test(sc safeCounter, emailTests []emailTest) {
		emails := make(map[string]struct{})

		var wg sync.WaitGroup
		for _, emailT := range emailTests {
			emails[emailT.email] = struct{}{}
			for i := 0; i < emailT.count; i++ {
				wg.Add(1)
				go func(emailT emailTest) {
					sc.inc(emailT.email)
					wg.Done()
				}(emailT)
			}
		}
		wg.Wait()

		emailsSorted := make([]string, 0, len(emails))
		for email := range emails {
			emailsSorted = append(emailsSorted, email)
		}
		sort.Strings(emailsSorted)

		for _, email := range emailsSorted {
			go sc.inc(email)
			fmt.Printf("Email: %s has %d emails\n", email, sc.val(email))
		}
		fmt.Println("=====================================")
	}

	func main() {
		sc := safeCounter{
			counts: make(map[string]int),
			mu:     &sync.Mutex{},
		}
		test(sc, []emailTest{
			{
				email: "john@example.com",
				count: 23,
			},
			{
				email: "john@example.com",
				count: 29,
			},
			{
				email: "jill@example.com",
				count: 31,
			},
			{
				email: "jill@example.com",
				count: 67,
			},
		})
		test(sc, []emailTest{
			{
				email: "kaden@example.com",
				count: 23,
			},
			{
				email: "george@example.com",
				count: 126,
			},
			{
				email: "kaden@example.com",
				count: 31,
			},
			{
				email: "george@example.com",
				count: 453,
			},
		})
	}

package main

import (

	"fmt"
	"sync"

)

	func main() {
		m := map[int]int{}

		mu := &sync.Mutex{}

		go writeLoop(m, mu)
		go readLoop(m, mu)

		// stop program from exiting, must be killed
		block := make(chan struct{})
		<-block
	}

	func writeLoop(m map[int]int, mu *sync.Mutex) {
		for {
			for i := 0; i < 100; i++ {
				mu.Lock()
				m[i] = i
				mu.Unlock()
			}
		}
	}

	func readLoop(m map[int]int, mu *sync.Mutex) {
		for {
			mu.Lock()
			for k, v := range m {
				fmt.Println(k, "-", v)
			}
			mu.Unlock()
		}
	}

package main

import (

	"fmt"
	"sort"
	"sync"
	"time"

)

	type safeCounter struct {
		counts map[string]int
		mu     *sync.RWMutex
	}

	func (sc safeCounter) inc(key string) {
		sc.mu.Lock()
		defer sc.mu.Unlock()
		sc.slowIncrement(key)
	}

	func (sc safeCounter) val(key string) int {
		sc.mu.RLock()
		defer sc.mu.RUnlock()
		return sc.counts[key]
	}

// don't touch below this line

	func (sc safeCounter) slowIncrement(key string) {
		tempCounter := sc.counts[key]
		time.Sleep(time.Microsecond)
		tempCounter++
		sc.counts[key] = tempCounter
	}

	type emailTest struct {
		email string
		count int
	}

	func test(sc safeCounter, emailTests []emailTest) {
		emails := make(map[string]struct{})

		var wg sync.WaitGroup
		for _, emailT := range emailTests {
			emails[emailT.email] = struct{}{}
			for i := 0; i < emailT.count; i++ {
				wg.Add(1)
				go func(emailT emailTest) {
					sc.inc(emailT.email)
					wg.Done()
				}(emailT)
			}
		}
		wg.Wait()

		emailsSorted := make([]string, 0, len(emails))
		for email := range emails {
			emailsSorted = append(emailsSorted, email)
		}
		sort.Strings(emailsSorted)

		sc.mu.RLock()
		defer sc.mu.RUnlock()
		for _, email := range emailsSorted {
			fmt.Printf("Email: %s has %d emails\n", email, sc.val(email))
		}
		fmt.Println("=====================================")
	}

	func main() {
		sc := safeCounter{
			counts: make(map[string]int),
			mu:     &sync.RWMutex{},
		}
		test(sc, []emailTest{
			{
				email: "john@example.com",
				count: 23,
			},
			{
				email: "john@example.com",
				count: 29,
			},
			{
				email: "jill@example.com",
				count: 31,
			},
			{
				email: "jill@example.com",
				count: 67,
			},
		})
		test(sc, []emailTest{
			{
				email: "kaden@example.com",
				count: 23,
			},
			{
				email: "george@example.com",
				count: 126,
			},
			{
				email: "kaden@example.com",
				count: 31,
			},
			{
				email: "george@example.com",
				count: 453,
			},
		})
	}

