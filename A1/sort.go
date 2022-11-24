//Name: Dana Curca
//Student Number: 250976773
//CS 4435 Assignment0
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	var numArray []float64 //initialize float64 array

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
		numArray = append(numArray, numbers) //append all numbers in txt file to numArray
	}
	file.Close()
	sort.Float64s(numArray) //sort the numbers in numArray

	outputFile, err := os.Create("sortedNumbers.txt") //create the empty sortedNumbers.txt file
	if err != nil {
		log.Fatalf("Output file could not be created, please try again")
	}

	for _, numbers := range numArray { //parse through numArray
		_, err := outputFile.WriteString(fmt.Sprintf("%.1f\n", numbers)) //for every number write it on the output file sortedNumbers.txt as a float

		if err != nil {
			log.Fatalf("Write to output file failed, please try again")
		}
	}
	outputFile.Close() //close the new txt file
}
