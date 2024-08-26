package model

type Lessons struct {
	LessonID int    `json:"lesson_id"`
	ModuleID int    `json:"module_id"`
	Title    string `json:"title"`
}
