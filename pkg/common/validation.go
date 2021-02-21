package common

import (
	"fmt"
	"net/http"

	"k8s.io/apimachinery/pkg/util/validation"
)

func (v *Quota) Bind(r *http.Request) error {
	if v.Cpu < -1 || v.Instance < -1 || v.Memory < -1 || v.NvidiaGpu < -1 || v.Storage < -1 {
		return fmt.Errorf("invalid value for quota. -1 <= value")
	}
	return nil
}

func (v *CreateServerRequest) Bind(r *http.Request) error {
	if v.Description == "" || v.Name == "" || v.TemplateId == 0 {
		return fmt.Errorf("missing field")
	}
	if errs := validation.IsDNS1123Label(v.Name); len(errs) != 0 {
		tmp := "invalid name format"
		for _, err := range errs {
			tmp = fmt.Sprintf("%s:%s", tmp, err)
		}
		return fmt.Errorf(tmp)
	}
	return nil
}

func (v *Flavor) Bind(r *http.Request) error {
	if v.Cpu == 0 || v.Memory == 0 {
		return fmt.Errorf("missing flavor field")
	}
	return nil
}

func (v *LoginRequest) Bind(r *http.Request) error {
	if v.Id == "" || v.Pw == "" {
		return fmt.Errorf("missing field")
	}
	return nil
}

func (v *Template) Bind(r *http.Request) error {
	if v.Body == "" || v.Name == "" || len(v.Body) == 0 {
		return fmt.Errorf("missing field")
	}
	return nil
}
