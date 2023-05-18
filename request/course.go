package request

import (
	"mime/multipart"

	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	Course struct {
		Title       string                `json:"title" form:"title"`
		Description string                `json:"desc" form:"desc"`
		Price       float64               `json:"price" form:"price"`
		Banner      *multipart.FileHeader `json:"banner" form:"banner"`
		CategoryID  uint                  `json:"category_id" form:"category_id"`
	}
)

func (req Course) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Title, validation.Required),
		validation.Field(&req.Price, validation.Required),
		validation.Field(&req.Banner, validation.Required),
		validation.Field(&req.CategoryID, validation.Required),
	)
}
