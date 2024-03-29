package skeletor

import (
	"os"
	"path/filepath"
	"skeletor/utils"
	"text/template"

	log "github.com/sirupsen/logrus"
)

type Operation interface {
	Execute() error
}

type textTemplateOperation struct {
	templates_directory string
	skeletons_directory string
	environment         Environment
	template            Template
}

func NewTextTemplateOperation(
	templates_directory string,
	skeletons_directory string,
	environment Environment,
	template Template,
) textTemplateOperation {
	ret := textTemplateOperation{}

	ret.templates_directory = templates_directory
	ret.skeletons_directory = skeletons_directory
	ret.environment = environment
	ret.template = template

	return ret
}

func (to textTemplateOperation) GetDestinationFilePath() string {
	dstPath := utils.PathUtils(to.skeletons_directory)

	dstPath.PostJoin(to.template.Target)

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

	log.Info("Obtaining template list from directory " + to.templates_directory)

	rootTpl, err := to.getRootTemplate()

	if len(rootTpl.Templates()) > 0 {
		log.Info("Templates: ", *rootTpl.Templates()[0])
	} else {
		log.Info("No hay templates en el directorio")
	}

	if err != nil {
		return err
	}

	dstPathExist := true

	if _, err := os.Stat(dstPath); os.IsNotExist(err) {
		dstPathExist = false
	}

	if to.template.Overwrite || !dstPathExist {
		dstFile, err := os.OpenFile(dstPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0660)

		if err != nil {
			return err
		}

		log.Info("Rendering template ", to.template.Source, " to ", dstPath, "(Environment: ", to.environment, ")")

		err = rootTpl.ExecuteTemplate(dstFile, to.template.Source, to.environment)
	} else {
		log.Info("Skiping template ", to.template.Source, " to ", dstPath, "(Environment: ", to.environment, ") because file exists and overwrite is false")
	}

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
