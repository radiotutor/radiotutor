package pages

import (
	"github.com/abaft/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pe5er/radiotutor/quiz"
	"strconv"
)

var (
	fullQuizSize = map[string]int{
		"F":  19,
		"I":  45,
		"AV": 40,
	}
)

func QuizGet(c *gin.Context) {
	session := sessions.Default(c)
	l := c.Param("licenceType")

	v := session.Get(l + "Started")
	if v == nil {
		c.HTML(200, "quiz.html", gin.H{"Licence": licenceCodeToName[l], "NoOfQuestions": fullQuizSize[l]})
		return
	}

	type state struct {
		Quiz    []quiz.Question
		Current int
		Length  int
		Sub     []byte
	}
	q := session.Get(l + "Quiz").([]quiz.Question)
	current := session.Get(l + "Current").(int)
	answers := session.Get(l + "QuizAnswers").([]byte)
	currentState := state{Quiz: q, Current: current, Length: v.(int), Sub: answers}

	if len(q) > current {
		c.HTML(200, "question.html", gin.H{"State": currentState})
	}

	//session.Save()
}

func QuizPost(c *gin.Context) {

	// Get Post Type
	postType, _ := c.GetPostForm("type")

	// Post Switch
	switch postType {
	case "start_quiz":
		startQuiz(c)
	case "question":
		questionHandle(c)
	default:
		QuizGet(c)
	}
}

func startQuiz(c *gin.Context) {
	session := sessions.Default(c)
	l := c.Param("licenceType")
	defer QuizGet(c)

	session.Clear()

	// Check Proper Quiz Start Post
	if sel, ok := c.GetPostForm("sel"); ok {
		sel, err := strconv.Atoi(sel)
		if err != nil || (sel != 1 && sel != 2) {
			return
		}

		// Set Quiz Size
		switch sel {
		case 1:
			session.Set(l+"Started", 10)
		case 2:
			session.Set(l+"Started", fullQuizSize[l])
		}
	}

	// Setup Quiz
	session.Set(l+"Current", int(0))
	quiz := quiz.ReturnQuiz(l, 1, session.Get(l+"Started").(int))
	if quiz == nil {
		// Session isn't saved so should be fine
		return
	}
	session.Set(l+"Quiz", quiz)
	answers := make([]byte, session.Get(l+"Started").(int))
	for i, _ := range answers {
		answers[i] = 5
	}
	session.Set(l+"QuizAnswers", answers)
	session.Save()
}

func questionHandle(c *gin.Context) {
	session := sessions.Default(c)
	l := c.Param("licenceType")

	// Check Valid Post and Session
	nav, _ := c.GetPostForm("nav")
	if session.Get(l+"Started") == nil ||
		(nav != "next" && nav != "previous" && nav != "exit") {
		QuizGet(c)
		return
	}

	// Check Exit
	if nav == "exit" {
		session.Clear()
		session.Save()
		QuizGet(c)
		return
	}

	current := session.Get(l + "Current").(int)
	answers := session.Get(l + "QuizAnswers").([]byte)
	q := session.Get(l + "Quiz").([]quiz.Question)
	numberOfQuestions := session.Get(l + "Started").(int)
	ans, _ := c.GetPostForm("answer")
	// Record Answer
	if i, _ := strconv.Atoi(ans); ans != "" && i >= 0 && i < 5 {
		answers[current] = byte(i)
	}
	// Increment Next
	if nav == "next" {
		current++
	} else if nav == "previous" && current != 0 {
		current--
	}

	// Check if quiz is over
	if current == numberOfQuestions {
		c.HTML(200, "result.html", gin.H{
			"Licence":       licenceCodeToName[l],
			"NoOfQuestions": strconv.Itoa(numberOfQuestions),
			"Score":         strconv.Itoa(quiz.Mark(q, answers)),
			"Questions":     q,
			"Sub":           answers,
		})
		session.Clear()
		session.Save()
		return
	}
	session.Set(l+"Current", current)
	session.Set(l+"QuizAnswers", answers)
	session.Save()
	QuizGet(c)
}
