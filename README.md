# HTTP CLI Tool

A Command-Line Interface (CLI) tool to make HTTP requests.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
    - [get](#get)
    - [post](#post)
    - [put](#put)
    - [delete](#delete)
- [Configuration](#configuration)
- [Advanced Features](#advanced-features)
- [Examples](#examples)

## Installation

To install the HTTP CLI tool, you need to have Go installed on your system. You can download and install it from [golang.org](https://golang.org/dl/).

Clone the repository and build the tool:

```sh
git clone https://github.com/yourusername/http-cli-tool.git
cd http-cli-tool
go build -o http-cli cmd/http-cli/main.go
mv http-cli /usr/local/bin/
```

## Usage
The HTTP CLI tool supports various HTTP methods like GET, POST, PUT, and DELETE.
```bash
http-cli [command] [flags]
```