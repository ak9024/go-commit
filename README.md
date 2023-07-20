<div align="center">
    <h2>go-commit</h2>
    <p>The CLI (Command Line Interface) to help you generated commit automatically, using karma style for git convention</p>
</div>

### Prerequisite

- [Go](https://go.dev/doc/install)
- [OpenAI Key](https://platform.openai.com/account/api-keys)

### Install

```bash
go install github.com/ak9024/go-commit@latest
```

## Install via CURL

```bash
# download install.sh from repo go-commit
curl -O https://raw.githubusercontent.com/ak9024/go-commit/main/install.sh
# please ope ./install.sh with your editor and suitable with your config
nvim ./install.sh
# open write access
chmod +x ./install.sh
# execute install.sh
./install.sh
# move go commit into your local binary
sudo mv go-commit /usr/local/bin/
# run the binary
go-commit
```

### Setup

```bash
export OPENAI_API_KEY=<token>

# or open file ~/.zshrc
# add "export OPENAI_API_KEY=<OpenAI Key>"
# source ~/.zshrc
```

### How to run?

```bash
cd <your-repo>
git add .
go-commit commit

# for help
go-commit -h
```
