/*
package main

import (

	"errors"
	"fmt"

)

	func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
		users := make(map[string]user)
		if len(names) != len(phoneNumbers) {
			return users, errors.New("invalid sizes")
		}
		for i, name := range names {
			users[name] = user {
				name,
				phoneNumbers[i],
			}
		}
		return users, nil
	}

// don't touch below this line

	type user struct {
		name        string
		phoneNumber int
	}

	func test(names []string, phoneNumbers []int) {
		fmt.Println("Creating map...")
		defer fmt.Println("====================================")
		users, err := getUserMap(names, phoneNumbers)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, name := range names {
			fmt.Printf("key: %v, value:\n", name)
			fmt.Println(" - name:", users[name].name)
			fmt.Println(" - number:", users[name].phoneNumber)
		}
	}

	func main() {
		test(
			[]string{"John", "Bob", "Jill"},
			[]int{14355550987, 98765550987, 18265554567},
		)
		test(
			[]string{"John", "Bob"},
			[]int{14355550987, 98765550987, 18265554567},
		)
		test(
			[]string{"George", "Sally", "Rich", "Sue"},
			[]int{20955559812, 38385550982, 48265554567, 16045559873},
		)
	}

package main

import (

	"errors"
	"fmt"
	"sort"

)

	func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
		user, ok := users[name]
		if !ok {
			return false, errors.New("not found")
		}

		if user.scheduledForDeletion {
			delete(users, name)
			return true, nil
		}

		return false, nil
	}

// don't touch below this line

	type user struct {
		name                 string
		number               int
		scheduledForDeletion bool
	}

	func test(users map[string]user, name string) {
		fmt.Printf("Attempting to delete %s...\n", name)
		defer fmt.Println("====================================")
		deleted, err := deleteIfNecessary(users, name)
		if err != nil {
			fmt.Println(err)
			return
		}
		if deleted {
			fmt.Println("Deleted:", name)
			return
		}
		fmt.Println("Did not delete:", name)
	}

	func main() {
		users := map[string]user{
			"john": {
				name:                 "john",
				number:               18965554631,
				scheduledForDeletion: true,
			},
			"elon": {
				name:                 "elon",
				number:               19875556452,
				scheduledForDeletion: true,
			},
			"breanna": {
				name:                 "breanna",
				number:               98575554231,
				scheduledForDeletion: false,
			},
			"kade": {
				name:                 "kade",
				number:               10765557221,
				scheduledForDeletion: false,
			},
		}
		test(users, "john")
		test(users, "musk")
		test(users, "santa")
		test(users, "kade")

		keys := []string{}
		for name := range users {
			keys = append(keys, name)
		}
		sort.Strings(keys)

		fmt.Println("Final map keys:")
		for _, name := range keys {
			fmt.Println(" - ", name)
		}
	}

package main

import (

	"crypto/md5"
	"fmt"
	"io"

)

	func getCounts(userIDs []string) map[string]int {
		counts := make(map[string]int)
		for _, userID := range userIDs {
			count := counts[userID]
			count++
			counts[userID] = count
		}

		return counts
	}

// don't edit below this line

	func test(userIDs []string, ids []string) {
		fmt.Printf("Generating counts for %v user IDs...\n", len(userIDs))

		counts := getCounts(userIDs)
		fmt.Println("Counts from select IDs:")
		for _, k := range ids {
			v := counts[k]
			fmt.Printf(" - %s: %d\n", k, v)
		}
		fmt.Println("=====================================")
	}

	func main() {
		userIDs := []string{}
		for i := 0; i < 10000; i++ {
			h := md5.New()
			io.WriteString(h, fmt.Sprint(i))
			key := fmt.Sprintf("%x", h.Sum(nil))
			userIDs = append(userIDs, key[:2])
		}

		test(userIDs, []string{"00", "ff", "dd"})
		test(userIDs, []string{"aa", "12", "32"})
		test(userIDs, []string{"bb", "33"})
	}

package main

import (
	"fmt"
)

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		initial := rune(name[0])
		if _, ok := counts[initial]; !ok {
			counts[initial] = make(map[string]int)
		}

		counts[initial][name]++
	}
	return counts
}

// don't edit below this line

func test(names []string, initial rune, name string) {
	fmt.Printf("Generating counts for %v names...\n", len(names))

	nameCounts := getNameCounts(names)
	count := nameCounts[initial][name]
	fmt.Printf("Count for [%c][%s]: %d\n", initial, name, count)
	fmt.Println("=====================================")
}

func main() {
	test(getNames(50), 'M', "Matthew")
	test(getNames(100), 'G', "George")
	test(getNames(150), 'D', "Drew")
	test(getNames(200), 'P', "Philip")
	test(getNames(250), 'B', "Bryant")
	test(getNames(300), 'M', "Matthew")
}

func getNames(length int) []string {
	names := []string{
		"Grant", "Eduardo", "Peter", "Matthew", "Matthew", "Matthew", "Peter", "Peter", "Henry", "Parker", "Parker", "Parker", "Collin", "Hayden", "George", "Bradley", "Mitchell", "Devon", "Ricardo", "Shawn", "Taylor", "Nicolas", "Gregory", "Francisco", "Liam", "Kaleb", "Preston", "Erik", "Alexis", "Owen", "Omar", "Diego", "Dustin", "Corey", "Fernando", "Clayton", "Carter", "Ivan", "Jaden", "Javier", "Alec", "Johnathan", "Scott", "Manuel", "Cristian", "Alan", "Raymond", "Brett", "Max", "Drew", "Andres", "Gage", "Mario", "Dawson", "Dillon", "Cesar", "Wesley", "Levi", "Jakob", "Chandler", "Martin", "Malik", "Edgar", "Sergio", "Trenton", "Josiah", "Nolan", "Marco", "Drew", "Peyton", "Harrison", "Drew", "Hector", "Micah", "Roberto", "Drew", "Brady", "Erick", "Conner", "Jonah", "Casey", "Jayden", "Edwin", "Emmanuel", "Andre", "Phillip", "Brayden", "Landon", "Giovanni", "Bailey", "Ronald", "Braden", "Damian", "Donovan", "Ruben", "Frank", "Gerardo", "Pedro", "Andy", "Chance", "Abraham", "Calvin", "Trey", "Cade", "Donald", "Derrick", "Payton", "Darius", "Enrique", "Keith", "Raul", "Jaylen", "Troy", "Jonathon", "Cory", "Marc", "Eli", "Skyler", "Rafael", "Trent", "Griffin", "Colby", "Johnny", "Chad", "Armando", "Kobe", "Caden", "Marcos", "Cooper", "Elias", "Brenden", "Israel", "Avery", "Zane", "Zane", "Zane", "Zane", "Dante", "Josue", "Zackary", "Allen", "Philip", "Mathew", "Dennis", "Leonardo", "Ashton", "Philip", "Philip", "Philip", "Julio", "Miles", "Damien", "Ty", "Gustavo", "Drake", "Jaime", "Simon", "Jerry", "Curtis", "Kameron", "Lance", "Brock", "Bryson", "Alberto", "Dominick", "Jimmy", "Kaden", "Douglas", "Gary", "Brennan", "Zachery", "Randy", "Louis", "Larry", "Nickolas", "Albert", "Tony", "Fabian", "Keegan", "Saul", "Danny", "Tucker", "Myles", "Damon", "Arturo", "Corbin", "Deandre", "Ricky", "Kristopher", "Lane", "Pablo", "Darren", "Jarrett", "Zion", "Alfredo", "Micheal", "Angelo", "Carl", "Oliver", "Kyler", "Tommy", "Walter", "Dallas", "Jace", "Quinn", "Theodore", "Grayson", "Lorenzo", "Joe", "Arthur", "Bryant", "Roman", "Brent", "Russell", "Ramon", "Lawrence", "Moises", "Aiden", "Quentin", "Jay", "Tyrese", "Tristen", "Emanuel", "Salvador", "Terry", "Morgan", "Jeffery", "Esteban", "Tyson", "Braxton", "Branden", "Marvin", "Brody", "Craig", "Ismael", "Rodney", "Isiah", "Marshall", "Maurice", "Ernesto", "Emilio", "Brendon", "Kody", "Eddie", "Malachi", "Abel", "Keaton", "Jon", "Shaun", "Skylar", "Ezekiel", "Nikolas", "Santiago", "Kendall", "Axel", "Camden", "Trevon", "Bobby", "Conor", "Jamal", "Lukas", "Malcolm", "Zackery", "Jayson", "Javon", "Roger", "Reginald", "Zachariah", "Desmond", "Felix", "Johnathon", "Dean", "Quinton", "Ali", "Davis", "Gerald", "Rodrigo", "Demetrius", "Billy", "Rene", "Reece", "Kelvin", "Leo", "Justice", "Chris", "Guillermo", "Matthew", "Matthew", "Matthew", "Kevon", "Steve", "Frederick", "Clay", "Weston", "Dorian", "Hugo", "Roy", "Orlando", "Terrance", "Kai", "Khalil", "Khalil", "Khalil", "Graham", "Noel", "Willie", "Nathanael", "Terrell", "Tyrone",
	}
	if length > len(names) {
		length = len(names)
	}
	return names[:length]
}
*/