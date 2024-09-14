package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/go-github/v64/github"
	"github.com/koki-sato/github-card/internal/card"
	"github.com/spf13/cobra"
)

type Flag struct {
	repo     string
	output   string
	fullName bool
}

var flag Flag

// rootCmd represents the base command when called without any subcommands
func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "github-card --repo <owner>/<repo>",
		Short: "GitHub Repository Card Generator",
		RunE:  run,
	}
	cmd.CompletionOptions.DisableDefaultCmd = true
	cmd.Flags().StringVarP(&flag.repo, "repo", "", "", "GitHub repository name (<owner>/<repo>)")
	cmd.Flags().StringVarP(&flag.output, "output", "o", "", "Output file path (default: <owner>-<repo>.svg)")
	cmd.Flags().BoolVarP(&flag.fullName, "full-name", "f", false, "Use full name in the card")
	cmd.MarkFlagRequired("repo") /* #nosec G104 */
	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if len(strings.Split(flag.repo, "/")) != 2 {
		return fmt.Errorf("invalid repository name: %s", flag.repo)
	}
	owner := strings.Split(flag.repo, "/")[0]
	repo := strings.Split(flag.repo, "/")[1]

	// Get GitHub repository information.
	client := github.NewClient(nil)
	if os.Getenv("GITHUB_TOKEN") != "" {
		client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_TOKEN"))
	}
	repository, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return err
	}

	// Get the number of commits.
	// See https://stackoverflow.com/questions/27931139/how-to-use-github-v3-api-to-get-commit-count-for-a-repo/70610670#70610670
	_, resp, err := client.Repositories.ListCommits(ctx, owner, repo, &github.CommitsListOptions{ListOptions: github.ListOptions{PerPage: 1, Page: 1}})
	if err != nil {
		return err
	}
	re := regexp.MustCompile(`<https://api.github.com/repositories/\d+/commits\?page=(\d+)&per_page=\d+>; rel="last"`)
	matches := re.FindStringSubmatch(resp.Header["Link"][0])
	if len(matches) != 2 {
		return fmt.Errorf("unable to find the number of commits")
	}
	nCommits, err := strconv.Atoi(matches[1])
	if err != nil {
		return err
	}

	// Generate a SVG from the repository information.
	svg := card.GenerateSVG(repository, nCommits, card.Option{UsesFullName: flag.fullName})

	// Write the SVG to the output file.
	outFile := flag.output
	if outFile == "" {
		outFile = fmt.Sprintf("%s-%s.svg", owner, repo)
	}
	if err = os.WriteFile(outFile, []byte(svg), 0644); err != nil /* #nosec G306 */ {
		log.Fatalf("unable to write file: %v", err)
	}

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cmd := rootCmd()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
