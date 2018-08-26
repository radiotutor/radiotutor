package pages

import (
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
	if len(q) > current {
		c.HTML(200, "question.html", gin.H{
			"Current":  strconv.Itoa(current + 1),
			"NumQ":     strconv.Itoa(v.(int)),
			"Question": q[current].Question,
			"Answer1":  q[current].Answers[0].Answer,
			"Answer2":  q[current].Answers[1].Answer,
			"Answer3":  q[current].Answers[2].Answer,
			"Answer4":  q[current].Answers[3].Answer})
	}

	//session.Save()
}

func QuizPost(c *gin.Context) {
	session := sessions.Default(c)
	l := c.Param("licenceType")

	// Initial stage post
	if numberOfQuizS, ok := c.GetPostForm("number"); ok {
		numberOfQuiz, err := strconv.Atoi(numberOfQuizS)
		if err != nil {
			QuizGet(c)
			return
		}
		session.Set(l+"Started", numberOfQuiz)
		session.Save()
	}

	v := session.Get(l + "Current")

	// Check Proper Post
	if session.Get(l+"Started") == nil {
		QuizGet(c)
		return
	}
	if session.Get(l+"Started").(int) < 0 || session.Get(l+"Started").(int) > 40 {
		session.Delete(l + "Started")
		QuizGet(c)
		session.Save()
		return
	}
	if v == nil {
		// Setup
		session.Set(l+"Current", int(0))
		correct := make([]bool, session.Get(l+"Started").(int))
		session.Set(l+"Quiz", quiz.ReturnQuiz(l, 1, session.Get(l+"Started").(int)))
		session.Set(l+"QuizCorrect", correct)
	} else {

		// Change question
		q := session.Get(l + "Quiz").([]quiz.Question)
		len := session.Get(l + "Started").(int)
		current := v.(int)
		correctIndex := -1
		correct := session.Get(l + "QuizCorrect").([]bool)

		for i, ans := range q[current].Answers {
			if ans.Correct {
				correctIndex = i
				break
			}
		}
		ans, _ := c.GetPostForm("answer")
		if i, _ := strconv.Atoi(ans); i == correctIndex {
			correct[current] = true
		} else {
			correct[current] = false
		}

		session.Set(l+"QuizCorrect", correct)

		if (len - 1) == current {
			score := int(0)
			for _, j := range correct {
				if j {
					score++
				}
			}

			session.Delete(l + "Started")
			session.Delete(l + "Current")
			session.Delete(l + "Quiz")
			session.Delete(l + "QuizCorrect")
			session.Save()
			c.String(200, strconv.Itoa(score)+"/"+strconv.Itoa(len))
			return
		} else {
			current++
		}
		session.Set(l+"Current", current)
	}

	session.Save()
	QuizGet(c)

	//c.JSON(200, gin.H{"number": numberOfQuiz})
}
