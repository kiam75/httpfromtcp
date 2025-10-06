package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	inputfile, err := os.Open("message.txt")
	if err != nil {
		log.Printf("Error open file. Errormessage: %v", err)
	}
	defer inputfile.Close()

	buffer := make([]byte, 8)

	for {
		n, err := inputfile.Read(buffer)

		// Print the actuall bytes read (only n bytes, not the full buffer)
		if n > 0 {
			fmt.Printf("read: %s\n", buffer[:n])
		}

		// Handle errors with switch
		switch err {
		case nil:
			// Continue reading
			continue
		case io.EOF:
			// End of file reached, exit cleanly
			//break
			os.Exit(0)
		default:
			// Handle other errors
			log.Printf("Error reading file: %v", err)
			//break
			os.Exit(1)
		}
	}

}
