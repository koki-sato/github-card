package card

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"

	humanize "github.com/dustin/go-humanize"
	"github.com/google/go-github/v64/github"
)

//go:embed templates/card.svg
var svgTemplate string

type Option struct {
	UsesFullName   bool
	DisplayCommits bool
}

func GenerateSVG(repository *github.Repository, nCommits int, option Option) (string, error) {
	repoNameInImage := repository.GetName()
	if option.UsesFullName {
		repoNameInImage = fmt.Sprintf("%s/%s", repository.GetOwner().GetLogin(), repository.GetName())
	}

	commits := ""
	if option.DisplayCommits {
		commits = humanize.Comma(int64(nCommits))
	}

	color := LanguageToColor[repository.GetLanguage()]
	description := wrapDescription(repository.GetDescription())

	width := 400
	height := 60 + (20 * len(description)) + 35
	card := Card{
		Width:       width,
		Height:      height,
		RepoURL:     repository.GetHTMLURL(),
		RepoName:    repoNameInImage,
		Description: description,
		Color:       color,
		Language:    repository.GetLanguage(),
		Stars:       formatCount(repository.GetStargazersCount()),
		Forks:       formatCount(repository.GetForksCount()),
		Commits:     commits,
	}

	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"mul": func(a, b int) int { return a * b },
	}
	template, err := template.New("svg").Funcs(funcMap).Parse(svgTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var svgBuffer bytes.Buffer
	err = template.Execute(&svgBuffer, card)
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return svgBuffer.String(), nil
}
