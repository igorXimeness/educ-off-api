package model

type Lesson struct {
	LessonID int    `json:"lesson_id"`
	ModuleID int    `json:"modules_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
