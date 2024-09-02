package model

type Modules struct {
	SubjectID  int    `json:"subject_Id"`
	ModulesID  int `json:"modules_Id"`
	ModuleName string `json:"module_name"`
	Done       bool   `json:"done"`
}
