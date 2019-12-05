package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Coords consist of x and y
type Coords struct {
	x int
	y int
}

func main() {
	// Part 1
	part1()

	// Part 2
	part2()
}

// Part 1
func part1() {
	// The 2 paths
	path1, path2 := getInput()

	// Store each step of the path in the traces
	trace1, trace2 := []Coords{Coords{0, 0}}, []Coords{Coords{0, 0}}
	// Store x and y
	var x, y int

	// Trace path1
	x, y = 0, 0
	for _, instruction := range path1 {
		direction := string(instruction[0])
		amount, _ := strconv.Atoi(instruction[1:])

		// Move
		for i := 0; i < amount; i++ {
			if direction == "U" {
				y++
			} else if direction == "D" {
				y--
			} else if direction == "R" {
				x++
			} else if direction == "L" {
				x--
			}
			trace1 = append(trace1, Coords{x, y})
		}
	}

	// Trace path2
	x, y = 0, 0
	for _, instruction := range path2 {
		direction := string(instruction[0])
		amount, _ := strconv.Atoi(instruction[1:])

		// Move
		for i := 0; i < amount; i++ {
			if direction == "U" {
				y++
			} else if direction == "D" {
				y--
			} else if direction == "R" {
				x++
			} else if direction == "L" {
				x--
			}
			trace2 = append(trace2, Coords{x, y})
		}
	}

	// Get the shortest manhattan distance
	distance := 0
	for i := 0; i < len(trace1); i++ {
		for j := 0; j < len(trace2); j++ {
			if trace1[i] == trace2[j] {
				result := abs(trace1[i].x) + abs(trace1[i].y)
				if result < distance || distance == 0 {
					distance = result
				}
			}
		}
	}

	fmt.Println("Shortest manhattan distance:")
	fmt.Println(distance)
}

// Part 2
func part2() {
	// The 2 paths
	path1, path2 := getInput()

	// Store each step of the path in the traces
	trace1, trace2 := []Coords{Coords{0, 0}}, []Coords{Coords{0, 0}}
	// Store x and y
	var x, y int

	// Trace path1
	x, y = 0, 0
	for _, instruction := range path1 {
		direction := string(instruction[0])
		amount, _ := strconv.Atoi(instruction[1:])

		// Move
		for i := 0; i < amount; i++ {
			if direction == "U" {
				y++
			} else if direction == "D" {
				y--
			} else if direction == "R" {
				x++
			} else if direction == "L" {
				x--
			}
			trace1 = append(trace1, Coords{x, y})
		}
	}

	// Trace path2
	x, y = 0, 0
	for _, instruction := range path2 {
		direction := string(instruction[0])
		amount, _ := strconv.Atoi(instruction[1:])

		// Move
		for i := 0; i < amount; i++ {
			if direction == "U" {
				y++
			} else if direction == "D" {
				y--
			} else if direction == "R" {
				x++
			} else if direction == "L" {
				x--
			}
			trace2 = append(trace2, Coords{x, y})
		}
	}

	// Get the smallest number of steps
	steps := 0
	for i := 0; i < len(trace1); i++ {
		for j := 0; j < len(trace2); j++ {
			if trace1[i] == trace2[j] {
				result := i + j
				if result < steps || steps == 0 {
					steps = result
				}
			}
		}
	}

	fmt.Println("Shortest manhattan distance:")
	fmt.Println(steps)
}

// Returns the absolute value of an int
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Get the input from input.txt
func getInput() ([]string, []string) {
	// Read the input
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while reading the input")
	}
	defer file.Close()

	// Create a scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Read the 2 lines
	scanner.Scan()
	line1 := scanner.Text()
	scanner.Scan()
	line2 := scanner.Text()

	// There are 2 paths separated by a new line
	path1 := strings.Split(line1, ",")
	path2 := strings.Split(line2, ",")

	return path1, path2
}
