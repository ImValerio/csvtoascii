package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic("You must specify the file name :( ")
	}

	fileName := os.Args[1]

	printIntro(fileName)

	dat, err := os.ReadFile(fmt.Sprint("./", fileName))
	if err != nil {
		panic(err)
	}

	fileContent := string(dat)
	lines := strings.Split(fileContent, "\n")

	for index, line := range lines {
		values := strings.Split(line, ",")
		formattedLine := strings.Join(values, " | ")
		separator := strings.Repeat("-", len(formattedLine))

		if index == 0 {
			separator = strings.Repeat("=", len(formattedLine))
			fmt.Println(separator)
		}

		fmt.Print(formattedLine)
		fmt.Println("")
		fmt.Println(separator)

	}

}

func printIntro(fileName string) {
	fmt.Println("=== CONVERTING ===")
	fmt.Println(fileName)
	fmt.Println("==================")
}

// returns a map containing the max length of words in each column
// colNum -> maxColLength
func generateMapMaxLength() map[string]int {
	rv := make(map[string]int)

	return rv
}
