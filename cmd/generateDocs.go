package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func GenerateDocuments() {
	fmt.Println("generating documents")
	err := doc.GenMarkdownTree(initCmd, "docs")
	if err != nil {
		cobra.CompError(err.Error())
		return
	}
}
