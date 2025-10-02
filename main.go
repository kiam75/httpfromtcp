package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	inputfile, err := os.Open("message.txt")
	if err != nil{
		log.Printf("Error open file. Errormessage: %v",err)
	}
	defer inputfile.Close()

	
}
