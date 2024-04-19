package main

import (
	"context"
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

//go:embed embedfs/*
var playbooks embed.FS

func main() {

	// Create a temporary directory with a prefix "ansibleplaybook-simple-embedfs-"
	tempDir, err := os.MkdirTemp("", "ansibleplaybook-simple-embedfs-")
	if err != nil {
		panic(err)
	}
	// Remove the directory when the program finishes
	defer os.RemoveAll(tempDir)

	// Copy the source directory and its contents to the destination directory
	err = copyDir(playbooks, "embedfs", tempDir)
	if err != nil {
		panic(err)
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		Inventory:  filepath.Join(tempDir, "inventory.ini"),
	}

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks(filepath.Join(tempDir, "site.yml"), filepath.Join(tempDir, "site2.yml")),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	exec := execute.NewDefaultExecute(
		execute.WithCmd(playbookCmd),
		execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
	)

	err = exec.Execute(context.TODO())
	if err != nil {
		panic(err)
	}
}

// copyDir recursively copies the directory and its contents from src to dest.
func copyDir(sourceFS embed.FS, src, dest string) error {
	// Get the contents of the source directory
	entries, err := sourceFS.ReadDir(src)
	if err != nil {
		return err
	}

	// Copy each entry in the source directory to the destination directory
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			// If the entry is a directory, recursively copy it to the destination directory
			err = copyDir(sourceFS, srcPath, destPath)
			if err != nil {
				return err
			}
		} else {
			// If the entry is a file, copy its contents to the destination file
			srcFile, err := sourceFS.Open(srcPath)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			destFile, err := os.Create(destPath)
			if err != nil {
				return err
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				return err
			}
			fmt.Printf("Copying file from the embedded filesystem '%s' to '%s'\n", srcPath, destPath)
		}
	}

	return nil
}
