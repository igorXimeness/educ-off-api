package model

type Lesson struct {
	LessonID int    `json:"lesson_id"`
	ModuleID int    `json:"modules_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}



type Question struct {
    QuestionID int    `json:"question_id"`
    LessonID   int    `json:"lesson_id"`      // Relaciona a pergunta à lição
    QuestionText string `json:"question_text"`  // Texto da pergunta
    OptionA    string `json:"option_a"`        // Opção A
    OptionB    string `json:"option_b"`        // Opção B
    OptionC    string `json:"option_c"`        // Opção C
    OptionD    string `json:"option_d"`        // Opção D
    RightOption string `json:"right_option"`    // A opção correta (A, B, C ou D)
}
