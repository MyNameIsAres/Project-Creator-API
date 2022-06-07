package validate

import (
	"github.com/go-playground/validator/v10"
	"github.com/org/project_creator_api/models"
)

func Validate(request models.Project) error {
	err := validator.New().Struct(request)

	if err != nil {
		return err
	}
	return nil
}
