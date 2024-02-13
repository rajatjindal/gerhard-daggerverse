package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/disgoorg/disgo/webhook"
)

type Notify struct{}

// EXAMPLE: dagger call discord --webhook-url env:DISCORD_WEBHOOK --message "Hi from Dagger Notify Module 👋 Learn more at https://github.com/gerhard/daggerverse"
func (n *Notify) Discord(ctx context.Context, webhookURL *Secret, message string) (string, error) {
	if message == "" {
		return "", errors.New("--message cannot be an empty string")
	}

	url, err := webhookURL.Plaintext(ctx)
	if err != nil {
		return "", err
	}

	client, err := webhook.NewWithURL(url)
	if err != nil {
		return "", err
	}
	defer client.Close(ctx)

	m, err := client.CreateContent(message)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("MESSAGE SENT AT: %s\n%s\n", m.CreatedAt, m.Content), err
}
