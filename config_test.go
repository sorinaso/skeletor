package skeletor

import (
	"path/filepath"
	"testing"
)

var TestNewConfigCases = []struct {
	path     string
	config   Config
	mustPass bool
}{
	{
		path: filepath.Join("testdata", "config", "good_config.yml"),
		config: Config{
			Path:          "/tmp",
			TemplatesPath: "tp",
			SkeletonsPath: "sp",
			Models: []struct {
				Name      string
				Templates []struct {
					Source string
					Target string
				}
			}{
				{
					Name: "tm",
					Templates: []struct {
						Source string
						Target string
					}{
						{Source: "tm_t1_src", Target: "tm_t1_tgt"},
					},
				},
				{
					Name: "tm2",
					Templates: []struct {
						Source string
						Target string
					}{
						{Source: "tm2_t1_src", Target: "tm2_t1_tgt"},
						{Source: "tm2_t2_src", Target: "tm2_t2_tgt"},
					},
				},
			},
			Skeletons: []struct {
				Name        string
				Model       string
				Environment Environment
			}{
				{Name: "test", Model: "test", Environment: Environment{"test": "test"}},
				{Name: "test2", Model: "test2", Environment: Environment{"test2": "test2"}},
			},
		},
		mustPass: true,
	},
	{
		path:     filepath.Join("testdata", "config", "no_existent_config.yml"),
		config:   Config{},
		mustPass: false,
	},
}

func TestNewConfig(t *testing.T) {
	for _, tt := range TestNewConfigCases {
		t.Run(tt.path, func(t *testing.T) {
			if tt.mustPass {
				assertNewConfigGood(t, tt.path, tt.config)
			} else {
				assertNewConfigBad(t, tt.path)
			}
		})
	}
}
