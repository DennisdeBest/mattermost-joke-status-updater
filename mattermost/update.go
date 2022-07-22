package mattermost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dennisdebest/joke-fetcher/api"
	"github.com/dennisdebest/mattermost-joke-status-updater/emojis"
	"github.com/dennisdebest/mattermost-joke-status-updater/helper"
	"io/ioutil"
	"log"
	"net/http"
)

type UserResponse struct {
	Id       string            `json:"id"`
	Username string            `json:"username"`
	Email    string            `json:"email"`
	Props    UserResponseProps `json:"props"`
}

type UserResponseProps struct {
	Status string `json:"customStatus"`
}

type CustomStatus struct {
	Emoji string `json:"emoji"`
	Text  string `json:"text"`
}

func Update() {
	api.FetchJoke()
	arguments := helper.GetArguments()
	getStatus(arguments.Url, arguments.Secret)
}

func getStatus(url string, secret string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/users/me", url), nil)
	req.Header.Add("Authorization", fmt.Sprintf("BEARER %s", secret))
	req.Header.Add("Accept", "application/json")
	resp, _ := client.Do(req)

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response UserResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		log.Fatal(err)
	}

	payload := CustomStatus{
		Emoji: emojis.GetEmoji(),
		Text:  api.Joke,
	}

	postBody, err := json.Marshal(payload)

	post, err := http.NewRequest("PUT", fmt.Sprintf("%s/users/%s/status/custom", url, response.Id), bytes.NewBuffer(postBody))
	post.Header.Add("Authorization", fmt.Sprintf("BEARER %s", secret))
	post.Header.Add("Accept", "application/json")
	resp, _ = client.Do(post)
	if err != nil {
		log.Fatal(err)
	}
}
