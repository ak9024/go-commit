![GitHub License](https://img.shields.io/github/license/ak9024/go-commit)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/ak9024/go-commit/.github%2Fworkflows%2Fgo.yml)
![GitHub Downloads (all assets, latest release)](https://img.shields.io/github/downloads/ak9024/go-commit/latest/total)
![GitHub top language](https://img.shields.io/github/languages/top/ak9024/go-commit)

# go-commit

The CLI (Command Line Interface) helps you generate commits automatically, using karma style for git convention

### Prerequisite

- [Install Go](https://go.dev/doc/install)
- [OPENAI_API_KEY](https://platform.openai.com/account/api-keys)

### Install

```bash
go install github.com/ak9024/go-commit@latest
```

## Install via CURL

> If you don't installed go binary, please follow this instruction to install with CURL.

```bash
# download install.sh from repo go-commit
curl -O https://raw.githubusercontent.com/ak9024/go-commit/main/install.sh
# please open ./install.sh with your editor and suitable with your config
# OS=Linux
# VERSION=v0.1.*
nvim ./install.sh
# open write access
chmod +x ./install.sh
# execute install.sh
./install.sh
# move go commit into your local binary
sudo mv go-commit /usr/local/bin/
```

### Setup

```bash
export OPENAI_API_KEY=<token>

# or open file ~/.zshrc
# add "export OPENAI_API_KEY=<OpenAI Key>"
# source ~/.zshrc
```

### How to run?

> Please setup `OPENAI_API_KEY` before usage go-commit

```bash
cd <your-repo>
git add .
go-commit commit

# for help
go-commit -h
```
