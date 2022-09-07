package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func getCsvData(problemsFile string) [][]string {
	// Read the csv file
	file, err := os.Open(problemsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Extract each line value
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func getUserInput(question string) int {
	// Get User input
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What is %s ?\n", question)
	guess_input, _ := reader.ReadString('\n')
	guess_input = strings.TrimSpace(guess_input)
	guess, _ := strconv.Atoi(guess_input)
	return guess
}

func calculateScore(results []bool) {
	var goodAnswers int8 = 0
	questionNumber := len(results)

	for _, result := range results {
		if result {
			goodAnswers++
		}
	}
	println("\n")
	fmt.Printf("Finished: You have scored %d good answers over %d questions\n", goodAnswers,
		questionNumber)
}

func main() {
	var results []bool

	// Parse file flag
	problemsFile := flag.String("file", "problems.csv",
		"string: Path to CSV file containing questions")
	flag.Parse()
	data := getCsvData(*problemsFile)

	//Start Game
	fmt.Println("Quiz Game Starting...")
	fmt.Println("Press Enter when ready to start...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	// Start Timer
	timer30sec := time.NewTimer(5 * time.Second)

	for _, val := range data {
		answerData := val[1]
		answerData = strings.TrimSpace(answerData)
		answer, _ := strconv.Atoi(answerData)

		guess := getUserInput(val[0])

		// Evaluate answer and store result
		if guess == answer {
			results = append(results, true)
		} else {
			results = append(results, false)
		}
		<-timer30sec.C
		fmt.Println("Timer fired")
	}
	calculateScore(results)
}
