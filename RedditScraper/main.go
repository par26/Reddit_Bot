package main

import (
	"log"
	"os"
)

func main() {
	questions, err := GetPost("Top", "Showerthoughts", 20)

	if err != nil {
		log.Fatal(err)
	}

	writeToFile(questions)
}

func writeToFile(questions []string) {
	f, err := os.Create("question.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	for _, question := range questions {
		_, err := f.WriteString(question + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

/*unc getUserInput() {

	var input string

	fmt.Scanln(input)

	switch input {
	case "1":

	default:
		fmt.Println("Invalid Input")
	}

	fmt.Println("Uh oh")
}*/
