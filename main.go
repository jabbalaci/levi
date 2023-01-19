package main

// Calculate the Levenshtein distance between two text files.

import (
	"fmt"
	"os"

	"github.com/jabbalaci/levi/lib/jfile"
	"github.com/jabbalaci/levi/lib/jtext"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: levi file1.txt file2.txt")
		os.Exit(1)
	}
	// else
	text1, err := jfile.Read(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file1:", err)
		os.Exit(1)
	}
	text2, err := jfile.Read(os.Args[2])
	if err != nil {
		fmt.Println("Error reading file2:", err)
		os.Exit(1)
	}
	result := jtext.LevDist(text1, text2)
	fmt.Println("Levenshtein distance:", result)
}
