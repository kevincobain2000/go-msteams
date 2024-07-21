package src

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetCard verifies that the getCard function creates a card with the correct content
func TestGetCard(t *testing.T) {
	title := "Test Title"
	subtitle := "Test Subtitle"
	subject := "Test Subject"
	message := "Test Message"

	card := getCard(title, subtitle, subject, message)

	// Check the top-level fields
	if card.Type != "message" {
		t.Errorf("expected card type to be 'message', got '%s'", card.Type)
	}
	if len(card.Attachments) != 1 {
		t.Errorf("expected 1 attachment, got %d", len(card.Attachments))
	}

	attachment := card.Attachments[0]
	if attachment.ContentType != "application/vnd.microsoft.card.adaptive" {
		t.Errorf("expected content type to be 'application/vnd.microsoft.card.adaptive', got '%s'", attachment.ContentType)
	}

	content := attachment.Content
	if content.Schema != "http://adaptivecards.io/schemas/adaptive-card.json" {
		t.Errorf("expected schema to be 'http://adaptivecards.io/schemas/adaptive-card.json', got '%s'", content.Schema)
	}
	if content.Type != "AdaptiveCard" {
		t.Errorf("expected content type to be 'AdaptiveCard', got '%s'", content.Type)
	}
	if content.Version != "1.4" {
		t.Errorf("expected version to be '1.4', got '%s'", content.Version)
	}
	if content.AccentColor != "bf0000" {
		t.Errorf("expected accent color to be 'bf0000', got '%s'", content.AccentColor)
	}

	// Check the body of the card
	body := content.Body
	if len(body) != 3 {
		t.Errorf("expected 3 body elements, got %d", len(body))
	}

	// Check the title text block
	titleBlock, ok := body[0].(textBlock)
	if !ok {
		t.Errorf("expected first body element to be textBlock, got %T", body[0])
	}
	if titleBlock.Text != title {
		t.Errorf("expected title text to be '%s', got '%s'", title, titleBlock.Text)
	}
	if titleBlock.Size != "large" {
		t.Errorf("expected title size to be 'large', got '%s'", titleBlock.Size)
	}
	if titleBlock.Weight != "bolder" {
		t.Errorf("expected title weight to be 'bolder', got '%s'", titleBlock.Weight)
	}
	if titleBlock.Color != "accent" {
		t.Errorf("expected title color to be 'accent', got '%s'", titleBlock.Color)
	}

	// Check the fact set
	factSetBlock, ok := body[1].(factSet)
	if !ok {
		t.Errorf("expected second body element to be factSet, got %T", body[1])
	}
	if len(factSetBlock.Facts) != 2 {
		t.Errorf("expected 2 facts, got %d", len(factSetBlock.Facts))
	}
	if factSetBlock.Facts[0].Value != subtitle {
		t.Errorf("expected subtitle to be '%s', got '%s'", subtitle, factSetBlock.Facts[0].Value)
	}
	if factSetBlock.Facts[1].Value != subject {
		t.Errorf("expected subject to be '%s', got '%s'", subject, factSetBlock.Facts[1].Value)
	}

	// Check the code block
	codeBlock, ok := body[2].(codeBlock)
	if !ok {
		t.Errorf("expected third body element to be codeBlock, got %T", body[2])
	}
	if codeBlock.CodeSnippet != message {
		t.Errorf("expected code snippet to be '%s', got '%s'", message, codeBlock.CodeSnippet)
	}
	if codeBlock.FontType != "monospace" {
		t.Errorf("expected font type to be 'monospace', got '%s'", codeBlock.FontType)
	}
	if !codeBlock.Wrap {
		t.Errorf("expected wrap to be true, got %v", codeBlock.Wrap)
	}
}

// TestDispatch verifies that the dispatch method sends a POST request with the correct payload
func TestDispatch(t *testing.T) {
	hookURL := "/"
	proxyURL := ""

	title := "Test Title"
	subtitle := "Test Subtitle"
	subject := "Test Subject"
	message := "Test Message"

	card := getCard(title, subtitle, subject, message)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected method POST, got %s", r.Method)
		}

		if r.URL.String() != hookURL {
			t.Errorf("expected URL %s, got %s", hookURL, r.URL.String())
		}

		var receivedCard msTeamCard
		err := json.NewDecoder(r.Body).Decode(&receivedCard)
		if err != nil {
			t.Fatalf("error decoding request body: %v", err)
		}

		expectedCard := card
		if receivedCard.Type != expectedCard.Type {
			t.Errorf("expected card type to be %s, got %s", expectedCard.Type, receivedCard.Type)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	err := card.dispatch(server.URL, proxyURL)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
