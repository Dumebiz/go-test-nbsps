// Title: Script to print system environment variables

// Author: Chu Dumebi Umezinne

package main

// Import required packages
import (
	"fmt"
	"os"
)

// This function prints the environment variables line by line
func main() {
	for _, data := range os.Environ() {
		fmt.Println(data)
	}
}
