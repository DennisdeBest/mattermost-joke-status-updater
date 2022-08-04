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
	"os"
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

type CallLog struct {
	Url  string `json:"url"`
	Joke string `json:"joke"`
}

var CallLogs []CallLog
var jokeMaxLength = 100
var joke string
var apiUrl string

func Update() []CallLog {
	arguments := helper.GetArguments()

	url := arguments.Url
	secret := arguments.Secret

	if url == "" {
		checkAndAssignEnvVar(&url, "MATTERMOST_URL")
	}
	if secret == "" {
		checkAndAssignEnvVar(&secret, "MATTERMOST_SECRET")
	}

	jokeResponse := getJoke(arguments.MaxTries)

	if jokeResponse != nil {
		joke = *jokeResponse
		userData := getUserData(url, secret)
		setStatus(url, secret, userData)
	}

	return CallLogs
}

func checkAndAssignEnvVar(value *string, variable string) string {
	var envVarExists bool
	*value, envVarExists = os.LookupEnv(variable)
	if !envVarExists {
		log.Fatal(fmt.Sprintf("No %s variable has been set", variable))
	}
	return *value
}

func getJoke(maxTries int) *string {
	var joke *string
	remainingTries := maxTries
	for remainingTries > 0 {
		remainingTries--
		jokeResponse := api.FetchJoke()
		apiUrl = api.LatestApiUrl
		CallLogs = append(CallLogs, CallLog{
			Url:  apiUrl,
			Joke: jokeResponse,
		})
		jokeLength := len(jokeResponse)
		if jokeLength <= jokeMaxLength {
			joke = &jokeResponse
			break
		}
	}

	if joke == nil {
		log.Print(fmt.Sprintf("No joke found in %d tries for max length %d", maxTries, jokeMaxLength))
	}

	return joke
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
