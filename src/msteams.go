package src

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// msTeamCard is Adaptive Card for Team notification
type msTeamCard struct {
	Type        string       `json:"type"`
	Attachments []attachment `json:"attachments"`
}

type attachment struct {
	ContentType string      `json:"contentType"`
	ContentURL  *string     `json:"contentUrl"`
	Content     cardContent `json:"content"`
}

type cardContent struct {
	Schema      string        `json:"$schema"`
	Type        string        `json:"type"`
	Version     string        `json:"version"`
	AccentColor string        `json:"accentColor"`
	Body        []interface{} `json:"body"`
	Actions     []action      `json:"actions"`
	MSTeams     msTeams       `json:"msteams"`
}

type textBlock struct {
	Type   string `json:"type"`
	Text   string `json:"text"`
	ID     string `json:"id,omitempty"`
	Size   string `json:"size,omitempty"`
	Weight string `json:"weight,omitempty"`
	Color  string `json:"color,omitempty"`
}

type fact struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type factSet struct {
	Type  string `json:"type"`
	Facts []fact `json:"facts"`
	ID    string `json:"id"`
}

type action struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

type msTeams struct {
	Width string `json:"width"`
}

type Details struct {
	Label   string `json:"label"`
	Message string `json:"value"`
}

func Send(title string, details []Details, hookURL, proxyURL string) (err error) {
	card := getCard(title, details)
	return card.dispatch(hookURL, proxyURL)
}

func getCard(title string, details []Details) msTeamCard {
	facts := []fact{}
	for _, v := range details {
		facts = append(facts, fact{
			Title: v.Label,
			Value: v.Message,
		})
	}

	return msTeamCard{
		Type: "message",
		Attachments: []attachment{
			{
				ContentType: "application/vnd.microsoft.card.adaptive",
				ContentURL:  nil,
				Content: cardContent{
					Schema:      "http://adaptivecards.io/schemas/adaptive-card.json",
					Type:        "AdaptiveCard",
					Version:     "1.4",
					AccentColor: "bf0000",
					Body: []interface{}{
						textBlock{
							Type:   "TextBlock",
							Text:   title,
							ID:     "title",
							Size:   "large",
							Weight: "bolder",
							Color:  "accent",
						},
						factSet{
							Type:  "FactSet",
							Facts: facts,
							ID:    "acFactSet",
						},
					},
					MSTeams: msTeams{
						Width: "Full",
					},
				},
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
	log.Println("response", resp.Status)

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}
