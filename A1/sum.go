//Name: Dana Curca
//Student Number: 250976773
//CS 4435 Assignment0
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var result float64

	args := os.Args[1:]
	if len(args) != 1 { //error checking to make sure arguments is not below or above 1
		log.Fatalf("Invalid number of arguments, please try again with only 1 argument")
	}

	filename := os.Args[1]
	file, err := os.Open(filename) //open txt file
	if err != nil {
		log.Fatalf("Failed to open file or no file exists")
	}

	scanner := bufio.NewScanner(file) //scan the txt file
	for scanner.Scan() {
		numbers, err := strconv.ParseFloat(scanner.Text(), 64) //convert all numbers in txt file to float64
		if err != nil {
			log.Fatalf("Failed to convert numbers to a float")
		}
		result += numbers
	}
	file.Close() //close the txt file
	fmt.Println("The sum of the numbers in the file is: ", result)
}
