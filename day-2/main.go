package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// Opcodes:
	// 1  - Add the 2 numbers at the positions specified by the next 2 ints
	//      and store them at the position specificed by the third int
	// 2  - Same as above but multiplication
	// 99 - Halt the execution

	input := getInput()

	// Replace position 1 and 2
	input[1] = 79
	input[2] = 60

	// Run the program
	index := 0
	for index < len(input) {
		// 1 opcode
		if input[index] == 1 {
			// Get the 3 positions we need
			pos1 := input[index+1]
			pos2 := input[index+2]
			pos3 := input[index+3]

			val1 := input[pos1]
			val2 := input[pos2]

			// Set the new value to the third position
			input[pos3] = val1 + val2

			// Go four steps ahead
			index += 4
			continue
		}

		// 2 opcode
		if input[index] == 2 {
			// Get the 3 positions we need
			pos1 := input[index+1]
			pos2 := input[index+2]
			pos3 := input[index+3]

			val1 := input[pos1]
			val2 := input[pos2]

			// Set the new value to the third position
			input[pos3] = val1 * val2

			// Go four steps ahead
			index += 4
			continue
		}

		// 99 opcode
		if input[index] == 99 {
			break
		}

		// Something else
		break
	}

	// Print result
	fmt.Println("Result:")
	fmt.Println(input[0])
}

// Gets the input from input.txt
func getInput() []int {
	// Read the input file
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Could not read the input file!")
	}
	data := string(file)

	// Store the input values here
	input := []int{}

	// Fill in the input slice
	splitted := strings.Split(data, ",")
	for _, number := range splitted {
		convNumber, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Error while converting the input data.")
		}
		input = append(input, convNumber)
	}

	return input
}
