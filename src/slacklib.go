package slacklib

import (
    "fmt"
    "bytes"
    "encoding/json"
    "net/http"
)


type SlackPost struct {
    Channel string
    Text    string
}



// BasicMessage takes a SlackPost struct and a webhook url as
// arguments and sends a message. Advanced webhook features
// are not supported with this function
func BasicMessage(message SlackPost, hook_url string) bool {
    result := sendPayload(buildPayload(message), hook_url)

    return result
}


func buildPayload(message SlackPost) []byte {
    var payload []byte
    var err error
	payload, err = json.Marshal(message)
    if err != nil {
        return nil
    } else {
        return payload
    }
}


func sendPayload(payload []byte, hook_url string) bool {
    req, err := http.NewRequest("POST", hook_url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil || resp.StatusCode != 200 {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}
