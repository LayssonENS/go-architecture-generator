package domain

import "errors"

var (
	// ErrRegistrationNotFound will throw if the searched person registration cannot be found
	ErrRegistrationNotFound = errors.New("person registration not found")
)

type ErrorResponse struct {
	ErrorMessage string `json:"error"`
}

type ProjectConfig struct {
	ProjectName string `json:"projectName"`
	Framework   string `json:"framework"`
	Database    string `json:"database"`
	Auth        bool   `json:"auth"`
	Cache       bool   `json:"cache"`
	StructName  string `json:"structName"`
}
