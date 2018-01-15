package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type AppStringArray []string

type AppGenDependency struct {
	Name      string `yaml:"name"`
	URL       string `yaml:"url"`
	Copyright string `yaml:"copyright"`
	License   string `yaml:"license"`
	Info      string `yaml:"info"`
}
type AppGenDependencyArray []AppGenDependency


type AppGenTarget struct {
	Id           string                `yaml:"id"`
	Image        string                `yaml:"image"`
	ShortDesc    string                `yaml:"shortdesc"`
	Desc         string                `yaml:"desc"`
	Tags         AppStringArray        `yaml:"tags"`
	Dependencies AppGenDependencyArray `yaml:"dependencies"`
	CodeGenInfo  string				   `yaml:"codegeninfo"`
}

type AppGenTargetArray []AppGenTarget

type AppGenTargets struct {
	Targets AppGenTargetArray `yaml:"targets"`
}

func (agt *AppGenTargets) AppGenTargetById(id string) *AppGenTarget {
	for _, target := range agt.Targets {
		if id == target.Id {
			return &target
		}
	}
	return nil
}

func ReadGeneratorsConfig() (*AppGenTargets, error) {
	b, err := ioutil.ReadFile(ServerConfig.Paths.GeneratorsDataPath)
	if err != nil {
		return nil, err
	}

	var t = &AppGenTargets{}
	if err := yaml.Unmarshal(b, t); err != nil {
		Error.Printf("Error reading generators config file. Check YAML: %s", err)
		return t, err
	}

//	Debug.Printf("targets=%s\n", spew.Sdump(t))
	return t, nil
}

