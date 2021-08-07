# Slack Notifier

A command line tool and a GitHub Action to post message notifications to Slack.

This app uses the latest recommended, non-legacy Slack API to post messages,
including Block Kit for the messages.

## GitHub Action

See [action.yml] for details on inputs and outputs.

[action.yml]: action.yml

Example:

```yaml
- uses: ronny/slacknotifier@v1
  with:
    slack-token: ${{ secrets.SLACK_TOKEN }}
    bot-name: deploybot
    bot-icon-emoji: ":mega:"
    channel-id: "C12345"
    message: "Deployed `todo-service`"
    message-context: "Env: `production` | By: `${{ github.actor }}` | Commit: `${{ github.sha }}`"
```

You can also use the docker image directly, but you need to specify all of the
inputs including the ones that have default values in [action.yml].

```yaml
- uses: docker://ronny/slacknotifier:1
  with:
    slack-token: ${{ secrets.SLACK_TOKEN }}
    bot-name: deploybot
    bot-icon-emoji: ":mega:"
    channel-id: "C12345"
    message: "Deployed `todo-service`"
    message-context: "Env: `production` | By: `${{ github.actor }}` | Commit: `${{ github.sha }}`"
    # These are the defaults:
    markdown: 'true'
    replace-timestamp: ''
    timeout: 30s
```

## CLI tool

```
go install github.com/ronny/slacknotifier/cmd/slack-notify
$GOPATH/bin/slack-notify
```

Or, if you have the source locally:

```
make install
$GOPATH/bin/slack-notify
```

## Slack token

The recommended token type is a [bot token]. Follow the [guide to create a new
Slack app](https://api.slack.com/authentication/basics), the scopes needed are:
`chat:write`, `chat:write.public` (if you want the bot to be able to post
messages without being invited to a channel), and `chat:write.customize` (if you
want to customize the bot's name and icon in a step).

[bot token]: https://api.slack.com/authentication/token-types#bot

Once you created your app, you can find the bot token from the "OAuth &
Permissions" section of your app, the URL of the page looks like
`https://api.slack.com/apps/Axxxxxxxx/oauth`.

## Channel ID

It’s recommended to use a channel’s canonical ID, which can be found at the
bottom of the pop-up dialog when you click a channel’s name.

Using `#name` or just `name` sometimes work, but not when updating/replacing a
message. So it’s best to use the canonical channel ID always.
