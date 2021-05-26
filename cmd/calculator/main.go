package main

import (
	"log"

	"github.com/namit-chandwani/Go-Calculator-CLI/pkg/cmdloader"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func main() {
	cmd := &cobra.Command{
		Use:          "calculator",
		Short:        "Calculator CLI",
		Long:         `A basic calculator CLI built in Golang using the Cobra library.`,
		SilenceUsage: true,
	}

	cmdloader.AddCommands(cmd)

	err := doc.GenMarkdownTree(cmd, "./docs/commands")
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
