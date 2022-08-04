package main

import (
	_ "embed"
	"fmt"
	"github.com/dennisdebest/mattermost-joke-status-updater/helper"
	"github.com/dennisdebest/mattermost-joke-status-updater/mattermost"
)

func main() {
	logs := mattermost.Update()
	fmt.Println(logs)
}

func init() {
	helper.DefineArguments()
}
