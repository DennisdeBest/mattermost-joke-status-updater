package emojis

import (
	_ "embed"
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

//go:embed dataset/emoji.json
var emojisData []byte

type Emojis struct {
	Emojis []string `json:"status_emoji"`
}

func GetEmoji() string {
	var emojis Emojis
	err := json.Unmarshal(emojisData, &emojis)
	if err != nil {
		println("Failed to read the emoji dataset")
		os.Exit(0)
	}
	emoji := getRandomEmoji(emojis.Emojis)

	return emoji
}

func getRandomEmoji(emojis []string) string {
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(emojis)
	return emojis[n]
}
