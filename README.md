[![Build status](https://img.shields.io/github/workflow/status/pushbits/cli/Main)](https://github.com/pushbits/cli/actions)
![License](https://img.shields.io/github/license/pushbits/cli)

# PushBits CLI

| :exclamation:  This software is currently in alpha phase.   |
|-------------------------------------------------------------|

## About

PushBits is a relay server for push notifications.
It enables your services to send notifications via a simple web API, and delivers them to you through [Matrix](https://matrix.org/).
This is similar to what [Pushover](https://pushover.net/) and [Gotify](https://gotify.net/) offer, but it does not require an additional app.

This command line tool enables users to create and modify applications.
Further, it can be used by administrators to add and remove users.

## Installation

It it easiest to download the binary from the [latest release](https://github.com/pushbits/cli/releases).
Alternatively, build it yourself:
```bash
go install github.com/pushbits/cli/cmd/...@latest
```

## Usage

If you are a normal user of PushBits, you will most likely just want to manage your applications.
As an admin, you are possibly trying to manage your users.
This tool provides help for both use cases.

In particular, the tool supports two separate subcommands, `application` and `user`
As the names suggest, the former lets you configure applications, while the latter deals with users.

To use this tool, you need to know the URL of your PushBits instance and the credentials (username and password) of your user.

### Managing Applications

If you forget how to use the tool, always remember to check out the help using the `-h` flag.
Let's do this for the `application` command.

```
$ pbcli application -h
Usage:
  pbcli [OPTIONS] application <command>

Application Options:
      --url=      The URL where the server listens for requests
      --username= The username for authenticating on the server

Help Options:
  -h, --help      Show this help message

Available commands:
  create  Create a new application for a user (aliases: c)
  delete  Delete an existing application for a user (aliases: d)
  list    List all existing applications of the user (aliases: l)
  show    Show details of an existing application of a user (aliases: s)
```

As you can see, there are subcommands for manipulation and displaying of your applications.
Each subcommand has its own help, too.

Let's try creating a new application!
As a first step, check out the help for this subcommand.

```
$ pbcli application create -h
Usage:
  pbcli [OPTIONS] application create name

Application Options:
      --url=      The URL where the server listens for requests
      --username= The username for authenticating on the server

Help Options:
  -h, --help      Show this help message

[create command arguments]
  name:           The name of the application
```

So we need to provide a name for the application.
Of course, we also need to provide the URL of the PushBits instance as well as the username.
This means, to create an application with the name `MyApplication` for the user `MyUsername`, we need to run the following command.

```bash
pbcli application create MyApplication --url https://pushbits.example.com --username MyUsername
```

You will then be asked for your password interactively.

Now you should have a basic idea of how this CLI tool works.
Happy for you to reach out if there are any questions.

## Development

The source code is located on [GitHub](https://github.com/pushbits/cli).
You can retrieve it by checking out the repository as follows.

```bash
git clone https://github.com/pushbits/cli.git
```
