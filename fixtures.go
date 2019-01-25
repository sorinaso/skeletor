package skeletor

func FixturesBadModels() []Model {
	return []Model{
		Model{Name: "", Templates: FixturesGoodTemplates()},
		Model{Name: "test_bad_model_1", Templates: FixturesBadTemplates()},
		Model{Name: "test_bad_model_2", Templates: FixturesBadTemplates()},
	}
}

func FixturesGoodModels() []Model {
	return []Model{
		NewModel("test_good_model", FixturesGoodTemplates()),
		NewModel("test_good_model_2", FixturesGoodTemplates()),
	}
}

func FixtureGoodSkeletons() []Skeleton {
	return []Skeleton{
		Skeleton{
			Name:        "test_good_skeleton",
			Model:       FixturesGoodModels()[0],
			Environment: Environment{"test_good_skeleton": "tetest_good_skeleton"},
		},
		Skeleton{
			Name:        "test_good_skeleton_2",
			Model:       FixturesGoodModels()[1],
			Environment: Environment{"test_good_skeleton_2": "tetest_good_skeletont_2"},
		},
	}
}

func FixturesBadTemplates() []Template {
	return []Template{
		Template{},
		Template{Source: "test_bad_template"},
		Template{Target: "test_bad_template_2"},
		Template{Source: "/test_bad_template_3", Target: "test_bad_template_3"},
		Template{Source: "test_bad_template_4", Target: "/test_bad_template_4"},
		Template{Source: "/test_bad_template_5", Target: "/test_bad_template_5"},
	}
}

func FixturesGoodTemplates() []Template {
	return []Template{
		NewTemplate("test_good_template", "test_good_template"),
		NewTemplate("test_good_template_2", "test_good_template_2"),
	}
}

func FixtureProjectGoodConfigYAML() Project {
	yamlWorkspaces := ProjectWorkspaces{
		Templates: NewWorkspace("/tp"),
		Skeletons: NewWorkspace("/sp"),
	}

	yamlModel := NewModel("tm", []Template{Template{Source: "tm_t1_src", Target: "tm_t1_tgt"}})

	yamlModel2 := NewModel(
		"tm2",
		[]Template{
			Template{Source: "tm2_t1_src", Target: "tm2_t1_tgt"},
			Template{Source: "tm2_t2_src", Target: "tm2_t2_tgt"},
		},
	)

	yamlSkeleton := NewSkeleton("test", yamlModel, Environment{"test": "test"})

	for _, t := range yamlSkeleton.Model.Templates {
		yamlSkeleton.Operations = append(
			yamlSkeleton.Operations,
			NewTextTemplateOperation(
				yamlWorkspaces.Templates.Path,
				yamlWorkspaces.Skeletons.Path,
				t.Source,
				yamlSkeleton.Environment,
				t.Target,
			),
		)
	}

	yamlSkeleton2 := NewSkeleton("test2", yamlModel2, Environment{"test2": "test2"})

	for _, t := range yamlSkeleton2.Model.Templates {
		yamlSkeleton2.Operations = append(
			yamlSkeleton2.Operations,
			NewTextTemplateOperation(
				yamlWorkspaces.Templates.Path,
				yamlWorkspaces.Skeletons.Path,
				t.Source,
				yamlSkeleton2.Environment,
				t.Target,
			),
		)
	}

	return Project{
		Models:     []Model{yamlModel, yamlModel2},
		Skeletons:  []Skeleton{yamlSkeleton, yamlSkeleton2},
		Workspaces: yamlWorkspaces,
	}
}
