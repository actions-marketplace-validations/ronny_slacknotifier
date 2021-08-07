package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ronny/slacknotifier"
	"github.com/slack-go/slack"
)

func main() {
	var token string
	flag.StringVar(&token, "token", "", "Slack app token, usually starts with 'xoxb-'")

	var channelID string
	flag.StringVar(&channelID, "chan", "", "the ID of the channel to send the notification message to")

	var message string
	flag.StringVar(&message, "msg", "", "the notification message")

	var markdown bool
	flag.BoolVar(&markdown, "markdown", true, "whether `msg` (and `msgContext` if given) is/are in markdown or not")

	var botname string
	flag.StringVar(&botname, "botname", "notifier", "the username of the bot")

	var botIconEmoji string
	flag.StringVar(&botIconEmoji, "boticonemoji", ":ghost:", "emoji name to use as the bot icon")

	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "how long to wait when sending notification to slack before giving up")

	var replaceMsgTimestamp string
	flag.StringVar(&replaceMsgTimestamp, "replace", "", "timestamp of a previously-posted message to replace with this message")

	var msgContext string
	flag.StringVar(&msgContext, "context", "", "optional message for context displayed below the main message")

	flag.Parse()

	if token == "" || channelID == "" || message == "" {
		flag.Usage()
		os.Exit(1)
	}

	notifier := &slacknotifier.Notifier{
		BotUsername:  botname,
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
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Sent to %s at %s\n", output.ChannelID, output.Timestamp)
}
