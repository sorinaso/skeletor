package utils

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertGoldenTwoFiles(t *testing.T, actualFilePath, goldenFilePath string) {
	actual, err := ioutil.ReadFile(actualFilePath)

	if err != nil {
		t.Error(err.Error())
	}

	expected, err := ioutil.ReadFile(goldenFilePath)

	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, actual, expected)
}
