package common

import "fmt"

//GetReason Create Reason Message
func GetReason(msg string) *Reason {
	return &Reason{
		Reason: msg,
	}
}

//GetReasonf Create Reason Message with formatting
func GetReasonf(format string, a ...interface{}) *Reason {
	return &Reason{
		Reason: fmt.Sprintf(format, a...),
	}
}
