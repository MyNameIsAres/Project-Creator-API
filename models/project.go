package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func foo() {
	validator := validator.New()
	fmt.Println(validator)
}

type Project struct {
	ProjectMeta   projectMeta  `json:"meta" validate:"required"`
	Language      string       `json:"language" validate:"required"`
	FoVersion     string       `json:"foundation_version" validate:"required"` // This might get removed as we'd naturally always want the latest version
	SpigotVersion string       `json:"spigot_version" validate:"required"`
	JavaVersion   int16        `json:"java" validate:"required"`
	BuildSystem   string       `json:"build_system" validate:"required"`
	Dependencies  dependencies `json:"dependencies" validate:"required"`
}

type projectMeta struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Artifact    string `json:"artifact" validate:"required"`
	Group       string `json:"group" validate:"required"`
}

type repositoryMeta struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type dependencyMeta struct {
	Group    string `json:"group"`
	Artifact string `json:"artifact"`
	Version  string `json:"version"`
}

type dependency struct {
	Repository repositoryMeta `json:"repository"`
	Dependency dependencyMeta `json:"dependency"`
}

// dependency struct goes here, bit complex.
type dependencies struct {
	Dependencies []dependency `json:"dependency_collection"`
}
