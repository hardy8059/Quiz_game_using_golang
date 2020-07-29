package commom_utilities

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Problem struct {
	Question string
	Answer   string
}

func InitArguments() (*string, *int) {
	csvFilename := flag.String("csv", "questions.csv", "a csv file in the format of 'question,answer'")
	timerDuration := flag.Int("time", 30, "time to finish the whole quiz.")
	flag.Parse()
	return csvFilename, timerDuration
}

func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func ReadCsvData(file *os.File) [][]string {
	result := csv.NewReader(file)
	lines, err := result.ReadAll()
	if err != nil {
		Exit(fmt.Sprintf("Unable to read the questions from csv file."))
	}
	return lines
}

func ParseCsvData(lines [][]string) []*Problem {
	problemArray := make([]*Problem, len(lines))
	for i, line := range lines {
		problemArray[i] = &Problem{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]),
		}
	}
	return problemArray
}
