package main

import (
	"fmt"
	"github.com/hardy8059/Quiz_game_using_golang/commom_utilities"
	"github.com/hardy8059/Quiz_game_using_golang/timed_quiz"
	"github.com/hardy8059/Quiz_game_using_golang/untimed_quiz"
)

func main() {
	var input int
	csvFilename, timerDuration := commom_utilities.InitArguments()
	//csvFilename := flag.String("csv", "questions.csv", "a csv file in the format of 'question,answer'")
	//timerDuration := flag.Int("time", 30, "time to finish the whole quiz.")
	//flag.Parse()
	fmt.Println("Hi! Welcome to Maths Quiz.\n" +
		"Choose the type of quiz you want to play:\n" +
		"1. Timed Quiz\n" +
		"2. UnTimed Quiz")
	fmt.Scanf("%d\n", &input)
	switch input {
	case 1:
		timed_quiz.TimedQuiz(csvFilename, timerDuration)
	case 2:
		untimed_quiz.UnTimedQuiz(csvFilename)
	default:
		fmt.Println("Wrong selection. Please restart and type only 1 or 2.")
	}
}
