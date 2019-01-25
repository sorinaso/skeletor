package skeletor

type ProjectWorkspaces = struct {
	Templates Workspace
	Skeletons Workspace
}

type Project struct {
	Config     Config
	Workspaces ProjectWorkspaces
	Skeletons  []Skeleton
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

func (pr *Project) GetSkeletons() []Skeleton {
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

func getSkeletonsFromConfig(config Config, projectWorkspaces ProjectWorkspaces) []Skeleton {
	var ret []Skeleton

	for _, skeleton := range config.Skeletons {
		if model, ok := ModelMap[skeleton.Model]; ok {
			newSkeleton := NewSkeleton(skeleton.Name, model, skeleton.Environment)

			for _, t := range newSkeleton.Model.Templates {
				newSkeleton.Operations = append(
					newSkeleton.Operations,
					NewTextTemplateOperation(
						projectWorkspaces.Templates.Path,
						projectWorkspaces.Skeletons.Path,
						t.Source, newSkeleton.Environment,
						t.Target,
					),
				)
			}

			ret = append(ret, newSkeleton)
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
	tw := NewWorkspace(config.TemplatesPath)
	sw := NewWorkspace(config.SkeletonsPath)

	return ProjectWorkspaces{Templates: tw, Skeletons: sw}, nil
}
