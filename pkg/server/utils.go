package server

import (
	"encoding/json"
	"net/http"
)

func parseBody(r *http.Request, target interface{}) error {
	err := json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		return err
	}
	/*
		TODO: Validate if field is not "nonempty"
	*/
	return nil
}
