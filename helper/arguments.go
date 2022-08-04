package helper

import (
	"flag"
	jokeFetcherHelper "github.com/dennisdebest/joke-fetcher/helper"
)

type Arguments struct {
	Secret   string
	Url      string
	MaxTries int
}

var secret string
var url string
var maxTries int

func DefineArguments() {
	flag.StringVar(&secret, "secret", "", "Mattermost secret")
	flag.StringVar(&url, "url", "", "Mattermost URL")
	flag.IntVar(&maxTries, "maxTries", 10, "The maxTries of request to get a joke inferioir to the max length")
	jokeFetcherHelper.DefineArguments()
}

func GetArguments() Arguments {
	flag.Parse()
	return Arguments{Secret: secret, Url: url, MaxTries: maxTries}
}
