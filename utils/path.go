package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type Path struct {
	Path string
}

func PathUtils(path string) Path {
	return Path{Path: path}
}

func (p *Path) CheckAbsolute() error {
	if err := p.CheckValid(); err != nil {
		return err
	}

	if strings.HasPrefix(p.Path, "/") {
		return nil
	}

	return errors.New("The path " + p.Path + " must be absolute.")
}

func (p *Path) MustBeRelative() {
	if err := p.CheckRelative(); err != nil {
		panic(err.Error())
	}
}

func (p *Path) CheckRelative() error {
	if err := p.CheckValid(); err != nil {
		return err
	}

	if strings.HasPrefix(p.Path, "/") {
		return errors.New("The path " + p.Path + " must be relative.")
	}

	return nil
}

func (p *Path) CheckFile() error {
	if err := p.CheckValid(); err != nil {
		return err
	}

	stat, err := os.Stat(p.Path)

	if os.IsNotExist(err) {
		return errors.New("El path " + p.Path + " no existe")
	}

	if stat.IsDir() {
		return errors.New("El path " + p.Path + " is a directory not a file")
	}

	return nil
}

func (p *Path) MustBeARelativeFile() {
	if err := p.CheckRelativeFile(); err != nil {
		panic(err.Error())
	}
}

func (p *Path) CheckRelativeFile() error {
	if err := p.CheckFile(); err != nil {
		return err
	}

	if err := p.CheckRelative(); err != nil {
		return err
	}

	return nil
}

func (p *Path) CheckDirectory() error {
	if err := p.CheckValid(); err != nil {
		return err
	}

	stat, err := os.Stat(p.Path)

	if os.IsNotExist(err) {
		return errors.New("El path " + p.Path + " no existe")
	}

	if !stat.IsDir() {
		return errors.New("El path " + p.Path + " is a file not a directory")
	}

	return nil
}

func (p *Path) MustBeARelativeDirectory() {
	if err := p.CheckRelativeDirectory(); err != nil {
		panic(err.Error())
	}
}

func (p *Path) CheckRelativeDirectory() error {
	if err := p.CheckDirectory(); err != nil {
		return err
	}

	if err := p.CheckRelative(); err != nil {
		return err
	}

	return nil
}

func (p *Path) MustBeAAbsoluteDirectory() {
	if err := p.CheckAbsoluteDirectory(); err != nil {
		panic(err.Error())
	}
}

func (p *Path) CheckAbsoluteDirectory() error {
	if err := p.CheckDirectory(); err != nil {
		return err
	}

	if err := p.CheckAbsolute(); err != nil {
		return err
	}

	return nil
}

func (p *Path) CheckValid() error {
	if p.Path == "" {
		return errors.New("The path can't be blank")
	}

	return nil
}

func (p *Path) Join(path string) {
	p.Path = filepath.Join(p.Path, path)
}

func (p *Path) CreateDirectory() error {
	if err := os.MkdirAll(p.Path, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func (p *Path) UnmarshalYAML(out interface{}) error {
	if err := p.CheckFile(); err != nil {
		return err
	}

	yamlFile, err := ioutil.ReadFile(p.Path)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(yamlFile, out)
}
