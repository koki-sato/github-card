package card

import (
	"fmt"
	"strings"

	humanize "github.com/dustin/go-humanize"
	"github.com/google/go-github/v64/github"
)

type Option struct {
	UsesFullName bool
}

func GenerateSVG(repository *github.Repository, nCommits int, option Option) string {
	repoNameInImage := repository.GetName()
	if option.UsesFullName {
		repoNameInImage = fmt.Sprintf("%s/%s", repository.GetOwner().GetLogin(), repository.GetName())
	}

	color := LanguageToColor[repository.GetLanguage()]

	description := wrapDescription(repository.GetDescription())
	stars := formatCount(repository.GetStargazersCount())
	forks := formatCount(repository.GetForksCount())

	width := 400
	descriptionY := 50 + (len(description) * 21)
	height := descriptionY + 43
	languageNextX := 16
	if repository.Language != nil {
		languageNextX = 60 + (5 * len(*repository.Language))
	}

	descriptionSVG := ""
	for i, line := range description {
		descriptionSVG += fmt.Sprintf(`
		<g fill="#586068" fill-opacity="1" stroke="#586068" stroke-opacity="1" stroke-width="1" stroke-linecap="square" stroke-linejoin="bevel" transform="matrix(1,0,0,1,0,0)">
			<text fill="#586068" fill-opacity="1" stroke="none" xml:space="preserve" x="17" y="%d" font-family="sans-serif" font-size="14" font-weight="400" font-style="normal">%s</text>
		</g>`, 60+i*21, line)
	}

	svg := fmt.Sprintf(`
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="%d" height="%d" version="1.2" baseProfile="tiny">
	<defs />
	<g fill="none" stroke="black" stroke-width="1" fill-rule="evenodd" stroke-linecap="square" stroke-linejoin="bevel">
		<g fill="#ffffff" fill-opacity="1" stroke="none" transform="matrix(1,0,0,1,0,0)">
			<rect x="0" y="0" width="%d" height="%d"/>
		</g>
		<!-- Boarder -->
		<rect x="0" y="0" width="%d" height="%d" stroke="#eaecef" stroke-width="2"/>
		<!-- Repo name -->
		<g fill="#0366d6" fill-opacity="1" stroke="#0366d6" stroke-opacity="1" stroke-width="1" stroke-linecap="square" stroke-linejoin="bevel" transform="matrix(1,0,0,1,0,0)">
			<a href="%s" target="_blank" rel="noopener noreferrer">
				<text fill="#0366d6" fill-opacity="1" stroke="none" xml:space="preserve" x="17" y="33" font-family="sans-serif" font-size="16" font-weight="630" font-style="normal">%s</text>
			</a>
		</g>
		<!-- Description -->
		%s
		<!-- Language -->
		<circle cx="23" cy="%d" r="7" stroke="none" fill="%s"/>
		<g fill="#24292e" fill-opacity="1" stroke="#24292e" stroke-opacity="1" stroke-width="1" stroke-linecap="square" stroke-linejoin="bevel" transform="matrix(1,0,0,1,0,0)">
			<text fill="#24292e" fill-opacity="1" stroke="none" xml:space="preserve" x="33" y="%d" font-family="sans-serif" font-size="12" font-weight="400" font-style="normal">%s</text>
		</g>
		<!-- Stars -->
		<g fill="#000000" fill-opacity="1" stroke="none" transform="matrix(1,0,0,1,%d,%d)">
			<path vector-effect="none" fill-rule="evenodd" d="M14,6 L9.1,5.36 L7,1 L4.9,5.36 L0,6 L3.6,9.26 L2.67,14 L7,11.67 L11.33,14 L10.4,9.26 L14,6" />
		</g>
		<text fill="#586068" fill-opacity="1" stroke="none" xml:space="preserve" x="%d" y="%d" font-family="sans-serif" font-size="12" font-weight="400" font-style="normal">%s</text>
		<!-- Forks -->
		<g fill="#000000" fill-opacity="1" stroke="none" transform="matrix(1,0,0,1,%d,%d)">
			<path vector-effect="none" fill-rule="evenodd" d="M4,9 L3,9 L3,8 L4,8 L4,9 M4,6 L3,6 L3,7 L4,7 L4,6 M4,4 L3,4 L3,5 L4,5 L4,4 M4,2 L3,2 L3,3 L4,3 L4,2 M12,1 L12,13 C12,13.55 11.55,14 11,14 L6,14 L6,16 L4.5,14.5 L3,16 L3,14 L1,14 C0.45,14 0,13.55 0,13 L0,1 C0,0.45 0.45,0 1,0 L11,0 C11.55,0 12,0.45 12,1 M11,11 L1,11 L1,13 L3,13 L3,12 L6,12 L6,13 L11,13 L11,11 M11,1 L2,1 L2,10 L11,10 L11,1" />
		</g>
		<text fill="#586068" fill-opacity="1" stroke="none" xml:space="preserve" x="%d" y="%d" font-family="sans-serif" font-size="12" font-weight="400" font-style="normal">%s</text>
		<!-- Commits -->
		<text fill="#586068" fill-opacity="1" stroke="none" xml:space="preserve" x="%d" y="%d" font-family="sans-serif" font-size="12" font-weight="400" font-style="normal">%s Commits</text>
	</g>
</svg>
`,
		width, height+1, width, height+1, width, height, repository.GetHTMLURL(), repoNameInImage, strings.TrimPrefix(descriptionSVG, "\n\t\t"),
		descriptionY+21, color, descriptionY+26, repository.GetLanguage(),
		languageNextX, descriptionY+13, languageNextX+20, descriptionY+26, stars,
		languageNextX+40+len(stars)*5, descriptionY+13, languageNextX+60+len(stars)*5, descriptionY+26, forks,
		languageNextX+80+len(stars)*5+len(forks)*5, descriptionY+26, humanize.Comma(int64(nCommits)),
	)
	svg = strings.ReplaceAll(strings.TrimPrefix(svg, "\n"), "\t", "  ")

	return svg
}
