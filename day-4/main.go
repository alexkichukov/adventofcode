package main

import (
	"fmt"
	"strconv"
)

func main() {
	max := 707912
	min := 165432

	p1count := 0
	p2count := 0
	for num := min; num < max; num++ {
		pwd := getPassword(num)

		// Make sure it is increasing (or staying the same)
		increasing := pwd[0] <= pwd[1] && pwd[1] <= pwd[2] && pwd[2] <= pwd[3] && pwd[3] <= pwd[4] && pwd[4] <= pwd[5]
		if !increasing {
			continue
		}

		// Make sure it has a double
		if !hasDouble(pwd) {
			continue
		}

		// At this point all checks (for part 1) have passed
		p1count++

		// Part 2 double check
		if !hasDoublePart2(pwd) {
			continue
		}

		// At this point all checks (for part 2) have passed
		p2count++
	}

	fmt.Println("Part 1, possible passwords:")
	fmt.Println(p1count)
	fmt.Println("Part 2, possible passwords:")
	fmt.Println(p2count)
}

// Checks if the password has a digit twice in it
func hasDouble(pwd [6]int) bool {
	// Check by making a slice with no duplicates (a set)
	// If the length of the slice decreases to 5 then there is just one double
	set := []int{}
	seen := map[int]bool{}
	for _, num := range pwd {
		if !seen[num] {
			seen[num] = true
			set = append(set, num)
		}
	}

	return len(set) <= 5
}

// Checks if the password has a digit twice in it and is not a part of a larger group
func hasDoublePart2(pwd [6]int) bool {
	seen := map[int]int{}
	for _, num := range pwd {
		seen[num]++
	}

	for _, times := range seen {
		if times == 2 {
			return true
		}
	}

	return false
}

// Convert an int to an array of its digits
func getPassword(num int) [6]int {
	// Get string password
	pwdstr := strconv.Itoa(num)

	// Add each digit to the array
	var pwd [6]int
	for i := 0; i < 6; i++ {
		pwddigit, _ := strconv.Atoi(string(pwdstr[i]))
		pwd[i] = pwddigit
	}

	// Return
	return pwd
}
