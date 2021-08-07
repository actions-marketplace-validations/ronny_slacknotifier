package slacknotifier

import (
	"context"
	"errors"

	"github.com/slack-go/slack"
)

type SlackClient interface {
	UpdateMessageContext(ctx context.Context, channelID string, timestamp string, options ...slack.MsgOption) (_chan string, _timestamp string, _text string, err error)
	PostMessageContext(ctx context.Context, channelID string, options ...slack.MsgOption) (_chan string, _timestamp string, err error)
}

type Notifier struct {
	BotUsername  string
	BotIconEmoji string
	Client       SlackClient
}

type NotifyInput struct {
	ChannelID                   string
	Message                     string
	MessageContext              string
	Markdown                    bool
	TimestampOfMessageToReplace string
}

type NotifyOutput struct {
	ChannelID string
	Timestamp string
}

func (n *Notifier) Notify(ctx context.Context, input NotifyInput) (*NotifyOutput, error) {
	if n.Client == nil {
		return nil, errors.New("BUG: SlackClient is nil")
	}

	elementType := slack.PlainTextType
	if input.Markdown {
		elementType = slack.MarkdownType
	}

	blocks := []slack.Block{
		slack.NewSectionBlock(
			slack.NewTextBlockObject(elementType, input.Message, !input.Markdown, false),
			nil,
			nil,
		),
	}
	if input.MessageContext != "" {
		blocks = append(blocks,
			slack.NewContextBlock("ctx",
				slack.NewTextBlockObject(elementType, input.MessageContext, !input.Markdown, false),
			),
		)
	}

	options := []slack.MsgOption{
		slack.MsgOptionIconEmoji(n.BotIconEmoji),
		slack.MsgOptionUsername(n.BotUsername),
		slack.MsgOptionText(input.Message, true), // true -> escapes &, <, and > to &amp;, &lt;, and &gt;
		slack.MsgOptionBlocks(blocks...),
	}

	output := NotifyOutput{}
	var err error
	{
		if input.TimestampOfMessageToReplace != "" {
			output.ChannelID, output.Timestamp, _, err = n.Client.UpdateMessageContext(ctx, input.ChannelID, input.TimestampOfMessageToReplace, options...)
		} else {
			output.ChannelID, output.Timestamp, err = n.Client.PostMessageContext(ctx, input.ChannelID, options...)
		}
	}
	if err != nil {
		return nil, err
	}
	return &output, nil
}
