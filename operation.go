package skeletor

import (
	"skeletor/utils"
	"os"
	"path/filepath"
	"text/template"

	log "github.com/sirupsen/logrus"
)

type Operation interface {
	Execute() error
}

type textTemplateOperation struct {
	templates_directory string
	skeletons_directory string
	name                string
	environment         Environment
	destination         string
}

func NewTextTemplateOperation(
	templates_directory string,
	skeletons_directory string,
	name string,
	environment Environment,
	destination string,
) textTemplateOperation {
	ret := textTemplateOperation{}

	ret.templates_directory = templates_directory
	ret.skeletons_directory = skeletons_directory
	ret.name = name
	ret.environment = environment
	ret.destination = destination

	return ret
}

func (to textTemplateOperation) GetDestinationFilePath() string {
	dstPath := utils.PathUtils(to.skeletons_directory)

	dstPath.Join(to.destination)

	return dstPath.Path
}

func (to textTemplateOperation) Execute() error {
	dstPath := to.GetDestinationFilePath()

	dstDirectory := utils.PathUtils(filepath.Dir(dstPath))

	if err := dstDirectory.CheckAbsolute(); err != nil {
		return err
	}

	if err := dstDirectory.CreateDirectory(); err != nil {
		return err
	}

	dstFile, err := os.OpenFile(dstPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0660)

	if err != nil {
		return err
	}

	log.Info("Obtaining template list from directory " + to.templates_directory)

	rootTpl, err := to.getRootTemplate()

	log.Info("Templates: ", *rootTpl.Templates()[0])

	if err != nil {
		return err
	}

	log.Info("Rendering template ", to.name, " to ", dstPath, "(Environment: ", to.environment, ")")

	err = rootTpl.ExecuteTemplate(dstFile, to.name, to.environment)

	if err != nil {
		return err
	}

	return nil
}

func (to *textTemplateOperation) getRootTemplate() (*template.Template, error) {
	ret, err := utils.GetTemplatesFromPath(
		to.templates_directory,
		template.FuncMap{"environment": func() Environment { return to.environment }},
	)

	if err != nil {
		return &template.Template{}, err
	}

	return ret, nil
}
