package models

type Submission struct {
    ID         string `json:"id"`
    Code       string `json:"code"`
    Language   string `json:"language"`
    ProblemID  string `json:"problem_id"`
    Status     string `json:"status"`     
    Verdict    string `json:"verdict"`   
}
