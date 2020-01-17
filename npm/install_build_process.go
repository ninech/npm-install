package npm

import (
	"os"
	"path/filepath"

	"github.com/cloudfoundry/packit/pexec"
)

func NewInstallBuildProcess(executable Executable) InstallBuildProcess {
	return InstallBuildProcess{
		executable: executable,
	}
}

type InstallBuildProcess struct {
	executable Executable
}

func (r InstallBuildProcess) Run(layerDir, cacheDir, workingDir string) error {
	err := os.Mkdir(filepath.Join(layerDir, "node_modules"), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Symlink(filepath.Join(layerDir, "node_modules"), filepath.Join(workingDir, "node_modules"))
	if err != nil {
		return err
	}

	_, _, err = r.executable.Execute(pexec.Execution{
		Args: []string{"install", "--unsafe-perm", "--cache", cacheDir},
		Dir:  workingDir,
	})
	if err != nil {
		return err
	}

	return nil
}
