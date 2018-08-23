package pages

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pe5er/radiotutor/quiz"
	"strconv"
)

func QuizGet(c *gin.Context) {
	session := sessions.Default(c)
	l := c.Param("licenceType")

	v := session.Get(l + "Started")
	if v == nil {
		c.HTML(200, "quiz.html", gin.H{"Licence": l})
		return
	}

	q := session.Get(l + "Quiz").([]quiz.Question)
	current := session.Get(l + "Current").(int)
	c.HTML(200, "question.html", gin.H{
		"Current":  strconv.Itoa(current + 1),
		"NumQ":     strconv.Itoa(v.(int)),
		"Question": q[current].Question,
		"Answer1":  q[current].Answers[0].Answer,
		"Answer2":  q[current].Answers[1].Answer,
		"Answer3":  q[current].Answers[2].Answer,
		"Answer4":  q[current].Answers[3].Answer})

	//session.Save()
}

func QuizPost(c *gin.Context) {
	quiz.QuestionsInit()
	session := sessions.Default(c)
	l := c.Param("licenceType")

	if numberOfQuizS, ok := c.GetPostForm("number"); ok {
		numberOfQuiz, err := strconv.Atoi(numberOfQuizS)
		if err != nil {
			c.Redirect(302, c.Params[0].Value)
			return
		}
		session.Set(l+"Started", numberOfQuiz)
		session.Save()
	}

	v := session.Get(l + "Current")
	if v == nil {
		session.Set(l+"Current", int(0))
		correct := make([]bool, session.Get(l+"Started").(int))
		session.Set(l+"Quiz", quiz.ReturnQuiz(l, 1, session.Get(l+"Started").(int)))
		fmt.Println(session.Get(l + "Quiz"))
		fmt.Println(quiz.ReturnQuestion("M0", 1))
		fmt.Println(session.Get(l + "Started").(int))
		fmt.Println(quiz.ReturnQuiz(l, 1, session.Get(l+"Started").(int)))
		session.Set(l+"QuizCorrect", correct)
	} else {
		current := v.(int)
		current++
		session.Set(l+"Current", current)
	}

	session.Save()
	c.Redirect(302, "exam")

	//c.JSON(200, gin.H{"number": numberOfQuiz})
}
