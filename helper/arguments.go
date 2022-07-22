package helper

import (
	"flag"
)

type Arguments struct {
	Secret string
	Url    string
}

var secret string
var url string

func DefineArguments() {
	flag.StringVar(&secret, "secret", "", "Mattermost secret")
	flag.StringVar(&url, "url", "", "Mattermost URL")
}

func GetArguments() Arguments {
	flag.Parse()
	return Arguments{Secret: secret, Url: url}
}
