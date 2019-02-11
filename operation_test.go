package skeletor

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"skeletor/utils"
	"testing"
)

var TestTextTemplateOperationExecuteCases = []struct {
	title    string
	to       textTemplateOperation
	mustPass bool
}{
	{
		title: "One existent template without parameters",
		to: textTemplateOperation{
			templates_directory: "./testdata/text_template/one_template_without_params",
			skeletons_directory: "/tmp",
			environment:         Environment{"a": "b"},
			template: Template{
				Source: "template1.tpl",
				Target: "/tmp/text_template_test_without_parameters",
			},
		},
		mustPass: true,
	},
	{
		title: "One existent template without parameters wihtout existing directory",
		to: textTemplateOperation{
			templates_directory: "./testdata/text_template/one_template_without_params",
			skeletons_directory: "/tmp",
			environment:         Environment{"a": "b"},
			template: Template{
				Source: "template1.tpl",
				Target: "/tmp/noexiste/text_template_test_without_parameters",
			},
		},
		mustPass: true,
	},
	{
		title: "One existent template with parameters",
		to: textTemplateOperation{
			templates_directory: "./testdata/text_template/one_template_with_params",
			skeletons_directory: "/tmp",
			environment:         Environment{"a": "b"},
			template: Template{
				Source: "template1.tpl",
				Target: "/tmp/text_template_test_with_parameter",
			},
		},
		mustPass: true,
	},

	// Nested directory
	{
		title: "One existent template without parameters nested",
		to: textTemplateOperation{
			templates_directory: "./testdata/text_template/one_template_without_params",
			skeletons_directory: "/tmp",
			environment:         Environment{"a": "b"},
			template: Template{
				Source: "template1.tpl",
				Target: "text_template_test_without_parameters_nested",
			},
		},
		mustPass: true,
	},
	{
		title: "One existent template without parameters wihtout existing directory nested",
		to: textTemplateOperation{
			templates_directory: "./testdata/text_template/one_template_without_params",
			skeletons_directory: "/tmp",
			environment:         Environment{"a": "b"},
			template: Template{
				Source: "nested/template1.tpl",
				Target: "/noexiste/text_template_test_without_parameters_nested",
			},
		},
		mustPass: true,
	},
	{
		title: "One existent template with parameters nested",
		to: textTemplateOperation{
			templates_directory: "./testdata/text_template/one_template_with_params",
			skeletons_directory: "/tmp",
			environment:         Environment{"a": "b"},
			template: Template{
				Source: "nested/template1.tpl",
				Target: "text_template_test_with_parameters_nested",
			},
		},
		mustPass: true,
	},
}

func TestTextTemplateOperationExecute(t *testing.T) {
	for _, tt := range TestTextTemplateOperationExecuteCases {
		t.Run(tt.title+"/no_preexistent_file", func(t *testing.T) {
			err := os.Remove(tt.to.GetDestinationFilePath())

			if err != nil {
				t.Log(err.Error())
			}

			to := NewTextTemplateOperation(
				tt.to.templates_directory,
				tt.to.skeletons_directory,
				Environment{"a": "b"},
				tt.to.template,
			)

			err = to.Execute()

			if tt.mustPass {
				if err != nil {
					t.Error(err.Error())
				}

				utils.AssertGoldenTwoFiles(t, to.GetDestinationFilePath(), filepath.Join(to.templates_directory, "result.golden"))
			} else {
				//Must Implement
			}
		})

		t.Run(tt.title+"/preexistent_file", func(t *testing.T) {
			to := NewTextTemplateOperation(
				tt.to.templates_directory,
				tt.to.skeletons_directory,
				Environment{"a": "b"},
				tt.to.template,
			)

			err := ioutil.WriteFile(to.GetDestinationFilePath(), []byte("hello\ngohkjfdhkdshfkjsfdhkjdfhsd\n"), 0644)

			err = to.Execute()

			if tt.mustPass {
				if err != nil {
					t.Error(err.Error())
				}

				utils.AssertGoldenTwoFiles(t, to.GetDestinationFilePath(), filepath.Join(to.templates_directory, "result.golden"))
			} else {
				//Must Implement
			}
		})

	}
}
