package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		panic("You must specify the file name :( ")
	}

	fileName := os.Args[1]
	var fileNames []string

	if strings.Contains(fileName, ";") {
		fileNames = strings.Split(fileName, ";")
	}

	if len(fileNames) > 1 {
		isCsvFile(fileNames)
		printIntro(fileNames)
		multiFileProcess(fileNames)
	} else {
		isCsvFile([]string{fileName})
		printIntro([]string{fileName})
		processCsv(fileName)
	}

	fmt.Println("Total process time: ", time.Since(start))
}

func isCsvFile(fileNames []string) bool {

	for _, fileName := range fileNames {
		ext := fileName[strings.LastIndex(fileName, "."):]

		if ext != ".csv" {
			return false
		}
	}

	return true
}

func multiFileProcess(fileNames []string) {
	var wg sync.WaitGroup

	for _, fileName := range fileNames {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processCsv(fileName)
		}()
	}
	wg.Wait()
}

func processCsv(fileName string) {
	start := time.Now()

	dat, err := os.ReadFile(fmt.Sprint("./", fileName))
	if err != nil {
		panic(err)
	}

	fileContent := string(dat)
	lines := strings.Split(fileContent, "\n")

	maxColLength := generateSliceMaxLength(lines)
	// fmt.Println(maxColLength)
	outFileName := generateOutFileName(fileName)

	f, err := os.Create(outFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for index, line := range lines {
		values := splitRespectingQuotes(line, ',')
		for i, value := range values {
			values[i] = value + strings.Repeat(" ", maxColLength[i]-len(value))
		}
		formattedLine := strings.Join(values, " | ")
		// fmt.Println(formattedLine)
		separator := strings.Repeat("-", len(formattedLine))

		if index == 0 {
			separator = strings.Repeat("=", len(formattedLine))
		}

		// fmt.Print(formattedLine)
		if hasTrailingNewline(formattedLine) {
			f.WriteString(formattedLine)
		} else {
			f.WriteString(formattedLine + "\n")
		}

		// fmt.Println(separator)
		f.WriteString(separator + "\n")

	}

	fmt.Println("Processed "+strconv.Itoa(len(lines)-1)+" lines in ", time.Since(start))

}
func hasTrailingNewline(str string) bool {
	return len(str) > 0 && str[len(str)-1] == '\n'
}

func printIntro(fileNames []string) {
	fmt.Println("=== CONVERTING ===")
	for _, fileName := range fileNames {
		fmt.Println(fileName)
	}
	fmt.Println("==================")
}

// returns a slice containing the max length of words in each column
// colNum -> maxColLength
func generateSliceMaxLength(lines []string) []int {
	var rv []int

	for _, line := range lines {

		values := splitRespectingQuotes(line, ',')
		for index, value := range values {
			wordLen := len(value)
			if len(rv) < len(values) {
				// fmt.Println("          ", len(rv), len(values))
				rv = append(rv, wordLen)
			}
			if rv[index] < wordLen {
				rv[index] = wordLen
			}
		}
	}
	return rv
}

func generateOutFileName(fileName string) string {
	nameSep := strings.LastIndex(fileName, ".")

	rawFileName := fileName[:nameSep]

	return rawFileName + "_table.txt"
}

func splitRespectingQuotes(str string, delimiter rune) []string {
	var result []string
	current := ""
	quoted := false

	for _, char := range str {
		if char == delimiter && !quoted {
			result = append(result, current)
			current = ""
		} else if char == '"' {
			quoted = !quoted
		} else {
			current += string(char)
		}
	}

	if len(current) > 0 {
		result = append(result, current)
	}

	return result
}
