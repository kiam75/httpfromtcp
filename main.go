package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	inputfile, err := os.Open("message.txt")
	if err != nil {
		log.Printf("Error open file. Errormessage: %v", err)
	}
	defer inputfile.Close()

	buffer := make([]byte, 8)
	// line accumulator
	lineBuffer := ""

	for {
		n, err := inputfile.Read(buffer)

		// Print the actuall bytes read (only n bytes, not the full buffer)
		if n > 0 {
			// Convert bytes to string and split on newline
			log.Printf("read into data: %s\n", buffer[:n])
			data := string(buffer[:n])
			parts := strings.Split(data, "\n")

			// process all parts except the last one
			for i := 0; i < len(parts)-1; i++ {
				// Complete line found: currentLine + this part
				completeLine := lineBuffer + parts[i]
				fmt.Printf("read: %s\n", completeLine)
				lineBuffer = "" // Reset for next part
			}

			// Add the last part to currentLine (might be empty)
			lineBuffer += parts[len(parts)-1]
			log.Printf("Content of lineBuffer: %s\n", lineBuffer)
		}

		// Handle errors with switch
		switch err {
		case nil:
			// Continue reading
			continue
		case io.EOF:
			// Print any remaining line before exiting
			if lineBuffer != "" {
				fmt.Printf("read: %s\n", lineBuffer)
			}
			// End of file reached, exit cleanly
			os.Exit(0)
		default:
			// Handle other errors
			log.Printf("Error reading file: %v", err)
			//break
			os.Exit(1)
		}
	}

}
