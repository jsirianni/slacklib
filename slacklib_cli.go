package main

import (
    "fmt"
    "flag"

    "./slacklib"
)


// Basic client cli for testing the slacklib package
//


var alert slacklib.SlackPost
var hook_url string


func main()  {
    flag.StringVar(&alert.Channel, "c", "#some_channel", "Pass a slack channel")
    flag.StringVar(&alert.Text, "m", "#some_channel", "Pass a message")
    flag.StringVar(&hook_url, "h", "https://hooks.slack.com", "Pass a slack webhook url")
    flag.Parse()

    fmt.Println(slacklib.BasicMessage(alert, hook_url))

    return
}
