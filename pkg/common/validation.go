package common

import (
	"fmt"
	"net/http"
)

func (v *LoginRequest) Bind(r *http.Request) error {
	if v.Id == "" || v.Pw == "" {
		return fmt.Errorf("missing field")
	}
	return nil
}

func (v *Template) Bind(r *http.Request) error {
	if v.Body == "" || v.Name == "" || len(v.Variables) == 0 {
		return fmt.Errorf("missing field")
	}
	return nil
}
