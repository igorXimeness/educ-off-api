package model

type Modules struct {
	SubjectID  int    `json:"subject_Id"`
	Name       string `json:"name"`
	ModuleName string `json:"module_name"`
	Done       bool   `json:"done"`
}
