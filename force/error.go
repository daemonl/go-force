package force

import "strings"

// Custom Error to handle salesforce api responses.
type ApiErrors []*ApiError

type ApiError struct {
	Fields           []string `json:"fields,omitempty" force:"fields,omitempty"`
	Message          string   `json:"message,omitempty" force:"message,omitempty"`
	ErrorCode        string   `json:"errorCode,omitempty" force:"errorCode,omitempty"`
	ErrorName        string   `json:"error,omitempty" force:"error,omitempty"`
	ErrorDescription string   `json:"error_description,omitempty" force:"error_description,omitempty"`
}

func (e ApiErrors) Error() string {
	return e.String()
}

func (e ApiErrors) String() string {
	s := make([]string, len(e))
	for i, err := range e {
		s[i] = err.String()
	}

	return strings.Join(s, "\n")
}

func (e ApiErrors) Validate() bool {
	if len(e) != 0 {
		return true
	}

	return false
}

func (e ApiError) Error() string {
	return e.String()
}

func (e ApiError) String() string {
	parts := []string{}
	if len(e.ErrorName) > 0 {
		parts = append(parts, e.ErrorName)
	}
	if len(e.ErrorDescription) > 0 {
		parts = append(parts, e.ErrorDescription)
	}
	if len(e.Message) > 0 {
		parts = append(parts, e.Message)
	}
	if len(e.Fields) > 0 {
		parts = append(parts, strings.Join(e.Fields, ", "))
	}

	if len(e.ErrorCode) > 0 {
		return e.ErrorCode + ": " + strings.Join(parts, " ")
	}
	return strings.Join(parts, " ")
}

func (e ApiError) Validate() bool {
	if len(e.Fields) != 0 || len(e.Message) != 0 || len(e.ErrorCode) != 0 || len(e.ErrorName) != 0 || len(e.ErrorDescription) != 0 {
		return true
	}

	return false
}
