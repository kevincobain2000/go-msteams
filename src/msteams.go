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

// msTeamCard is Adaptive Card for Team notification
type msTeamCard struct {
	Type    string    `json:"type"`
	Version string    `json:"version"`
	Body    []body    `json:"body"`
	Actions []action  `json:"actions"`
}

// body is sub-struct of msTeamCard
type body struct {
	Type     string   `json:"type"`
	Text     string   `json:"text"`
	Items    []item   `json:"items"`
}

// item is sub-struct of body
type item struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
}

// action is sub-struct of msTeamCard
type action struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

func Send(title, subtitle, subject, color, message, hookURL, proxyURL string) (err error) {
	card := getCard(title, subtitle, subject, color, message)
	return card.dispatch(hookURL, proxyURL)
}

func getCard(title, subtitle, subject, color, message string) msTeamCard {
	return msTeamCard{
		Type:    "AdaptiveCard",
		Version: "1.2",
		Body: []body{
			{
				Type: "TextBlock",
				Text: title,
			},
			{
				Type: "TextBlock",
				Text: subtitle,
			},
			{
				Type: "TextBlock",
				Text: subject,
			},
			{
				Type: "TextBlock",
				Text: message,
			},
		},
		Actions: []action{
			{
				Type:  "Action.OpenUrl",
				Title: "Learn More",
				URL:   "https://adaptivecards.io",
			},
		},
	}
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
