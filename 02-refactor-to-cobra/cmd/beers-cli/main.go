package main

import (
	"github.com/CodelyTV/golang-examples/02-refactor-to-cobra/internal/cli/beers"
	"github.com/CodelyTV/golang-examples/02-refactor-to-cobra/internal/cli/stores"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(beers.InitBeersCmd())
	rootCmd.AddCommand(stores.InitStoresCmd())
	rootCmd.Execute()
}
