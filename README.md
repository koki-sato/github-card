# github-card

[![Test](https://github.com/koki-sato/github-card/actions/workflows/test.yml/badge.svg)](https://github.com/koki-sato/github-card/actions/workflows/test.yml)
[![Release](https://github.com/koki-sato/github-card/actions/workflows/release.yml/badge.svg)](https://github.com/koki-sato/github-card/actions/workflows/release.yml)

GitHub Repository Card Generator :octocat:

[![image](./image/github-card.svg)](https://github.com/koki-sato/github-card)

## Usage

```bash
$ github-card --help
GitHub Repository Card Generator

Usage:
  github-card --repo <owner>/<repo> [flags]

Flags:
  -c, --commit          Show the number of commits
  -f, --full-name       Use full name in the card
  -h, --help            help for github-card
  -o, --output string   Output file path (default: <owner>-<repo>.svg)
      --repo string     GitHub repository name (<owner>/<repo>)
```
