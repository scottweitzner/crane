package main

import (
	"github.com/scottweitzner/crane/cmd"
)

func main() {
	rootCmd := cmd.NewRootCommand()
	rootCmd.Execute()
}
