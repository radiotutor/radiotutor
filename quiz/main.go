package quiz

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
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
	case "AV":
		return advanced[n%len(advanced)]
	case "I":
		return intermediate[n%len(intermediate)]
	case "F":
		return foundation[n%len(foundation)]
	}
	return Question{}
}

func ReturnQuiz(t string, n int, num int) []Question {
	rtn := make([]Question, num)
	copy := make([]int64, num)
	for i, _ := range copy {
		copy[i] = -1
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	sanityCounter := 0
OUTER:
	for i := 0; i < num; i++ {
		z := r.Intn(1000)
		q := ReturnQuestion(t, z)
		for _, c := range copy {
			if q.Number == c {
				i--
				// For the sake of this never becoming an infinate loop
				if sanityCounter++; sanityCounter > 10000 {
					return nil
				}
				continue OUTER
			}
		}
		rtn[i] = q
		copy[i] = q.Number
	}
	return rtn

}

func Mark(questions []Question, answers []byte) int {
	score := 0
	for i, q := range questions {
		for z, ans := range q.Answers {
			if ans.Correct {
				if int(answers[i]) == z {
					score++

				}
				break
			}
		}
	}
	return score
}
