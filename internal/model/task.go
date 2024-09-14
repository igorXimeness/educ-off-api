package model




type Task struct {
    TaskID      int    `json:"task_id"`
    Description string `json:"description"`
    Date        string `json:"date"`  // Manter como string no modelo
    Status      string `json:"status"`
}