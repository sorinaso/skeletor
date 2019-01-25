package utils

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

func yamlFileToMap(path string) (map[interface{}]interface{}, error) {
	var ret map[interface{}]interface{}
	yamlFile, err := ioutil.ReadFile(path)

	if err != nil {
		return map[interface{}]interface{}{}, err
	}

	err = yaml.Unmarshal(yamlFile, &ret)

	if err != nil {
		return map[interface{}]interface{}{}, err
	}

	return ret, nil
}
