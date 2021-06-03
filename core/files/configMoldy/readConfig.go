package moldyConfig

import (
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

type AdminProjects struct {
	Changelogs            bool
	ConventionalCommits   bool
	conventionalWorkflows bool
	GitInit               bool
	SemverMode            bool
}

type AparienceOptions struct {
	AsciiArt    bool
	ColorsMode  bool
	ProgressBar bool
}

type MoldyPackageInfo struct {
	Author      string
	Description string
	Name        string
	Url         string
	Version     string
}

type MoldyRunner struct {
	Test string
}

type ConfigStruct struct {
	AdminProjects    AdminProjects
	AparienceOptions AparienceOptions
	MoldyPackageInfo MoldyPackageInfo
	MoldyRunner      MoldyRunner
}

func Settings() ConfigStruct {
	content, _ := ioutil.ReadFile("MoldyFile.toml")
	config := ConfigStruct{}

	toml.Unmarshal(content, &config)
	return config
}
