package main

import (
	_ "embed"
	jokeFetcherHelper "github.com/dennisdebest/joke-fetcher/helper"
	"github.com/dennisdebest/mattermost-joke-status-updater/helper"
	"github.com/dennisdebest/mattermost-joke-status-updater/mattermost"
)

func main() {
	mattermost.Update()
}

func init() {
	jokeFetcherHelper.DefineArguments()
	helper.DefineArguments()
}
