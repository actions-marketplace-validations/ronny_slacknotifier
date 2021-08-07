package main

import (
	"context"
	"time"

	"github.com/ronny/slacknotifier"
	"github.com/sethvargo/go-githubactions"
	"github.com/slack-go/slack"
)

const (
	DefaultBotUsername  = "slacknotifier"
	DefaultBotIconEmoji = ":ghost:"
)

func main() {
	token := githubactions.GetInput("slack-token")
	if token == "" {
		githubactions.Fatalf("missing 'slack-token'")
	}

	botUsername := githubactions.GetInput("bot-name")
	if botUsername == "" {
		botUsername = DefaultBotUsername
	}

	botIconEmoji := githubactions.GetInput("bot-icon-emoji")
	if botIconEmoji == "" {
		botIconEmoji = DefaultBotIconEmoji
	}

	channelID := githubactions.GetInput("channel-id")
	if channelID == "" {
		githubactions.Fatalf("missing 'channel-id'")
	}

	message := githubactions.GetInput("message")
	if message == "" {
		githubactions.Fatalf("missing 'message'")
	}

	msgContext := githubactions.GetInput("message-context")
	markdown := githubactions.GetInput("markdown") == "true"
	replaceMsgTimestamp := githubactions.GetInput("replace-timestamp")
	timeoutString := githubactions.GetInput("timeout")
	timeout, err := time.ParseDuration(timeoutString)
	if err != nil {
		githubactions.Fatalf("timeout: time.ParseDuration: %s\n", err)
	}

	notifier := &slacknotifier.Notifier{
		BotUsername:  botUsername,
		BotIconEmoji: botIconEmoji,
		Client:       slack.New(token),
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), timeout)
	defer cancelCtx()

	output, err := notifier.Notify(ctx, slacknotifier.NotifyInput{
		ChannelID:                   channelID,
		Message:                     message,
		MessageContext:              msgContext,
		Markdown:                    markdown,
		TimestampOfMessageToReplace: replaceMsgTimestamp,
	})
	if err != nil {
		githubactions.Fatalf("notifer.Notify: %s\n", err)
	}

	githubactions.SetOutput("channel-id", output.ChannelID)
	githubactions.SetOutput("timestamp", output.Timestamp)
}
