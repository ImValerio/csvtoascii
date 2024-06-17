package main;
import (
	"fmt"
	"os"
	"strings"
)

func main(){
	filename := os.Args[0]
	if filename == "" {
		filename = "test.csv"
	}

	fmt.Println("=== CONVERTING ===")
   	fmt.Println(filename);	
	fmt.Println("==================")
	dat, err := os.ReadFile("./test.csv")
	if err != nil{
		panic(err)
	}
    	fileContent := string(dat)
	lines := strings.Split(fileContent, "\n")

	for index, line := range lines {
		values := strings.Split(line, ",")
		formattedLine := strings.Join(values, " | ")
		separator := strings.Repeat("-",len(formattedLine))

		if index == 0 {
			separator = strings.Repeat("=",len(formattedLine))
			fmt.Println(separator)
		}

		fmt.Print(formattedLine)
		fmt.Println("")
		fmt.Println(separator)
		
	}


}


