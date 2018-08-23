package quiz

import (
	"encoding/json"
	"io/ioutil"
)

type Question struct {
	Answers []struct {
		Answer  string `json:"answer"`
		Correct bool   `json:"correct"`
	} `json:"answers"`
	Number   int64  `json:"number"`
	Question string `json:"question"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var (
	foundation   []Question
	intermediate []Question
	advanced     []Question
)

func QuestionsInit() {
	dat, err := ioutil.ReadFile("resources/foundation.json")
	check(err)
	err = json.Unmarshal(dat, &foundation)
	check(err)

	dat, err = ioutil.ReadFile("resources/intermediate.json")
	check(err)
	err = json.Unmarshal(dat, &intermediate)
	check(err)

	dat, err = ioutil.ReadFile("resources/advanced.json")
	check(err)
	err = json.Unmarshal(dat, &advanced)
	check(err)
}

func ReturnQuestion(t string, n int) Question {
	switch t {
	case "M0":
		return advanced[n%len(advanced)]
	case "2E0":
		return intermediate[n%len(intermediate)]
	case "M6":
		return foundation[n%len(foundation)]
	}
	return Question{}
}

func ReturnQuiz(t string, n int, num int) []Question {
	rtn := make([]Question, num)
	for i := 0; i < num; i++ {
		rtn[i] = ReturnQuestion(t, n+i)
	}
	return rtn

}
