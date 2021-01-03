package server

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
