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

	if err := json.Unmarshal(bytes, &eb); err != nil {
		log.Fatal(err)
		return
	}

	for _, e := range eb {
		fmt.Printf("%v%2s%v\n\n", e.Character, "", e.Code)
	}
}

func main() {
	ch := make(chan string)
	fmt.Println("Enter emoji name: ")

	go func(ch chan string) {
		reader := bufio.NewReader(os.Stdin)
		for {
			name, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
				continue
			}

			name = strings.TrimSpace(name)
			switch name {
			case "exit":
				os.Exit(0)
			case "":
				fmt.Print("Enter emoji name: ")
			default:
				ch <- name
			}
		}
	}(ch)

	for {
		go func(input string) {
			apiRequest(input)
			fmt.Print("Enter emoji name: ")
		}(<-ch)
	}
}
