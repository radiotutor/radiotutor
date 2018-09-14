package main

import "github.com/pe5er/radiotutor/quiz"

func tmplInc(i int) int {
	return i + 1
}

func tmplCheck(questions []quiz.Question, answers []byte, index int) bool {
	for z, ans := range questions[index].Answers {
		if ans.Correct {
			if int(answers[index]) == z {
				return true
			}
			break
		}
	}

	return false
}
