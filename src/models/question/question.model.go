package question

// import (
// 	"github.com/google/uuid"
// )

type Question struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	GroupId string `json:"groupId"`
}

var QuestionData = []Question{
	{ID: "1", Title: "Selecciona la alternativa verdadera", GroupId: "1"},
}
