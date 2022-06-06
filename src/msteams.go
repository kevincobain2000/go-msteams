package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// msTeamCard is MessageCard for Team notification
type msTeamCard struct {
	Type       string    `json:"@type"`
	Context    string    `json:"@context"`
	Summary    string    `json:"summary"`
	ThemeColor string    `json:"theme_color"`
	Title      string    `json:"title"`
	Sections   []section `json:"sections"`
}

// section is sub-struct of msTeamCard
type section struct {
	ActivityTitle    string `json:"activityTitle"`
	ActivitySubtitle string `json:"activitySubtitle"`
	ActivityImage    string `json:"activityImage"`
	Facts            []fact `json:"facts"`
}

// fact is sub-struct of section
type fact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func Send(title, subtitle, subject, color, message, hookURL, proxyURL string) (err error) {
	card := msTeamCard{
		Type:       "MessageCard",
		Context:    "http://schema.org/extensions",
		Summary:    subject,
		ThemeColor: color,
		Title:      title,
		Sections: []section{
			{
				ActivityTitle:    subject,
				ActivitySubtitle: subtitle,
				ActivityImage:    "",
				Facts: []fact{
					{
						Name:  "Message",
						Value: message,
					},
				},
			},
		},
	}
	return card.dispatch(hookURL, proxyURL)
}

// dispatch is send message to webhook
func (card *msTeamCard) dispatch(hookURL, proxyURL string) (err error) {
	requestBody, err := json.Marshal(card)
	if err != nil {
		return err
	}

	var client http.Client
	timeout := time.Duration(5 * time.Second)

	if proxyURL != "" {
		proxy, err := url.Parse(proxyURL)
		if err != nil {
			return err
		}
		transport := &http.Transport{Proxy: http.ProxyURL(proxy)}
		client = http.Client{
			Transport: transport,
			Timeout:   timeout,
		}
	} else {
		client = http.Client{
			Timeout: timeout,
		}
	}

	request, err := http.NewRequest("POST", hookURL, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		return err
	}

	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	body := string(respBody)
	if body != "1" {
		return fmt.Errorf("error: %v", body)
	}
	return nil
}
