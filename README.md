# SASM

Send A Slack Message (SASM).

## Create an App, Channel and Slack token

* Navigate to [Your Apps](https://api.slack.com/apps)
* Click on `Create New App`
* `From scratch`
* Choose an `App Name`
* `Pick a workspace to develop your app in`
* `Create App`
* Add features and functionality: `Permissions`
* Add an OAuth Scope in the `Bot Token Scopes`: `chat:write`
* `Install to Workspace`
* Create a channel and invite the bot
* Navigate to [Your Apps](https://api.slack.com/apps) again
* Click on app that you created
* Click on `OAuth & Permissions`
* Copy `Bot User OAuth Token`

## Configure sasm

Create a directory:

```bash
mkdir ~/.sasm
```

Create a config file:

```bash
vim ~/.sasm/config.yml
```

Add a Slack token:

```bash
---
slack_channel: someSlackChannelID-clickOnChannelAndCopyTheID
slack_token: someSlackToken
```

Save the file. Update the permissions to ensure that you are the
only one that is able to read and write to the config file:

```bash
chmod 0600 ~/.sasm/config.yml
```

## Troubleshooting

```bash
FATA[0000] channel_not_found
```

Add the created bot to the channel by inviting it.
