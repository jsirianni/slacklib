package slacklib

import (
    "fmt"
    "bytes"
    "strconv"
    "encoding/json"
    "net/http"
)


type SlackPost struct {
    Channel string `json:"channel"`
    Text    string `json:"text"`
}



// BasicMessage takes a SlackPost struct and a webhook url as
// arguments and sends a message. Advanced webhook features
// are not supported with this function
func BasicMessage(message SlackPost, hook_url string) bool {
    payload := buildPayload(message)
    if payload == nil {
        return false
    }
    fmt.Println(string(payload))
    return sendPayload(payload, hook_url)
}


func buildPayload(message SlackPost) []byte {
    var payload []byte
    var err error
	payload, err = json.Marshal(message)
    if err != nil {
        fmt.Println("Failed to marshal payload")
        fmt.Println(err.Error())
        return nil
    } else {
        return payload
    }
}


func sendPayload(payload []byte, hook_url string) bool {
    req, err := http.NewRequest("POST", hook_url, bytes.NewBuffer(payload))
    if err != nil {
        fmt.Println(err.Error())
        return false
    }
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
    if err != nil {
        fmt.Println(err.Error())
    }

	if resp.StatusCode != 200 {
		fmt.Println("Slack returned status: " + strconv.Itoa(resp.StatusCode))
		return false
	} else {
		return true
	}
}
