package skeletor

import (
	"skeletor/utils"
	"path/filepath"
)

type Workspace struct {
	Path string
}

func NewWorkspace(path string) Workspace {
	p := utils.PathUtils(path)

	if err := p.CheckAbsolute(); err != nil {
		panic("NewWorkspace: " + err.Error())
	}

	return Workspace{Path: path}
}

func (w *Workspace) SubPath(path string) (string, error) {
	p := utils.PathUtils(path)

	if err := p.CheckRelative(); err != nil {
		return "", err
	}

	return filepath.Join(w.Path, path), nil
}
