package main

import (
	"fmt"
	"github.com/pe5er/radiotutor/quiz"
	"strconv"
)

func main() {
	r := routes()
	r.LoadHTMLGlob("templates/*")

	quiz.QuestionsInit()

	score := 0
	qnum := 0

	for true {
		fmt.Println("Score " + strconv.Itoa(score))
		q := quiz.ReturnQuestion("M0", qnum)
		printQuestion(q)
		if checkAnswer(getAnswer(), q) {
			fmt.Println("Welldone")
			score += 1
		} else {
			fmt.Println("The answer is incorrect\n The correct answer is: " + returnAnswer(q))
		}
		qnum += 1
	}

	// Listen and Server in https://127.0.0.1:8080
	//r.Run(":8080")
}

func printQuestion(q quiz.Question) {
	fmt.Println(q.Question)
	fmt.Println("1) " + q.Answers[0].Answer)
	fmt.Println("2) " + q.Answers[1].Answer)
	fmt.Println("3) " + q.Answers[2].Answer)
	fmt.Println("4) " + q.Answers[3].Answer)
}

func getAnswer() int {
	var i int
	fmt.Scanf("%d", &i)
	return i
}

func checkAnswer(i int, q quiz.Question) bool {
	return q.Answers[i-1].Correct
}

func returnAnswer(q quiz.Question) string {
	for _, a := range q.Answers {
		if a.Correct {
			return a.Answer
		}
	}
	return ""
}
