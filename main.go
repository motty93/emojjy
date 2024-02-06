package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

var (
	apiKey string

	eb []EmojiResponse
)

type EmojiResponse struct {
	Code      string `json:"code,omitempty"`
	Character string `json:"character,omitempty"`
	Image     string `json:"image,omitempty"`
	Name      string `json:"name,omitempty"`
	Group     string `json:"group,omitempty"`
	Subgroup  string `json:"subgroup,omitempty"`
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("godotenv load error: %v\n", err)
	}

	apiKey = os.Getenv("NINJA_EMOJI_API_KEY")
}

func apiRequest(name string) {
	uri := "https://api.api-ninjas.com/v1/emoji?name=" + name
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("X-Api-Key", apiKey)

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	spew.Dump(string(bytes))

	if err := json.Unmarshal(bytes, &eb); err != nil {
		log.Fatal(err)
		return
	}

	spew.Dump(eb)
}

func main() {
	ch := make(chan string)
	fmt.Println("Enter emoji name:")

	go func(ch chan string) {
		reader := bufio.NewReader(os.Stdin)
		for {
			name, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
				continue
			}
			ch <- strings.TrimSpace(name)
		}
	}(ch)

	for {
		go func(input string) {
			apiRequest(input)
			fmt.Print("Enter emoji name:")
		}(<-ch)
	}
}
