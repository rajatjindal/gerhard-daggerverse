package main

import (
	"context"

	"main/internal/dagger"
	"github.com/slack-go/slack"
)

type Slack struct {
}

// Send a message to a specific slack channel or reply in a thread
func (s *Slack) SendMessage(
	ctx context.Context,
	// The slack token to authenticate with the slack organization
	token *dagger.Secret,
	// The sidebar color of the message
	color string,
	// The content of the notification to send
	message string,
	// The channel where to post the message
	channelId string,
	// Set a title to the message
	// +optional
	title string,
	// Set a footer to the message
	// +optional
	footer string,
	// Set an icon in the footer, the icon should be a link
	// +optional
	footerIcon string,
	// Add an image in the message
	// +optional
	ImageUrl string,
	// The thread id if we want to reply to a message or in a thread
	// +optional
	threadId string,
) (string, error) {
	clearToken, err := token.Plaintext(ctx)
	api := slack.New(clearToken)
	attachment := slack.Attachment{
		Color:      color,
		Text:       message,
		MarkdownIn: []string{"text"},
		Title:      title,
		Footer:     footer,
		FooterIcon: footerIcon,
		ImageURL:   ImageUrl,
	}

	options := []slack.MsgOption{
		slack.MsgOptionText("", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	}

	if threadId != "" {
		options = append(options, slack.MsgOptionTS(threadId))
	}

	_, ts, err := api.PostMessage(
		channelId,
		options...,
	)

	if err != nil {
		return "", err
	}

	return ts, nil
}
