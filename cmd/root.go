package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v64/github"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github-card <owner>/<repo>",
	Short: "GitHub Repository Card Generator",
	Args:  cobra.ExactArgs(1),
	RunE:  run,
}

func run(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if len(strings.Split(args[0], "/")) != 2 {
		return fmt.Errorf("invalid repository name: %s", args[0])
	}
	owner := strings.Split(args[0], "/")[0]
	repo := strings.Split(args[0], "/")[1]

	client := github.NewClient(nil)
	repository, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return err
	}
	fmt.Printf("repository: %v\n", repository)

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here, will be global for your application.
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
