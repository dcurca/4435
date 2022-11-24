//Name: Dana Curca
//Student Number: 250976773
//CS 4435 Assignment0

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var result float64

	args := os.Args[1:]
	if len(args) != 3 { //error checking to make sure arguments is not below or above 3
		log.Fatal("Invalid number of arguments, please try again with 3 arguments")
	}

	op := args[0] //store command line argument 0 in op
	i := args[1]  //store command line argument 1 in i
	j := args[2]  //store command line argument 2 in j

	x, err1 := strconv.ParseFloat(i, 64) //convert string i into variable x float64
	y, err2 := strconv.ParseFloat(j, 64) //convert string j into variable y float64
	if err1 != nil && err2 != nil {
		log.Fatal("Could not convert to float")
	}

	if op == "add" {
		result = x + y
		fmt.Println("The sum of these numbers is:", result)
	} else if op == "sub" {
		result = x - y
		fmt.Println("The difference of these numbers is:", result)
	} else if op == "div" {
		if y == 0 {
			log.Fatal("Divider cannot be 0")
		} else {
			result = x / y
			fmt.Println("The quotient of these numbers is:", result)
		}
	} else if op == "mult" {
		result = x * y
		fmt.Println("The product of these numbers is:", result)
	} else {
		fmt.Println("Your operator is incorrect, please try again with a valid operator")
	}
}
