package main

import (
	"context"
	"embed"
	"os"
	"path/filepath"

	"github.com/apenella/go-ansible/v2/examples/ansibleplaybook-embed-python/internal/ansibleplaybook-embed-python/data"
	"github.com/apenella/go-ansible/v2/pkg/execute"
	executable "github.com/apenella/go-ansible/v2/pkg/execute/exec"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
	"github.com/kluctl/go-embed-python/embed_util"
	"github.com/kluctl/go-embed-python/python"
)

//go:embed resources
var resources embed.FS

// PythonExec struct implements the executable.Exec interface and is used to execute Python commands using the embedded Python interpreter
type PythonExec struct {
	ep              *python.EmbeddedPython
	pythonLibFsPath string
	resourcesFsPath string
}

func NewPythonExec() *PythonExec {

	tmpDir := filepath.Join(os.TempDir(), "go-ansible")
	pythonDir := tmpDir + "-python"
	pythonLibDir := tmpDir + "-python-lib"
	resourcesDir := tmpDir + "-resources"

	pythonLibFs, _ := embed_util.NewEmbeddedFilesWithTmpDir(data.Data, pythonLibDir, true)
	pythonLibFsPath := pythonLibFs.GetExtractedPath()
	resourcesFs, _ := embed_util.NewEmbeddedFilesWithTmpDir(resources, resourcesDir, true)
	resourcesFsPath := resourcesFs.GetExtractedPath()

	ep, _ := python.NewEmbeddedPythonWithTmpDir(pythonDir, true)
	ep.AddPythonPath(pythonLibFs.GetExtractedPath())
	ep.AddPythonPath(resourcesFs.GetExtractedPath())

	return &PythonExec{
		ep:              ep,
		pythonLibFsPath: pythonLibFsPath,
		resourcesFsPath: resourcesFsPath,
	}
}

// Command is a wrapper of exec.Command
func (e *PythonExec) Command(name string, arg ...string) executable.Cmder {

	cmd := append([]string{name}, arg...)

	return e.ep.PythonCmd2(cmd)
}

// CommandContext is a wrapper of exec.CommandContext
func (e *PythonExec) CommandContext(ctx context.Context, name string, arg ...string) executable.Cmder {
	cmd := append([]string{name}, arg...)
	return e.ep.PythonCmd2(cmd)
}

func main() {

	var err error

	epExec := NewPythonExec()

	//
	// You can use the following line to install additional ansible collections or roles
	//
	// galaxyEP := ep.PythonCmd("-m", "ansible", "galaxy", "collection", "install", "community.general", "--force")
	// galaxyEP.Stdout = os.Stdout
	// galaxyEP.Stderr = os.Stderr
	// err = galaxyEP.Run()
	// if err != nil {
	// 	panic(err)
	// }

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory:  "127.0.0.1,",
		Connection: "local",
	}

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks(
			filepath.Join(epExec.resourcesFsPath, "resources", "ansible", "site.yml"),
		),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
		playbook.WithBinary(
			filepath.Join(epExec.pythonLibFsPath, "bin", "ansible-playbook"),
		),
	)

	ansiblePlaybookExecutor := execute.NewDefaultExecute(
		execute.WithCmd(playbookCmd),
		execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
		execute.WithExecutable(epExec),
	)

	err = ansiblePlaybookExecutor.Execute(context.TODO())
	if err != nil {
		panic(err)
	}
}
