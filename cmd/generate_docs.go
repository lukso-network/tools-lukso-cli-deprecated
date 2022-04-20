package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func generateDocuments() {
	fmt.Println("generating docs")
	err := doc.GenMarkdownTree(rootCmd, "docs")
	if err != nil {
		cobra.CompError(err.Error())
		return
	}
}
