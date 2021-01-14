![License](https://img.shields.io/github/license/pushbits/cli)

# PushBits CLI

## About

PushBits is a relay server for push notifications.
It enables your services to send notifications via a simple web API, and delivers them to you through [Matrix](https://matrix.org/).
This is similar to what [PushBullet](https://www.pushbullet.com/), [Pushover](https://pushover.net/), and [Gotify](https://gotify.net/) offer, but a lot less complex.

This command line tool enables users to create and modify applications.
Further, it can be used by administrators to administrate to add and remove users.

## Installation

The following command will download, build, and install the tool.

```bash
go get -u github.com/pushbits/cli/cmd/pbcli
```

Note that this requires [a working Go 1.12+ installation](https://golang.org/doc/install).

## Development

The source code is located on [GitHub](https://github.com/pushbits/cli).
You can retrieve it by checking out the repository as follows.

```bash
git clone https://github.com/pushbits/cli.git
```
