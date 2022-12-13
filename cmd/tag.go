package cmd

import (
	"fmt"
	"log"

	"github.com/prashantkhandelwal/boom/pkg/id3"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "tags",
	Short: "Read all tags.",
	Long:  `Read all the tags from the given MP3 file`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Reading ID Tags")
		i, err := id3.Read(args[0])
		if err != nil {
			log.Fatalln("ERROR: Unable to process the file.")
		}

		fmt.Println(i)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
