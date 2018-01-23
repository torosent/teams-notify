# Microsoft Teams-Notify

:star: Forked from: `https://github.com/technosophos/slack-notify`

Teams-Notify is a very simple tool for sending a Microsoft Teams notification via a
Teams Incoming Webhook.

It is designed to function as a 12-factor app, receiving configuration via
environment variables. To keep things simple, it is rigid in what it allows you
to set. 

## Teams Incoming Webhooks

This tool uses [Teams Incoming Webhooks](https://docs.microsoft.com/en-us/microsoftteams/platform/concepts/connectors)
to send a message to your teams channel.

Before you can use this tool, you need to log into Microsoft Teams and configure
this.

## Usage

Running `teams-notify` in a shell prompt goes like this:

```console
$ export TEAMS_WEBHOOK=https://outlook.office.com/webhook/Txxxxxx/IncomingWebhook/Bxxxxxx/xxxxxxxx
$ TEAMS_MESSAGE="hello" teams-notify
```

Running the Docker container goes like this:

```console
$ export TEAMS_WEBHOOK=https://outlook.office.com/webhook/Txxxxxx/IncomingWebhook/Bxxxxxx/xxxxxxxx
$ docker run -e TEAMS_WEBHOOK=$TEAMS_WEBHOOK -e TEAMS_MESSAGE="hello" torosent/teams-notify
```

### In Brigade

You can easily use this inside of brigade hooks. Here is an example from
[hello-helm](https://github.com/technosophos/hello-helm):


```javascript
const {events, Job} = require("brigadier")

events.on("imagePush", (e, p) => {

  var teams = new Job("teams-notify", "torosent/teams-notify:latest", ["/teams-notify"])

  // This doesn't need access to storage, so skip mounting to speed things up.
  teams.storage.enabled = false
  teams.env = {
    // It's best to store the teams webhook URL in a project's secrets.
    TEAMS_WEBHOOK: p.secrets.TEAMS_WEBHOOK,
    TEAMS_TITLE: "Message Title",
    TEAMS_MESSAGE: "Message Body",
    TEAMS_COLOR: "#0000ff"
  }
  teams.run()
})
```



## Environment Variables

```shell
# The Microsoft Teams self-assigned webhook
TEAMS_WEBHOOK=https://outlook.office.com/webhook/Txxxxxx/IncomingWebhook/Bxxxxxx/xxxxxxxx

# The title of the message
TEAMS_TITLE="Hello World"
# The body of the message
TEAMS_MESSAGE="Today is a fine day"
# RGB color to for message formatting. (Teams determines what is colored by this)
TEAMS_COLOR="#efefef"
```

## Build It

Configure:

```
make bootstrap
```

Compile:

```
make build
```

Publish to DockerHub

```
make docker-build docker-push
```
