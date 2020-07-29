package timed_quiz

import (
	"fmt"
	"github.com/hardy8059/Quiz_game_using_golang/commom_utilities"
	"os"
	"time"
)

func takeUserInput(answerChannel chan string) {
	var answer string
	fmt.Scanf("%s\n", &answer)
	answerChannel <- answer
}

func startTimedQuiz(problemArray []*commom_utilities.Problem, d *int) {
	timer := time.NewTimer(time.Duration(*d) * time.Second)
	correct := 0
	answerChannel := make(chan string)
	for i, p := range problemArray {
		fmt.Printf("Problem %d: %s is ?\n", i+1, p.Question)
		go takeUserInput(answerChannel)
		select {
		case <-timer.C:
			fmt.Printf("Time Up! You correctly answered %d out of %d questons!\n", correct, len(problemArray))
			return
		case answer := <-answerChannel:
			if answer == p.Answer {
				correct++
			}
		}
	}
	fmt.Printf("You correctly answered %d out of %d questons!\n", correct, len(problemArray))
}

func TimedQuiz(csvFilename *string, timerDuration *int) {
	file, err := os.Open(*csvFilename)
	if err != nil {
		commom_utilities.Exit(fmt.Sprintf("Unable to open file %s", *csvFilename))
	} else {
		fmt.Println("Questions Loaded!")
	}
	lines := commom_utilities.ReadCsvData(file)
	problemArray := commom_utilities.ParseCsvData(lines)
	startTimedQuiz(problemArray, timerDuration)
}
