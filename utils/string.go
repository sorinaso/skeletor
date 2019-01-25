package utils

import "strings"

type String struct {
	Content string
}

func StringUtils(str string) String {
	return String{Content: str}
}

func (s *String) MustNotBeBlank() {
	if len(strings.TrimSpace(s.Content)) == 0 {
		panic("The string can't be blank")
	}
}
