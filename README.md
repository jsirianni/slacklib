# slacklib
A simple package that makes it easy to generate slack notifications

## Usage
See slacklib_cli.go for an example
```
go build slacklib_cli.go
./slacklib_cli -c "#alert_channel" -m "Server abc failed to start" -h "https://hooks.slack.com/some/hook/endpoint"
```

Example client code
```
var alert slacklib.SlackPost
var hook_url string

func main()  {
    flag.StringVar(&alert.Channel, "c", "#some_channel", "Pass a slack channel")
    flag.StringVar(&alert.Text, "m", "some message", "Pass a message")
    flag.StringVar(&hook_url, "h", "https://hooks.slack.com", "Pass a slack webhook url")
    flag.Parse()

    fmt.Println(slacklib.BasicMessage(alert, hook_url))

    return
}
```

Import into an existing project with `go get github.com/jsirianni/slacklib/slacklib`
```
import github.com/jsirianni/slacklib/slacklib
```
