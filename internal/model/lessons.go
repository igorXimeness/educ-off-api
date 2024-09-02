package model

type Lessons struct {
	LessonID int    `json:"lesson_id"`
	ModulesID int    `json:"modules_id"`
	Title    string `json:"title"`
}
