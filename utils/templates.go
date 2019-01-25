package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

func GetTemplatesFromPath(path string, funcMap template.FuncMap) (*template.Template, error) {
	p := PathUtils(path)
	if err := p.CheckDirectory(); err != nil {
		return &template.Template{}, err
	}

	cleanRoot := filepath.Clean(path)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() {
			if e1 != nil {
				return e1
			}

			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx:]
			t := root.New(name).Funcs(funcMap)
			t, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}

		return nil
	})

	return root, err
}
