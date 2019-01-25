package skeletor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertNewProjectGood(t *testing.T, configurationPath string, project Project) Project {
	ret, err := NewProject(configurationPath)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, project.GetSkeletons(), ret.GetSkeletons())
	assert.Equal(t, project.GetModels(), ret.GetModels())
	assert.Equal(t, project.GetWorkSpaces(), ret.GetWorkSpaces())

	return ret
}

func assertNewProjectBad(t *testing.T, path string) error {
	_, err := NewProject(path)

	if err == nil {
		t.Error("EL repositorio con el path " + path + " deberia devolver un error")
	}

	return err
}

func assertNewConfigGood(t *testing.T, configurationPath string, config Config) Config {
	ret, err := NewConfig(configurationPath)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, config, ret)

	return ret
}

func assertNewConfigBad(t *testing.T, path string) error {
	_, err := NewConfig(path)

	if err == nil {
		t.Error("La configuracion deberia devolver un error")
	}

	return err
}
