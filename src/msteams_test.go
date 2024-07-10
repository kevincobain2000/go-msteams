package src

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendNoHook(t *testing.T) {
	title := "My Title"
	subtitle := "My Summary"
	subject := "My Subject"
	color := ""
	message := "My Message"
	hook := "https://"
	proxy := ""
	err := Send(title, subtitle, subject, color, message, hook, proxy)
	fmt.Println(err)
	assert.NotNil(t, err)
}

func TestCard(t *testing.T) {
	title := "My Title"
	subtitle := "My Summary"
	subject := "My Subject"
	color := ""
	message := "My Message"
	card := getCard(title, subtitle, subject, color, message)
	assert.NotNil(t, card)

	expectedCard := msTeamCard{
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

	assert.Equal(t, expectedCard, card)
}
