package data

type ModuleInfo struct {
	ID             int    `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	ModuleName     string `json:"module_name"`
	ModuleDuration int    `json:"module_duration"`
	ExamType       string `json:"exam_type"`
	Version        int    `json:"version"`
}
