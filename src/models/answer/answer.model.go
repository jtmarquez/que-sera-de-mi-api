package answer

// import (
// 	"github.com/google/uuid"
// )

type Answer struct {
	ID            string `json:"id"`
	QuestionId    string `json:"questionId"`
	AlternativeId string `json:"alternativeId"`
}

var AnswerData = []Answer{
	{ID: "1", QuestionId: "1", AlternativeId: "1"},
}
