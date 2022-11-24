//Name: Dana Curca
//Student Number: 250976773
//CS 4435 Assignment0
package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

//struct for json file
type bankAccount struct {
	Name      string  `json:"Name"`
	AccountID int     `json:"AccountID"`
	Balance   float64 `json:"Balance"`
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 { //error checking to make sure arguments is not below or above 1
		log.Fatalf("Invalid number of arguments, please try again with only 1 argument")
	}

	inputFile := os.Args[1]
	jsonFile, err := os.Open(inputFile) //open json file
	if err != nil {
		log.Fatalf("Error opening json file, please try again")
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile) //read json as a bytearray
	if err != nil {
		log.Fatalf("Error reading json file, please try again")
	}

	var account []bankAccount           //initialize bankAccount array that stores the accounts
	json.Unmarshal(byteValue, &account) //unmarshal bytearray into accounts
	for i := range account {
		account[i].Balance += 100
	}

	byteValue, _ = json.Marshal(account) //marshal bytearray back into json
	outputFile, err := os.Create("accountsUpdated.json")
	if err != nil {
		log.Fatalf("Output file could not be created, please try again")
	}

	outputFile.Write(byteValue) //write the values onto the output file
	outputFile.Close()
}
