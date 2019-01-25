package skeletor

import (
	"skeletor/utils"
)

type Config struct {
	Path string

	TemplatesPath string `yaml:"templates_path"`
	SkeletonsPath string `yaml:"skeletons_path"`

	Models []struct {
		Name      string
		Templates []Template
	}

	Skeletons []struct {
		Name        string
		Model       string
		Environment Environment
	}
}

func NewConfig(configurationPath string) (Config, error) {
	ret := Config{}

	cpu := utils.PathUtils(configurationPath)

	utils.MustNotReturnAnError(cpu.CheckFile())

	ret.Path = cpu.Path

	p := utils.PathUtils(configurationPath)

	if err := p.UnmarshalYAML(&ret); err != nil {
		return Config{}, err
	}

	return ret, nil
}
