package filemanipulation

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var file string
var FileCmd = &cobra.Command{
	Use:     "file",
	Short:   "This for performing file manipulation",
	Aliases: []string{"filem", "FILE"},
	Run: func(cmd *cobra.Command, args []string) {
		newFile, err := os.Create(file)
		if err != nil {
			fmt.Printf("failed to create file: %s", file)
			return
		}
		fmt.Println("File created", file)
		newFile.Close()
	},
}

func init() {
	FileCmd.Flags().StringVarP(&file, "name", "n", "", "Name of file to be created")
	FileCmd.MarkFlagRequired("name")
}
