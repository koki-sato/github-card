package main

import "github.com/koki-sato/github-card/cmd"

// version of the application, set during build by GoReleaser.
// See https://goreleaser.com/cookbooks/using-main.version/
var version string

func main() {
	cmd.Version = version
	cmd.Execute()
}
