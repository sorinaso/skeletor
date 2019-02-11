package skeletor

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkeletonUpdate(t *testing.T) {
	testProjectPath := "testdata/acceptance/skeletor-update"
	skeletorYmlPath := filepath.Join(testProjectPath, "/skeletor.yml")

	t.Run(skeletorYmlPath, func(t *testing.T) {
		project, err := NewProject(skeletorYmlPath)

		if err != nil {
			t.Errorf(err.Error())
		}

		project.Create()

		assert.DirExists(t, filepath.Join(testProjectPath, "skeletons/test/tested"))
		assert.FileExists(t, filepath.Join(testProjectPath, "skeletons/test/tested/template1"))
		assert.FileExists(t, filepath.Join(testProjectPath, "skeletons/test/tested/template2"))
	})

	os.RemoveAll(filepath.Join(testProjectPath, "skeletons"))
}
