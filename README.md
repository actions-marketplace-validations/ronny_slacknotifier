# Slack Notifier

A command line tool and a GitHub Action to post message notifications to Slack.

## GitHub Action

See [action.yml] for details on inputs and outputs.

[action.yml]: action.yml

Example:

```yaml
# `uses: ronny/slacknotifier@v1` should work too but it will build the docker
# image from scratch every time, which will take about a minute or so, that's
# probably not what you want 🙂
uses: docker://ronny/slacknotifier:1
with:
  slack-token: ${{ secrets.SLACK_TOKEN }}
  bot-name: deploybot
  bot-icon-emoji: ":mega:"
  channel-id: "C12345"
  message: "Deployed `todo-service`"
  message-context: "Env: `production` | By: `${{ github.actor }}` | Commit: `${{ github.sha }}`"
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
