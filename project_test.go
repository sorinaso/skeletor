package skeletor

import (
	"path/filepath"
	"testing"
)

func TestNewProject(t *testing.T) {
	testCases := []struct {
		path     string
		project  Project
		mustPass bool
	}{
		{
			path:     filepath.Join("testdata", "project", "good_project", "config.yml"),
			project:  FixtureProjectGoodConfigYAML(),
			mustPass: true,
		},
		{
			path:     filepath.Join("testdata", "project", "bad_project1", "config.yml"),
			project:  Project{},
			mustPass: false,
		},
		{
			path:     filepath.Join("testdata", "project", "bad_project2", "config.yml"),
			project:  Project{},
			mustPass: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.path, func(t *testing.T) {
			if tt.mustPass {
				assertNewProjectGood(t, tt.path, tt.project)
			} else {
				assertNewProjectBad(t, tt.path)
			}
		})
	}
}
