package skeletor

import "skeletor/utils"

var ModelMap = map[string]Model{}

type Template = struct {
	Source    string
	Target    string
	Overwrite bool
}

func NewTemplate(source string, target string, overwrite bool) Template {
	ps := utils.PathUtils(source)
	utils.MustNotReturnAnError(ps.CheckRelative())

	pt := utils.PathUtils(target)
	utils.MustNotReturnAnError(pt.CheckRelative())

	return Template{Source: source, Target: target, Overwrite: overwrite}
}

type Model struct {
	Name string

	Templates []Template
}

func NewModel(name string, templates []Template) Model {
	ret := Model{}

	s := utils.StringUtils(name)
	s.MustNotBeBlank()

	ret.Name = name

	for _, tpl := range templates {
		ret.Templates = append(ret.Templates, NewTemplate(tpl.Source, tpl.Target, tpl.Overwrite))
	}

	ModelMap[ret.Name] = ret

	return ret
}
