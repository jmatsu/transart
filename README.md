[![CircleCI](https://circleci.com/gh/jmatsu/artifact-transfer/tree/master.svg?style=svg)](https://circleci.com/gh/jmatsu/artifact-transfer/tree/master)

# transart

transart - Transfer artifacts from services to the single service


## Installation

Download the binary

```
curl -sL "https://raw.githubusercontent.com/jmatsu/transart/master/install.bash" | bash
curl -sL "https://raw.githubusercontent.com/jmatsu/transart/master/install.bash" | VERSION=<...> bash
```

Or build on your local

```
dep ensure
go build .
```

## Getting started

```
# Prepare the configuration file
transart init --save-dir ".transart"

# Configure GitHub Release as a *source* service
transart add github-release --source \
    --username jmatsu --reponame transart --api-token-name GITHUB_TOKEN
    --file-name-pattern ".*.tar"

# Configure Local file system as a *destination* service
transart add local --destination --path tmp

export GITHUB_TOKEN=...

# Download artifacts from GitHub Release and copy only tar files of them to 'tmp' directory
transart transfer

# Download artifacts only
transart download

# Upload artifacts only
transart upload
```

### Bash/Zsh completion

```
// For Bash
eval $(transart --init-completion bash)
transart --init-completion bash >> ~/.bashrc


// For Zsh
eval $(transart --init-completion zsh)
transart --init-completion zsh >> ~/.zshrc
```

## LICENSE

Under MIT License. See [LICENSE](./LICENSE)
