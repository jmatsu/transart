[![CircleCI](https://circleci.com/gh/jmatsu/transart/tree/master.svg?style=svg)](https://circleci.com/gh/jmatsu/transart/tree/master)

# transart

transart - Transfer artifacts from multiple services to the single service

Concurrent downloading, uploading and copying are supported by Goroutine.

```bash
transart [-f <configuration file>] <command> [command options...]
```

## Installation

Download the binary directly

```bash
# The latest binary
curl -sL "https://raw.githubusercontent.com/jmatsu/transart/master/install.bash" | bash

# The specific version
curl -sL "https://raw.githubusercontent.com/jmatsu/transart/master/install.bash" | VERSION=<...> bash
```

Or build on your local

```bash
dep ensure
go build .
```

## Getting started

```bash
# Create .transart.yml
transart init --save-dir ".transart"

# Configure GitHub Release as a *source* service
transart add github-release \
    --source \
    --username jmatsu \
    --reponame transart \
    --api-token-name GITHUB_TOKEN

# Configure Local file system as a *destination* service
# save files into tmp directory
transart add local \
    --destination \
    --path tmp \
    --file-name-pattern ".*\.tar$" # tar files only

export GITHUB_TOKEN=...

# Download artifacts from GitHub Release and copy only tar files of them to 'tmp' directory
transart transfer

# Download artifacts only
transart download

# Upload artifacts only
transart upload
```

### Supported services

Service|Source|Destination
:---|:---|:---
CircleCI| :white_check_mark: | :x:
GitHub Release| :white_check_mark: | :white_check_mark:
Local File System| :white_check_mark: | :white_check_mark:

### Configurations

#### Basic structure

```
version: 1
save_dir: <path>
source:
  locations:
  - <location config>
destination:
  location: <location config>
```

#### Location config

ref. [key definitions](../config/option_key.go)

`add` command is available to add each configuration of services.

*CircleCI*

```yaml
type: circleci
vcs-type: <github|bitbucket>
username: <the username of the project> # case sensitive
reponame: <the name of the project> # case sensitive

# optional
branch: <the name of the branch>
api-token-name: <environment name> #CIRCLECI_TOKEN is used by default
```

*GitHub Release*

```yaml
type: github-release
username: <the username of the repository>
reponame: <the name of the respoitory>

# optional
strategy: <draft|create|draft-or-create> # draft-or-create is used by default
api-token-name: <environment name> #GITHUB_TOKEN is used by default
```

*Local file system*

```yaml
type: local
path: <path>

# optional
file-name-pattern: <regexp pattern>
```

### Bash/Zsh completion

```bash
// For Bash
eval $(transart --init-completion bash)
transart --init-completion bash >> ~/.bashrc


// For Zsh
eval $(transart --init-completion zsh)
transart --init-completion zsh >> ~/.zshrc
```

## LICENSE

Under MIT License. See [LICENSE](./LICENSE)
