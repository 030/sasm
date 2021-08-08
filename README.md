# SASM

![Docker Pulls](https://img.shields.io/docker/pulls/utrecht/sasm.svg)

Send A Slack Message (SASM).

## Create a Slack token

* Open Slack
* Click on a workspace
* `Administration` or `Settings & administration`
* Manage apps
* Query for `bots`
* Click on `Bots Connect a bot to the Slack Real Time Messaging API.`
* Add to Slack
* Enter a username for the bot in the `Username` field
* Click on `Add bot integration`
* Copy the API Token that appears on the screen

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

## Send a slack message

[![dockeri.co](https://dockeri.co/image/utrecht/sasm)](https://hub.docker.com/r/utrecht/sasm)

```bash
docker run \
    -v /home/${USER}/.sasm/:/home/sasm/.sasm/ utrecht/sasm:1.0.0 \
    sasm plain --config /home/sasm/.sasm/config.yml -thello
```

## Troubleshooting

```bash
FATA[0000] channel_not_found
```

Add the created bot to the channel by inviting it.
