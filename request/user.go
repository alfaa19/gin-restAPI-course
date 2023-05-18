package request

import validation "github.com/go-ozzo/ozzo-validation"

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"-"`
}

func (req User) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Username, validation.Required),
		validation.Field(&req.Password, validation.Required),
	)
}
