# SASM

Send A Slack Message (SASM).

## Create a Slack token

* Open Slack
* Click on a workspace
* Administration
* Manage apps
* Query for `hooks`
* Incoming WebHooks
* Add to Slack
* Create a new channel
* Create
* Add people
* Add Incoming WebHooks Integration

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
slack_channel: someSlackChannel
slack_token: someSlackToken
```

Save the file. Update the permissions to ensure that you are the
only man that is able to read and write to the config file:

```bash
chmod 0600 ~/.sasm/config.yml
```

## Send a slack message

```bash
docker run \
    -v /home/${USER}/.sasm/:/home/sasm/.sasm/ utrecht/sasm:0.1.0 \
    sasm plain --config /home/sasm/.sasm/config.yml -thello
```
