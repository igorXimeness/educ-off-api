package model

type Lesson struct {
    LessonID    int    `json:"lesson_id"`     // Gerado automaticamente
    Title       string `json:"title"`         // Título da lição
    Content     string `json:"content"`       // Conteúdo da lição
    SubjectName string `json:"subject_name"`  // Nome da matéria
    ModuleName  string `json:"module_name"`   // Nome do módulo
}
