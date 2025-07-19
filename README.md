# github-card

[![Test](https://github.com/koki-sato/github-card/actions/workflows/test.yaml/badge.svg)](https://github.com/koki-sato/github-card/actions/workflows/test.yaml)
[![Release](https://github.com/koki-sato/github-card/actions/workflows/release.yaml/badge.svg)](https://github.com/koki-sato/github-card/actions/workflows/release.yaml)

GitHub Repository Card Generator :octocat:

[![image](./image/github-card.svg)](https://github.com/koki-sato/github-card)

The card design is heavily inspired by [gh-card](https://github.com/nwtgck/gh-card).

## Installation

### mise

You can install github-card using [mise](https://mise.jdx.dev/).

```bash
mise install ubi:koki-sato/github-card@latest
```

### GitHub Releases

You can download a binary from [GitHub Releases](https://github.com/koki-sato/github-card/releases).

### Install from Source Code

You can build and install github-card from source code via `go install`.

```bash
go install github.com/koki-sato/github-card@latest
```

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
  -v, --version         version for github-card
```
