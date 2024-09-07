# goup

`goup` is a CLI tool designed to help developers quickly initialize and manage Go projects with ease. It automates the setup process, including the creation of Go modules, tidying up dependencies, and setting up an initial `main.go` file. Whether you're starting a new Go project or managing existing ones, `goup` simplifies the process.

## Installation

To install `goup` via Homebrew, simply run the following command:

```bash
% brew install RyotaroSeto/tap/goup
```

This will download and install `goup` onto your system, making it available as a global command.

## Upgrade

To upgrade `goup` to the latest version, use the following command:

```bash
% brew upgrade RyotaroSeto/tap/goup
```

## Usage

Here is how you can use `goup` to initialize a new Go project:

1. Initialize a new Go project by running the following command:

```bash
% goup init <project-name>
```

This will:

- Run `go mod init <project-name>` to initialize the Go module.
- Run `go mod tidy` to clean up and synchronize dependencies.
- Create a basic `main.go` file with a "Hello, World!" example.

## Example
```bash
% goup init my-new-project
```
This will create a new Go project named `my-new-project` with the necessary setup files.

## LICENSE
[MIT](./LICENSE)
