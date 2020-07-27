package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func initArguments() *string {
	csvFilename := flag.String("csv", "questions.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	return csvFilename
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func readCsvData(file *os.File) [][]string {
	result := csv.NewReader(file)
	lines, err := result.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Unable to read the questions from csv file."))
	}
	return lines
}

func parseCsvData(lines [][]string) []*problem {
	problemArray := make([]*problem, len(lines))
	for i, line := range lines {
		problemArray[i] = &problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return problemArray
}

func startQuiz(problemArray []*problem) {
	var answer string
	correct := 0
	for i, p := range problemArray {
		fmt.Printf("Problem %d: %s is ?\n", i+1, p.question)
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			fmt.Println("Correct Answer!")
			correct++
		}
	}
	fmt.Printf("You correctly answered %d out of %d questons!\n", correct, len(problemArray))
}

func main() {
	csvFilename := initArguments()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Unable to open file %s", *csvFilename))
	} else {
		fmt.Println("Questions Loaded!")
	}
	lines := readCsvData(file)
	problemArray := parseCsvData(lines)
	startQuiz(problemArray)
}
