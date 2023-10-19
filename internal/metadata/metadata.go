package metadata

import (
	"errors"
	"fmt"

	"github.com/pterm/pterm"
	"github.com/spf13/afero"
)

type Metadata struct {
	ProgrammingLanguage string
	ProjectName         string
	GoVersion           string
	GitBranch           string
	CICDPlatform        string
	PipelineTasks       map[string]bool
	LocalTasks          map[string]bool
}

func New() *Metadata {
	return &Metadata{}
}

func (cfg *Metadata) SetProgrammingLanguage(n string) {
	cfg.ProgrammingLanguage = n
}

func (cfg *Metadata) SetProjectName(n string, fs *afero.Fs) error {
	if n == "" {
		switch cfg.ProgrammingLanguage {
		case "go":
			fmt.Println("Project name was not set, extracting from go.mod file")
			gopn, err := cfg.ExtractProjectNameFromGoModFile(fs)
			if err != nil {
				return fmt.Errorf("failed to extract project name from go.mod file: %v", err)
			}
			pterm.Println("Project name extracted from go.mod:", pterm.Yellow(*gopn))
			cfg.ProjectName = *gopn
			return nil
		case "rust":
			fmt.Println("Project name was not set, extracting from Cargo.toml file")
			rpn, err := cfg.ExtractProjectNameFromCargoFile(fs)
			if err != nil {
				return fmt.Errorf("failed to extract project name from Cargo.toml file: %v", err)
			}
			pterm.Println("Project name extracted from go.mod:", pterm.Yellow(*rpn))
			cfg.ProjectName = *rpn
			return nil
		}
	}
	cfg.ProjectName = n
	return errors.New("something went wrong")
}

func (cfg *Metadata) SetGitBranch(n string) {
	if n == "" {
		pterm.Println(pterm.White("Git branch was not set, defaulting to ") + pterm.Yellow("main"))
		cfg.SetGitBranch("main")
		return
	}
	cfg.GitBranch = n
}

func (cfg *Metadata) SetCICDPlatform(n string) {
	cfg.CICDPlatform = n
}

func (cfg *Metadata) SetPipelineTasks(n []string) {
	cfg.PipelineTasks = make(map[string]bool)
	for _, t := range n {
		cfg.PipelineTasks[t] = true
	}
}

func (cfg *Metadata) SetLocalTasks(n []string) {
	cfg.LocalTasks = make(map[string]bool)
	for _, t := range n {
		cfg.LocalTasks[t] = true
	}
}
