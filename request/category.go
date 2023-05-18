package request

import validation "github.com/go-ozzo/ozzo-validation"

type Category struct {
	Name string `json:"name" form:"name"`
}

func (req Category) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required),
	)
}
