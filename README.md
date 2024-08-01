<h1 align="center">Send message to Microsoft Teams.</h1>

<p align="center">
  Send messages to teams from the command line or Go program.
  <br>
  Go module to send messages to MS Teams channel
  <br>
  using webhooks.
</p>

<p align="center">
  <img src="screenshot1.png" alt="teams">
</p>


**Command line:** Post messages to MS Teams channel from command line.

**Proxy support:** Works behind a proxy

**Flexible:** Flexible and simple interface to use in Go Programs.

![unit-test-run-time](https://coveritup.app/badge?org=kevincobain2000&repo=go-msteams&branch=feature/cover&type=unit-test-run-time)
![coverage](https://coveritup.app/badge?org=kevincobain2000&repo=go-msteams&branch=feature/cover&type=coverage)
![build-time](https://coveritup.app/badge?org=kevincobain2000&repo=go-msteams&branch=feature/cover&type=build-time)
![go-binary-size](https://coveritup.app/badge?org=kevincobain2000&repo=go-msteams&branch=feature/cover&type=go-binary-size)
![go-mod-dependencies](https://coveritup.app/badge?org=kevincobain2000&repo=go-msteams&branch=feature/cover&type=go-mod-dependencies)


![unit-test-run-time](https://coveritup.app/chart?org=kevincobain2000&repo=go-msteams&branch=feature/cover&type=unit-test-run-time&theme=light&line=fill&width=150&height=150&output=svg)
![coverage](https://coveritup.app/chart?org=kevincobain2000&repo=go-msteams&branch=feature/cover&type=coverage&theme=light&line=fill&width=150&height=150&output=svg)
![build-time](https://coveritup.app/chart?org=kevincobain2000&repo=go-msteams&branch=feature/cover&type=build-time&theme=light&line=fill&width=150&height=150&output=svg)
![go-binary-size](https://coveritup.app/chart?org=kevincobain2000&repo=go-msteams&branch=feature/cover&type=go-binary-size&theme=light&line=fill&width=150&height=150&output=svg)
![go-mod-dependencies](https://coveritup.app/chart?org=kevincobain2000&repo=go-msteams&branch=feature/cover&type=go-mod-dependencies&theme=light&line=fill&width=150&height=150&output=svg)








## Installation

Install binary

```sh
curl -sLk https://raw.githubusercontent.com/kevincobain2000/go-msteams/master/install.sh | sh
./go-msteams --help
```

Or using Go

```sh
go install github.com/kevincobain2000/go-msteams@latest
```

### From Command Line:

```sh
go-msteams -hook=<required ms-teams webhook> -title= -subtitle= -subject= -message="<h1>HTML OK</h1>
```

### From Go program:

```go
import (
    gmt "github.com/kevincobain2000/go-msteams/src"
)

func main() {
    hook := "https://required" //required
    proxy := "" //optional

	details := []gmt.Details{
		{
			Label:   "Subtitle",
			Message: "subtitle",
		},
		{
			Label:   "Subject",
			Message: "subject",
		},
		{
			Label:   "Message",
			Message: "message",
		},
	}

    err := gmt.Send(title, details, hook, proxy)
    if err != nil {
        log.Println(err)
    } else {
        log.Println("Successfully sent!")
    }
}
```

--

### Command Line Params Help:

```sh
go-msteams -h
  -color string
    	Description: Your theme color
  -hook string
    	(required) Description: MS Teams webhook URL
  -message string
    	Description: Message body. HTML allowed. (default "My Message")
  -proxy string
    	Description: Hit behind this proxy
  -subject string
    	Description: Your subject (default "My Subject")
  -subtitle string
    	Description: Your summary (default "My Summary")
  -title string
    	Description: Your title (default "My Title")
```
