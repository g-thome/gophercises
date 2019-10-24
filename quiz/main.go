package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func ask(question string, answer int, ch chan int) {
	var input int
	fmt.Print(question)
	fmt.Scan(&input)
	ch <- input
}

func main() {
	file := flag.String("file", "problems.csv", "csv file with questions and answers in the format \"question,answer\"")
	limit := flag.Int("limit", 99999, "time limit in seconds for each question")
	flag.Parse()

	var score int
	ch := make(chan int)
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
		go ask(currentQuestion, currentAnswer, ch)

		select {
		case answer := <-ch:
			if answer == currentAnswer {
				score++
			}
		case <-time.After(time.Duration(*limit) * time.Second):
			fmt.Println("timeout")
		}
	}

	fmt.Printf("your score is %d \n", score)
}
