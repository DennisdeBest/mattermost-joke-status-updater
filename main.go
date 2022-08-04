package main

import (
	_ "embed"
	"github.com/dennisdebest/mattermost-joke-status-updater/helper"
	"github.com/dennisdebest/mattermost-joke-status-updater/mattermost"
)

func main() {
	mattermost.Update()
}

func init() {
	helper.DefineArguments()
}
