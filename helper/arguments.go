package helper

import (
	"flag"
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
}

func GetArguments() Arguments {
	flag.Parse()
	return Arguments{Secret: secret, Url: url, MaxTries: maxTries}
}
