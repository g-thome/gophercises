package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ask(question string, answer int) bool {
	var input int
	fmt.Print(question)
	fmt.Scan(&input)

	if input == answer {
		return true
	} else {
		return false
	}
}

func main() {
	file := flag.String("file", "problems.csv", "csv file with questions and answers in the format \"question,answer\"")
	flag.Parse()

	var score int
	csvFile, _ := os.Open(*file)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	defer csvFile.Close()

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		}

		currentQuestion := fmt.Sprintf("Problem #%d: %s = ", 1, line[0])
		currentAnswer, _ := strconv.Atoi(line[1])
		if ask(currentQuestion, currentAnswer) {
			score++
		}
	}

	fmt.Printf("your score is %d \n", score)
}
