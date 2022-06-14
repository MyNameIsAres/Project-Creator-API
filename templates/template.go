package templates

import (
	"io"
	"text/template"

	"github.com/org/project_creator_api/models"
)

func CreateTemplate(templateFile, templatePath string) (*template.Template, error) {
	templatePath = templatePath + templateFile

	template, err := template.New(templateFile).ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}

	return template, nil
}

func ApplyTemplate(template *template.Template, writer io.Writer, project models.Project) error {
	return template.Execute(writer, project)
}
