package skeletor

import (
	"path/filepath"
	"skeletor/utils"
)

type ProjectWorkspaces = struct {
	Templates Workspace
	Skeletons Workspace
}

type Project struct {
	Config     Config
	Workspaces ProjectWorkspaces
	Skeletons  map[string]Skeleton
	Models     []Model
}

func NewProject(configurationPath string) (Project, error) {
	var config Config
	var ret Project

	config, err := NewConfig(configurationPath)

	if err != nil {
		return Project{}, err
	}

	pws, err := getWorkspacesFromConfig(config)

	if err != nil {
		return Project{}, err
	}

	models := getModelsFromConfig(config)
	skeletons := getSkeletonsFromConfig(config, pws)

	ret.Config = config
	ret.Workspaces = pws
	ret.Skeletons = skeletons
	ret.Models = models

	return ret, nil
}

func (pr *Project) GetWorkSpaces() ProjectWorkspaces {
	return pr.Workspaces
}

func (pr *Project) GetSkeletons() map[string]Skeleton {
	return pr.Skeletons
}

func (pr *Project) GetModels() []Model {
	return pr.Models
}

func (pr *Project) Create() error {
	for _, skeleton := range pr.GetSkeletons() {
		if err := skeleton.Create(); err != nil {
			return err
		}
	}

	return nil
}

func getSkeletonsFromConfig(config Config, projectWorkspaces ProjectWorkspaces) map[string]Skeleton {
	ret := map[string]Skeleton{}

	for _, skeleton := range config.Skeletons {
		if model, ok := ModelMap[skeleton.Model]; ok {
			newSkeleton := NewSkeleton(skeleton.Name, model, skeleton.Environment)

			for _, t := range newSkeleton.Model.Templates {
				newSkeleton.Operations = append(
					newSkeleton.Operations,
					NewTextTemplateOperation(
						projectWorkspaces.Templates.Path,
						filepath.Join(projectWorkspaces.Skeletons.Path, newSkeleton.Name),
						t.Source, newSkeleton.Environment,
						t.Target,
					),
				)
			}

			ret[newSkeleton.Name] = newSkeleton
		} else {
			panic("The model " + skeleton.Model + " is not defined.")
		}

	}

	return ret
}

func getModelsFromConfig(config Config) []Model {
	var ret []Model

	for _, model := range config.Models {
		ret = append(ret, NewModel(model.Name, model.Templates))
	}

	return ret
}

func getWorkspacesFromConfig(config Config) (ProjectWorkspaces, error) {
	templatesPath := utils.PathUtils(config.TemplatesPath)
	skeletonPath := utils.PathUtils(config.SkeletonsPath)
	configPathDirectory := filepath.Dir(config.Path)

	configPathDirectory, err := filepath.Abs(configPathDirectory)

	if err != nil {
		return ProjectWorkspaces{}, err
	}

	if templatesPath.IsRelative() {
		templatesPath.PreJoin(configPathDirectory)
	}

	if skeletonPath.IsRelative() {
		skeletonPath.PreJoin(configPathDirectory)
	}

	tw := NewWorkspace(templatesPath.Path)
	sw := NewWorkspace(skeletonPath.Path)

	return ProjectWorkspaces{Templates: tw, Skeletons: sw}, nil
}
