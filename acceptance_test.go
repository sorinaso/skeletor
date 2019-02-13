package skeletor

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkeletonUpdate(t *testing.T) {
	testProjectPath := "testdata/acceptance/skeletor-update"
	skeletorYmlPath := filepath.Join(testProjectPath, "/skeletor.yml")
	template1Path := filepath.Join(testProjectPath, "skeletons/test/tested/template1")
	template2Path := filepath.Join(testProjectPath, "skeletons/test/tested/template2")

	t.Run(skeletorYmlPath, func(t *testing.T) {
		project, err := NewProject(skeletorYmlPath)

		if err != nil {
			t.Errorf(err.Error())
		}

		project.Create()

		assert.DirExists(t, testProjectPath)
		assert.FileExists(t, template1Path)
		assert.FileExists(t, template2Path)
	})

	template1, err := ioutil.ReadFile(template1Path)

	if err != nil {
		t.Error(err)
	}

	template2, err := ioutil.ReadFile(template2Path)

	if err != nil {
		t.Error(err)
	}

	fmt.Println("template1:\n", string(template1))

	fmt.Println("template2:\n", string(template2))

	os.RemoveAll(filepath.Join(testProjectPath, "skeletons"))
}
