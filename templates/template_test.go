package templates_test

import (
	"testing"

	"github.com/org/project_creator_api/templates"
)

func TestCreateTemplateError(t *testing.T) {
	templateFileName := "i_dont_exist.tmpl"
	// pathString := "<path>" // TODO Replace with test env key
	template, err := templates.CreateTemplate(templateFileName, "")

	if template != nil {
		t.Errorf("Wanted nil. Got %s instead", template.Name())
	}

	if err == nil {
		t.Errorf("Wanted error but got a template instead. Check if the file has accidentally been created.")
	}
}

func TestCreateTemplate(t *testing.T) {
	templateFileName := "foobar.tmpl"
	// templatePath := "<path>" // TODO Replace with test env key
	template, err := templates.CreateTemplate(templateFileName, "")

	if template == nil {
		t.Errorf("We expected Template not to be nil. Got %s instead", template.Name())
	}

	if err != nil {
		t.Errorf("Expected template, got %s instead", err)
	}
}
