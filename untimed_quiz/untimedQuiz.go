package untimed_quiz

import (
	"fmt"
	"github.com/hardy8059/Quiz_game_using_golang/commom_utilities"
	"os"
)

func startUnTimedQuiz(problemArray []*commom_utilities.Problem) {
	var answer string
	correct := 0
	for i, p := range problemArray {
		fmt.Printf("Problem %d: %s is ?\n", i+1, p.Question)
		fmt.Scanf("%s\n", &answer)
		if answer == p.Answer {
			fmt.Println("Correct Answer!")
			correct++
		}
	}
	fmt.Printf("You correctly answered %d out of %d questons!\n", correct, len(problemArray))
}

func UnTimedQuiz(csvFilename *string) {
	file, err := os.Open(*csvFilename)
	if err != nil {
		commom_utilities.Exit(fmt.Sprintf("Unable to open file %s", *csvFilename))
	} else {
		fmt.Println("Questions Loaded!")
	}
	lines := commom_utilities.ReadCsvData(file)
	problemArray := commom_utilities.ParseCsvData(lines)
	startUnTimedQuiz(problemArray)
}
