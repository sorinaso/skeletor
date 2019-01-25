package main

import (
	"fmt"
	"skeletor"
	//	"skeletor/templates"
	//	"os"
)

func main() {
	project, err := skeletor.NewProject("resources/config.yml")

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", project.GetSkeletons())

	if err := project.Create(); err != nil {
		panic(err.Error())
	}
	// tpl, err := templates.NewYAMLFromResource("openshift/buildconfig/symfony-with-pipeline.yml")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }

	// tpl2, err := templates.NewYAMLFromFields(tpl.Fields)

	// fmt.Println(tpl2.Content)
	//fmt.Println("test")
}
