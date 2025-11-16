package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [module path], [project creation folder]",
	Short: "Initialize a new Project-Zero template",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		modulePath := args[0]

		creationFolder := "./temp"
		if len(args) == 2 {
			creationFolder = args[1]
		}

		fmt.Println("Module Path:", modulePath)
		fmt.Println("Project Folder:", creationFolder)

		err := copyAndReplace("./code", creationFolder, "MODULE_PATH", modulePath)
		if err != nil {
			fmt.Println("Error copying template:", err)
			return
		}

		fmt.Println("Project initialized successfully!")
	},
}

func copyAndReplace(srcDir, dstDir, oldModule, newModule string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Determine the target path in the destination directory
		relPath, _ := filepath.Rel(srcDir, path)
		relPath = filepath.ToSlash(relPath)

		// Ignore the modules folder
		if strings.HasPrefix(relPath, "pkg/backend/modules") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		targetPath := filepath.Join(dstDir, relPath)

		// If directory, create it
		if info.IsDir() {
			return os.MkdirAll(targetPath, os.ModePerm)
		}

		// Read the file contents
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Replace module placeholders
		content := strings.ReplaceAll(string(data), oldModule, newModule)

		// Write to the destination
		return os.WriteFile(targetPath, []byte(content), 0644)
	})
}

func init() {
	RootCmd.AddCommand(initCmd)
}
