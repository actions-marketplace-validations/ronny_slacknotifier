# Slack Notifier

A command line tool and a GitHub Action to post message notifications to Slack.

## GitHub Action

See [action.yml] for details on inputs and outputs.

[action.yml]: action.yml

Example:

```yaml
uses: ronny/slacknotifier@v1
with:
  slack-token: xoxb-whatever
  bot-name: deploybot
  bot-icon-emoji: ":mega:"
  channel-id: "C12345"
  message: "Deployed `todo-service`"
  message-context: "Env: `production` | By: $GITHUB_ACTOR | Commit: $GITHUB_SHA"
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
