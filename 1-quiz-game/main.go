package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func loadCsvRecords(filename string) [][]string {
	// try to open the file, and defer it's closure
	// https://golang.org/pkg/os
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// create a csv reader pointer (*Reader) and configure it
	// to expect two fields per csv line (record)
	// https://golang.org/pkg/encoding/csv
	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = 2

	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	return records
}

func playGame(records [][]string) chan int {
	c := make(chan int)

	// start the actual game as a goroutine
	// once it is done, it will send message to the chan,
	// signaling the game is over
	go func() {
		correctAnswers := 0
		answerReader := bufio.NewReader(os.Stdin)
		for questionNumber, record := range records {
			question := record[0]
			answer := record[1]

			var userAnswer string
			for {
				fmt.Printf("%s: %s ->  ", strconv.Itoa(questionNumber+1), question)

				rawAnswer, err := answerReader.ReadString('\n')
				if err != nil {
					fmt.Println("Error: try again...")
					continue
				}

				userAnswer = strings.TrimSpace(rawAnswer)

				break
			}

			if userAnswer == answer {
				correctAnswers++
			}
		}

		c <- correctAnswers
	}()

	return c
}

func main() {
	filename := "problems.csv"

	records := loadCsvRecords(filename)

	// channel that signals once the game has ended
	gameChan := playGame(records)

	select {
	case correctAnswers := <-gameChan:
		fmt.Printf("\nCorrect Answers: %s\n", strconv.Itoa(correctAnswers))
		fmt.Println("Game Over")
	case <-time.After(30 * time.Second):
		fmt.Println("\n\nTime out")
	}
}
