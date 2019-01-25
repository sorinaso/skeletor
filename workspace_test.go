package skeletor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkspaceSubpath(t *testing.T) {
	testCases := []struct {
		workspacePath    string
		workspaceSubPath string
		expectedSubpath  string
		mustPass         bool
	}{
		{workspacePath: "/tmp", workspaceSubPath: "test", expectedSubpath: "/tmp/test", mustPass: true},
		{workspacePath: "/tmp", workspaceSubPath: "/test", expectedSubpath: "", mustPass: false},
	}

	for _, tt := range testCases {
		t.Run(tt.workspacePath+" , "+tt.workspaceSubPath, func(t *testing.T) {

			if tt.mustPass {
				w := NewWorkspace(tt.workspacePath)

				subPath, err := w.SubPath(tt.workspaceSubPath)

				if err != nil {
					t.Error(err)
				}

				assert.Equal(t, tt.expectedSubpath, subPath)
			} else {
				w := NewWorkspace(tt.workspacePath)

				subPath, err := w.SubPath(tt.workspaceSubPath)

				if err == nil {
					t.Error("The subpath ", subPath, " must throw an error.")
				}
			}
		})
	}
}
