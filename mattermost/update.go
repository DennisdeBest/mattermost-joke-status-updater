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

var jokeMaxLength = 100
var joke string

func Update() {
	arguments := helper.GetArguments()

	url := arguments.Url
	secret := arguments.Secret

	if url == "" {
		log.Fatal("No Mattermost API url has been set")
	}
	if secret == "" {
		log.Fatal("No Mattermost secret has been set")
	}

	joke = getJoke(arguments.MaxTries)

	userData := getUserData(arguments.Url, arguments.Secret)
	setStatus(arguments.Url, arguments.Secret, userData)
}

func getJoke(maxTries int) string {
	var joke *string
	remainingTries := maxTries
	for remainingTries > 0 {
		remainingTries--
		jokeResponse := api.FetchJoke()
		jokeLength := len(jokeResponse)
		if jokeLength <= jokeMaxLength {
			joke = &jokeResponse
			break
		}
	}

	if joke == nil {
		log.Fatal(fmt.Sprintf("No joke found in %d tries for max length %d", maxTries, jokeMaxLength))
	}

	return *joke
}

func getUserData(url string, secret string) UserResponse {
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

	return response
}

func setStatus(url string, secret string, userData UserResponse) {
	client := &http.Client{}
	payload := CustomStatus{
		Emoji: emojis.GetEmoji(),
		Text:  joke,
	}

	postBody, err := json.Marshal(payload)

	post, err := http.NewRequest("PUT", fmt.Sprintf("%s/users/%s/status/custom", url, userData.Id), bytes.NewBuffer(postBody))
	post.Header.Add("Authorization", fmt.Sprintf("BEARER %s", secret))
	post.Header.Add("Accept", "application/json")
	_, err = client.Do(post)
	if err != nil {
		log.Fatal(err)
	}
}
