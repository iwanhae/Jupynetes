package ent

import "github.com/iwanhae/Jupynetes/pkg/common"

//ToCommonType Convert db schema to common type
func (v *Template) ToCommonType() *common.Template {
	res := &common.Template{
		Id:          v.ID,
		Name:        v.Name,
		Description: v.Description,
		Body:        v.Template,
		Variables:   common.TemplateVariables{},
	}
	if v.Variables == nil {
		return res
	}
	for _, variable := range *v.Variables {
		res.Variables = append(res.Variables,
			common.TemplateVariable{
				Name:  variable.Name,
				Value: variable.Value,
			},
		)
	}

	return res
}
