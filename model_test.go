package skeletor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewModel(t *testing.T) {
	for _, m := range FixturesBadModels() {
		assert.Panics(t, func() { NewModel(m.Name, m.Templates) }, "This NewModel should panic")
	}

	for _, m := range FixturesGoodModels() {
		NewModel(m.Name, m.Templates)
	}
}

func TestGetModelByName(t *testing.T) {
	tm1 := FixturesGoodModels()[0]
	assert.Equal(t, ModelMap[tm1.Name], tm1)

	tm2 := FixturesGoodModels()[1]
	assert.Equal(t, ModelMap[tm2.Name], tm2)
}
